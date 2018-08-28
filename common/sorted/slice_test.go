package sorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testCmp(a interface{}, b interface{}) int {
	ai := a.(int)
	bi := b.(int)
	if ai < bi {
		return -1
	} else if ai > bi {
		return 1
	} else {
		return 0
	}
}

func TestSlice(t *testing.T) {
	slice := NewSlice(testCmp)
	slice.Push(3)
	slice.Push(2)
	slice.Push(4)
	assert.Equal(t, slice.Left(), 2)
	assert.Equal(t, slice.Right(), 4)
	assert.Equal(t, slice.PopLeft(), 2)
	slice.Del(4)
	slice.Push(1)
	assert.Equal(t, slice.Right(), 3)
	assert.Equal(t, slice.Left(), 1)
}
