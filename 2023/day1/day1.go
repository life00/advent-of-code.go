package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// function that fixes word numbers that overlap e.g. (tw(o)ne)
func fixInput(rawInput *[]byte) {
	// temporarily save it as other variable
	input := string(*rawInput)
	// all broken word number combinations
	brokenStrings := [8]string{"oneight", "twone", "threeight", "fiveight", "sevenine", "eightwo", "eighthree", "nineight"}
	strings := [8]string{"oneeight", "twoone", "threeeight", "fiveeight", "sevennine", "eighttwo", "eightthree", "nineeight"}
	// loop through arrays
	for i := 0; i < 8; i++ {
		// replace strings[i] with ints[i]
		reg := regexp.MustCompile(brokenStrings[i])
		input = reg.ReplaceAllString(input, strings[i])
	}
	// assign it back
	*rawInput = []byte(input)
}

// function that replaces proper word numbers with numbers
func replaceNumbers(rawInput *[]byte) {
	// define strings and ints
	strings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	ints := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	// temporarily save it as other variable
	input := string(*rawInput)
	// loop through arrays
	for i := 0; i < 9; i++ {
		// replace strings[i] with ints[i]
		reg := regexp.MustCompile(strings[i])
		input = reg.ReplaceAllString(input, ints[i])
	}
	// assign it back
	*rawInput = []byte(input)
}

func main() {
	// read the input
	rawInput, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fixInput(&rawInput)
	replaceNumbers(&rawInput)
	// split it into an array based on newline character
	arrInput := strings.Split(string(rawInput), "\n")
	// remove the last faulty element in the array
	if len(arrInput) > 0 {
		arrInput = arrInput[:len(arrInput)-1]
	}
	// preparing environment
	reg, err := regexp.Compile("[^0-9]+") // regex for removing non digits
	result := 0                           // saving all the digits
	if err != nil {
		fmt.Println(err)
	}
	// loop through every line
	for _, s := range arrInput {
		// remove all non digit characters
		s = reg.ReplaceAllString(s, "")
		// get the first and last character of a string
		s = string(s[0]) + string(s[len(s)-1])
		// convert string to integer
		d, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		// add the value to the sum
		result += d
	}
	// print out the result
	fmt.Printf("%d\n", result)
}
