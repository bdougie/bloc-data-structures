package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeIsInitialized(t *testing.T) {
	b := new(BinaryTree)
	assert.NotNil(t, b)
}

func TestNodeIsInitialized(t *testing.T) {
	n := new(Node)
	assert.NotNil(t, n)
}

func TestRootIsNotInitialized(t *testing.T) {
	b := new(BinaryTree)
	assert.Nil(t, b.root)
}

func TestRootIsAddedIfNil(t *testing.T) {
	b := new(BinaryTree)
	b.Insert(10)
	assert.NotNil(t, b.root)
}

func TestRootIsNotChanged(t *testing.T) {
	b := new(BinaryTree)
	b.Insert(10)
	b.Insert(12)
	assert.Equal(t, b.root.value, 10)
}

func TestValueGreaterThenRootIsSetRight(t *testing.T) {
	b := new(BinaryTree)
	b.Insert(10)
	b.Insert(12)
	assert.Equal(t, b.root.Left, 12)
}
