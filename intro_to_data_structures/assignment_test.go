package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// A line of people at an amusement park ride

func TestInitializationOfLine(t *testing.T) {
	SpaceMountain := AmusementRide{[]string{"bill", "susan", "mary", "sam"}}
	if length := len(SpaceMountain.line); length != 4 {
		t.Errorf("Expected line length of 4, but it was %d instead.", length)
	}
}

func TestAddingToLine(t *testing.T) {
	SpaceMountain := AmusementRide{[]string{"bill", "susan", "mary", "sam"}}
	addToEndOfLine(&SpaceMountain, "wayne")

	assert.Equal(t, len(SpaceMountain.line), 5)
}

func TestRemovingFromLine(t *testing.T) {
	SpaceMountain := AmusementRide{[]string{"bill", "susan", "mary", "sam"}}
	removeFromLine(&SpaceMountain, "sam")

	if length := len(SpaceMountain.line); length != 3 {
		t.Errorf("Expected line length of 3, but it was %d instead.", length)
	}
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
	length := len(c.pixels)

	if length != 3 {
		t.Errorf("Expected pixel length of 3, but it was %d instead.", length)
	}

	addPixelToScreen(&c, createPixel())

	if length != 3 {
		t.Errorf("Expected pixel length of 4, but it was %d instead.", length)
	}
}

// The geo-location of every person on Earth

func TestPeopleTracking(t *testing.T) {
	someone := trackSomeone()
	assert.NotNil(t, someone.location)
}
