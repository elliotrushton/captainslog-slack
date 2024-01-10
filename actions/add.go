package actions

import (
	"github.com/elliotrushton/captainslog-slack/clock"
	"github.com/elliotrushton/captainslog-slack/logbook"
)

func Add(captainsLog *logbook.Log, txt string) string {
	when := clock.RealClock{}
	captainsLog.Add(txt, when.Now())
	return txt
}
