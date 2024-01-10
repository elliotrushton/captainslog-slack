package logbook

import (
	"fmt"
	"time"
)

type Entry struct {
	message string
	when    time.Time
}

type Todo struct {
	Message string
}

type Log struct {
	contents []Entry
	Todos    []Todo
}

func (l *Log) Add(addendum string, when time.Time) *Log {
	entry := Entry{message: addendum, when: when}
	l.contents = append(l.contents, entry)
	return l
}

func (l *Log) AddTodo(message string) *Log {
	todo := Todo{Message: message}
	l.Todos = append(l.Todos, todo)
	return l
}

func (l *Log) CompleteTodo(idx int) *Log {
	var updated []Todo
	// there's gotta be a better way
	for i, t := range l.Todos {
		if i != idx-1 {
			updated = append(updated, t)
		}
	}
	l.Todos = updated
	return l
}

func (l *Log) String(tz *time.Location) string {

	report := ""
	for _, entry := range l.contents {
		report = fmt.Sprintf("%s%s - %s\n", report, entry.when.In(tz).Format("15:04"), entry.message)
	}
	return report
}
