package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// A line of people at an amusement park ride

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

// Pixels on a computer screen

func TestPixelCreation(t *testing.T) {
	p := createPixel()
	assert.NotNil(t, p)
	assert.NotNil(t, p.colors)
	assert.NotNil(t, p.coordinates)
}

func TestComputerScreenInitHasPixels(t *testing.T) {
	c := fillComputerScreen()
	assert.NotNil(t, c)
}

func TestPixelIsAddedToTheScreen(t *testing.T) {
	c := fillComputerScreen()
	assert.Equal(t, len(c.pixels), 3)

	addPixelToScreen(&c, createPixel())
	assert.Equal(t, len(c.pixels), 4)
}

// The geo-location of every person on Earth

func TestPeopleTracking(t *testing.T) {
	someone := trackSomeone()
	assert.NotNil(t, someone.location)
}
