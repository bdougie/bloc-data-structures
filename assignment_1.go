package main

import (
	"bufio"
	"fmt"
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
	fmt.Print("Who would you like to add to the line: ")

	// Input
	text, _ := reader.ReadString('\n')

	addToEndOfLine(&SpaceMountain, text)

	fmt.Print("Who would you like to remove from the line? ")
	text2, _ := reader.ReadString('\n')

	removeFromLine(&SpaceMountain, text2)
}

func main() {
	manipulateTheLine()
}
