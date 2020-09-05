package doublylinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Init(t *testing.T) {
	l := &List{}
	list := l.Init()

	assert.NotNil(t, list)
}

func TestList_New(t *testing.T) {
	list := New()

	assert.NotNil(t, list)
}

func TestList_First_Returns_Nil_On_Empty_List(t *testing.T) {
	list := New()
	assert.Nil(t, list.First())
}

func TestList_Last_Returns_Nil_On_Empty_List(t *testing.T) {
	list := New()
	assert.Nil(t, list.Last())
}

func TestList_First(t *testing.T) {
	list := New()
	list.Push(1)

	assert.NotNil(t, list.First())
}

func TestList_Last(t *testing.T) {
	list := New()
	list.Push(1)

	assert.NotNil(t, list.Last())
}

func TestList_Len_Returns_Zero_Value_On_Empty_List(t *testing.T) {
	list := New()
	value := list.Len()

	assert.Zero(t, value)
}

func TestList_Len(t *testing.T) {
	list := New()
	list.Push(1)
	list.Push(1)

	count := list.Len()

	assert.Equal(t, 2, count)
}