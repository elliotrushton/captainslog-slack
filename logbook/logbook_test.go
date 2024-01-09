package logbook

import (
	"testing"

	"github.com/elliotrushton/captainslog-slack/clock"
	"github.com/stretchr/testify/assert"
)

func TestLog_String(t *testing.T) {
	log := Log{}
	when := clock.TestClock{}
	log.Add("Hello world", when.Now())
	when.AddMinutes(5)
	log.Add("Let's eat lunch", when.Now())

	out := log.String()
	assert.Equal(t, "13:26 - Hello world\n13:31 - Let's eat lunch\n", out)
}
