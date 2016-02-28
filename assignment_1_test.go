package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitializationOfLine(t *testing.T) {
	SpaceMountain := AmusementRide{[]string{"bill", "susan", "mary", "sam"}}
	assert.Equal(t, len(SpaceMountain.line), 4)
}

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

func TestPixelCreation(t *testing.T) {
	p := createPixel()
	assert.NotNil(t, p)
	assert.NotNil(t, p.colors)
	assert.NotNil(t, p.coordinates)
}

func TextComputerScreenInitHasPixels(t *testing.T) {
	c := fillComputerScreen()
	assert.NotNil(t, c)
}
