package doublylinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_Next_Return_Proper_Pointer(t *testing.T) {
	list := New()

	list.Push(1)
	list.Push(13)

	actual := list.First().Next()
	assert.NotNil(t, actual)
}

func TestNode_Prev_Return_Proper_Pointer(t *testing.T) {
	list := New()

	list.Push(1)
	list.Push(13)

	actual := list.Last()
	assert.NotNil(t, actual)
}

func TestNode_Prev_Returns_Nil_On_Currently_Nil_Node(t *testing.T) {
	list := New()

	assert.Nil(t, list.First().Prev())
}

func TestNode_Next_Returns_Nil_On_Currently_Nil_Node(t *testing.T) {
	list := New()

	assert.Nil(t, list.First().Next())
}

func TestNode_Next_Return_Nil_On_Last_Node(t *testing.T) {
	list := New()
	list.Push(1)

	actual := list.Last().Next()

	assert.Nil(t, actual)
}

func TestNode_Prev_Return_Nil_On_First_Node(t *testing.T) {
	list := New()
	list.Push(1)

	actual := list.First().Prev()
	assert.Nil(t, actual)
}

func TestNode_Can_Return_Value(t *testing.T) {
	list := New()
	list.Push(1)

	value := list.First().Value
	assert.Equal(t, 1, value)
}
