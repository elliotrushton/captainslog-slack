package actions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/elliotrushton/captainslog-slack/logbook"
)

func Todo(captainsLog *logbook.Log, tokens []string) string {
	if len(tokens) == 1 {
		// no todo description, so interpret as "show me my todos"
		output := ""

		for idx, todo := range captainsLog.Todos {
			output += fmt.Sprintf("%d - %s\n", idx+1, todo.Message)
		}
		return output
	} else {
		// add todo
		todoBody := strings.Join(tokens[1:], " ")
		captainsLog.AddTodo(todoBody)
		return "Added Todo: " + todoBody
	}
}

func CompletedTodo(captainsLog *logbook.Log, tokens []string) string {
	var dones []int

	for _, doneStr := range tokens[1:] {
		doneIdx, err := strconv.Atoi(doneStr)
		if err != nil {
			return "Error - marked todo must be identified by number. I didn't understand: " + doneStr
		}
		dones = append(dones, doneIdx-1)
	}

	sort.Ints(dones)
	reverseSlice(dones)

	for idx, _ := range dones {
		captainsLog.CompleteTodo(idx)
	}

	return fmt.Sprintf("Completed %d todos!\n", len(dones))
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
