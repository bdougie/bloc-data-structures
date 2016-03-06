package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitializationOfMystack(t *testing.T) {
	m := new(Stack)
	assert.NotNil(t, m)
}

func TestStackIsInitializedEmpty(t *testing.T) {
	s := new(Stack)
	assert.Equal(t, true, s.isEmpty())
}

func TesdtItemCanBePushedOntoStack(t *testing.T) {
	s := new(Stack)
	s.Push("12")
	assert.Equal(t, false, s.isEmpty())
}

func TesdtItemCanBePopOffStack(t *testing.T) {
	s := new(Stack)
	s.Push("7")
	assert.Equal(t, false, s.isEmpty())
	s.Pop()
	assert.Equal(t, true, s.isEmpty())
}
