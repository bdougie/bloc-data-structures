package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// A line of people at an amusement park ride

var line []string

func line_at_ride() {

	// CMD Line interface
	reader := bufio.NewReader(os.Stdin)
	line = []string{"bill", "susan", "mary", "sam"}

	fmt.Println("The following people are in line", line)
	fmt.Print("Who would you like to add to the line: ")

	// Input
	text, _ := reader.ReadString('\n')

	line = addToEndOfLine(text, line[:])
	fmt.Println(line)

	fmt.Print("Who would you like to remove from the line? ")
	text2, _ := reader.ReadString('\n')

	line = removeFromLine(text2, line[:])
	fmt.Println(line)
}

func addToEndOfLine(person string, line []string) []string {
	return append(line, person)
}

func removeFromLine(person string, line []string) []string {
	for i, p := range line {
		match, _ := regexp.MatchString(p, strings.ToLower(person))

		if match {
			line = append(line[:i], line[i+1:]...)
		}
	}

	return line
}

func main() {
	line_at_ride()
}
