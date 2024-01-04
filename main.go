package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

const startupMessage = `Captains Log booting up...`

func logRequest(r *http.Request) {
	uri := r.RequestURI
	method := r.Method
	fmt.Println("Got request!", method, uri)
}

type Log struct {
	contents []string
}

func (l *Log) Add(addendum string) *Log {
	l.contents = append(l.contents, addendum)
	return l
}

func (l *Log) String() string {
	dump := strings.Join(l.contents, "\n")
	return dump + "\n"
}

func main() {
	captainsLog := &Log{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error parsing form from request: " + err.Error())
			return
		}

		txt := r.Form.Get("text")
		if txt[0] == '!' {
			// This is a command string
			tokens := strings.Split(txt, " ")
			cmd := strings.ToLower(tokens[0])
			switch cmd {
			case "!show":
				fmt.Fprintf(w, captainsLog.String())
			case "!search":
				fmt.Fprintf(w, "Search command\n")
			default:
				fmt.Fprintf(w, "Unrecognized command!\n")
			}
		} else {
			// Just a message to log
			captainsLog.Add(txt)
			fmt.Fprintf(w, "%s\n", txt)
		}
	})

	http.HandleFunc("/cached", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		maxAgeParams, ok := r.URL.Query()["max-age"]
		if ok && len(maxAgeParams) > 0 {
			maxAge, _ := strconv.Atoi(maxAgeParams[0])
			w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", maxAge))
		}
		responseHeaderParams, ok := r.URL.Query()["headers"]
		if ok {
			for _, header := range responseHeaderParams {
				h := strings.Split(header, ":")
				w.Header().Set(h[0], strings.TrimSpace(h[1]))
			}
		}
		statusCodeParams, ok := r.URL.Query()["status"]
		if ok {
			statusCode, _ := strconv.Atoi(statusCodeParams[0])
			w.WriteHeader(statusCode)
		}
		requestID := uuid.Must(uuid.NewV4())
		fmt.Fprint(w, requestID.String())
	})

	http.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		keys, ok := r.URL.Query()["key"]
		if ok && len(keys) > 0 {
			fmt.Fprint(w, r.Header.Get(keys[0]))
			return
		}
		headers := []string{}
		headers = append(headers, fmt.Sprintf("host=%s", r.Host))
		for key, values := range r.Header {
			headers = append(headers, fmt.Sprintf("%s=%s", key, strings.Join(values, ",")))
		}
		fmt.Fprint(w, strings.Join(headers, "\n"))
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		keys, ok := r.URL.Query()["key"]
		if ok && len(keys) > 0 {
			fmt.Fprint(w, os.Getenv(keys[0]))
			return
		}
		envs := []string{}
		envs = append(envs, os.Environ()...)
		fmt.Fprint(w, strings.Join(envs, "\n"))
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		codeParams, ok := r.URL.Query()["code"]
		if ok && len(codeParams) > 0 {
			statusCode, _ := strconv.Atoi(codeParams[0])
			if statusCode >= 200 && statusCode < 600 {
				w.WriteHeader(statusCode)
			}
		}
		requestID := uuid.Must(uuid.NewV4())
		fmt.Fprint(w, requestID.String())
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	for _, encodedRoute := range strings.Split(os.Getenv("ROUTES"), ",") {
		if encodedRoute == "" {
			continue
		}
		pathAndBody := strings.SplitN(encodedRoute, "=", 2)
		path, body := pathAndBody[0], pathAndBody[1]
		http.HandleFunc("/"+path, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, body)
		})
	}

	bindAddr := fmt.Sprintf(":%s", port)
	lines := strings.Split(startupMessage, "\n")
	fmt.Println()
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println()
	fmt.Printf("==> Server listening at %s 🚀\n", bindAddr)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}
