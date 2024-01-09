package logbook

import (
	"fmt"
	"time"
)

type Entry struct {
	message string
	when    time.Time
}

type Log struct {
	contents []Entry
}

func (l *Log) Add(addendum string, when time.Time) *Log {
	entry := Entry{message: addendum, when: when}
	l.contents = append(l.contents, entry)
	return l
}

func (l *Log) String(tz *time.Location) string {

	report := ""
	for _, entry := range l.contents {
		report = fmt.Sprintf("%s%s - %s\n", report, entry.when.In(tz).Format("15:04"), entry.message)
	}
	return report
}
