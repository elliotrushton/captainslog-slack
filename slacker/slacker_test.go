package slacker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlacker_GetTimezoneForUser(t *testing.T) {
	sl := NewSlacker()
	_, err := sl.GetTimezoneForUser("U06B8HBD2D8")
	assert.Nil(t, err)
}
