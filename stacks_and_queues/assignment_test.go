package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitializationOfMystack(t *testing.T) {
	m := new(MyStack)
	assert.NotNil(t, m)
}

func TesdtItemCanBePushedOntoStack(t *testing.T) {
}
