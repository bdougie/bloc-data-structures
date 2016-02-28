package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddingToLine(t *testing.T) {
	SpaceMountain := AmusementRide{[]string{"bill", "susan", "mary", "sam"}}
	addToEndOfLine(&SpaceMountain, "wayne")

	assert.Equal(t, len(SpaceMountain.line), 5)
}

func TestRemovingFromLine(t *testing.T) {
	SpaceMountain := AmusementRide{[]string{"bill", "susan", "mary", "sam"}}
	removeFromLine(&SpaceMountain, "sam")

	assert.Equal(t, len(SpaceMountain.line), 3)
}
