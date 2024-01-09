package clock

import "time"

type Clock interface {
	Now() time.Time
	// After(d time.Duration) <-chan time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }

// func (realClock) After(d time.Duration) <-chan time.Time { return time.After(d) }

type TestClock struct {
	deltaMin int
}

func (t TestClock) Now() time.Time {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	return time.Date(2023, 03, 05, 13, 26+t.deltaMin, 00, 00, loc)
}

func (t *TestClock) AddMinutes(mins int) {
	t.deltaMin += mins
}
