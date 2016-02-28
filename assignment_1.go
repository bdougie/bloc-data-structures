package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

type AmusementRide struct {
	line []string
}

func addToEndOfLine(a *AmusementRide, person string) {
	a.line = append(a.line, person)
	fmt.Println(a.line)
}

func removeFromLine(a *AmusementRide, person string) {
	for i, p := range a.line {
		match, _ := regexp.MatchString(p, strings.ToLower(person))

		if match {
			a.line = append(a.line[:i], a.line[i+1:]...)
			fmt.Println(a.line)
		}
	}
}

func manipulateTheLine() {
	SpaceMountain := AmusementRide{[]string{"bill", "susan", "mary", "sam"}}

	// CMD Line interface creation
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("The following people are in line", SpaceMountain.line)
	fmt.Println("Who would you like to add to the line: ")

	// Input
	text, _ := reader.ReadString('\n')

	addToEndOfLine(&SpaceMountain, text)

	fmt.Println("Who would you like to remove from the line? ")
	text2, _ := reader.ReadString('\n')

	removeFromLine(&SpaceMountain, text2)
}

type Pixel struct {
	colors      [3]int
	coordinates [2]int
}

type ComputerScreen struct {
	pixels []Pixel
}

func createPixel() Pixel {
	p := Pixel{
		colors:      [3]int{rand.Intn(100), rand.Intn(100), rand.Intn(100)},
		coordinates: [2]int{rand.Intn(100), rand.Intn(100)}}
	return p
}

func fillComputerScreen() ComputerScreen {
	j, k, l := createPixel(), createPixel(), createPixel()
	ps := []Pixel{j, k, l}

	return ComputerScreen{ps}
}

func main() {
	manipulateTheLine()
}
