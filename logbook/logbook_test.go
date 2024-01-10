package logbook

import (
	"testing"
	"time"

	"github.com/elliotrushton/captainslog-slack/clock"
	"github.com/stretchr/testify/assert"
)

func TestLog_String(t *testing.T) {
	log := Log{}
	when := clock.TestClock{}
	log.Add("Hello world", when.Now())
	when.AddMinutes(5)
	log.Add("Let's eat lunch", when.Now())

	loc, _ := time.LoadLocation("America/New_York")
	out := log.String(loc)
	assert.Equal(t, "13:26 - Hello world\n13:31 - Let's eat lunch\n", out)
}

func TestAddTodo(t *testing.T) {
	log := Log{}
	log.AddTodo("test todo")
	assert.Equal(t, "test todo", log.Todos[0].Message)
}

func TestCompleteTodo(t *testing.T) {
	log := Log{}
	log.AddTodo("test1").AddTodo("test2").AddTodo("test3")
	log.CompleteTodo(2)

	assert.Equal(t, 2, len(log.Todos))
	assert.Equal(t, "test1 test3", log.Todos[0].Message+" "+log.Todos[1].Message)
}
