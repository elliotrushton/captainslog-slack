package actions

import (
	"strings"
	"testing"

	"github.com/elliotrushton/captainslog-slack/logbook"
	"github.com/stretchr/testify/assert"
)

func TestTodo(t *testing.T) {
	cl := &logbook.Log{}

	body := Todo(cl, strings.Split("!todo buy velvet pants", " "))

	assert.Equal(t, "Added Todo: buy velvet pants", body)
}

func TestShowTodos(t *testing.T) {
	cl := &logbook.Log{}
	cl.AddTodo("buy velvet pants")
	cl.AddTodo("pet dogs")

	body := Todo(cl, []string{"!todo"})
	assert.Equal(t, "1 - buy velvet pants\n2 - pet dogs\n", body)
}
