package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashIsInitialized(t *testing.T) {
	h := new(HashClass)
	assert.NotNil(t, h)
}

func TestHashIsInitializedWithAHash(t *testing.T) {
	h := new(HashClass)
	assert.NotNil(t, h.hash)
}

func TestHashCanInsertAKeyValuePair(t *testing.T) {
	h := new(HashClass)
	h.Insert("Key", "Value")
	assert.NotNil(t, h.store)
	assert.Equal(t, 1, len(h.store))
}

func TestHashCanInsertTwoKeyValuePairs(t *testing.T) {
	h := new(HashClass)
	h.Insert("The Lord of the Rings: The Fellowship of the Ring", "3 hours, 48 minutes")
	h.Insert("The Hobbit: An Unexpected Journey", "3 hours, 2 minutes")
	assert.Equal(t, 2, len(h.store))
}
