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
	fmt.Printf("Fixing broken word numbers...\n")
	// temporarily save it as other variable
	input := string(*rawInput)
	// all broken word number combinations
	brokenStrings := [8]string{"oneight", "twone", "threeight", "fiveight", "sevenine", "eightwo", "eighthree", "nineight"}
	strings := [8]string{"oneeight", "twoone", "threeeight", "fiveeight", "sevennine", "eighttwo", "eightthree", "nineeight"}
	// loop through arrays
	for i := 0; i < 8; i++ {
		fmt.Printf("%s --> %s\n", brokenStrings[i], strings[i])
		// replace strings[i] with ints[i]
		reg := regexp.MustCompile(brokenStrings[i])
		input = reg.ReplaceAllString(input, strings[i])
	}
	fmt.Printf("Done!\n")
	// assign it back
	*rawInput = []byte(input)
}

// function that replaces proper word numbers with numbers
func replaceNumbers(rawInput *[]byte) {
	fmt.Printf("Replacing word numbers with numbers...\n")
	// define strings and ints
	strings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	ints := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	// temporarily save it as other variable
	input := string(*rawInput)
	// loop through arrays
	for i := 0; i < 9; i++ {
		fmt.Printf("%s --> %s\n", strings[i], ints[i])
		// replace strings[i] with ints[i]
		reg := regexp.MustCompile(strings[i])
		input = reg.ReplaceAllString(input, ints[i])
	}
	fmt.Printf("Done!\n")
	// assign it back
	*rawInput = []byte(input)
}

// extracts the numbers to an array
func getNumbers(rawInput *[]byte) (numbers []int) {
	fmt.Printf("Getting numbers from the data...\n")
	// split it into an array based on newline character
	arrInput := strings.Split(string(*rawInput), "\n")
	// remove the last faulty element in the array
	if len(arrInput) > 0 {
		arrInput = arrInput[:len(arrInput)-1]
	}
	// preparing environment
	reg, err := regexp.Compile("[^0-9]+") // regex for removing non digits
	if err != nil {
		fmt.Println(err)
	}
	// loop through every line
	for _, s := range arrInput {
		fmt.Printf("%s --> ", s)
		// remove all non digit characters
		s = reg.ReplaceAllString(s, "")
		// get the first and last character of a string
		s = string(s[0]) + string(s[len(s)-1])
		// convert string to integer
		d, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d\n", d)
		// append the digit to an array
		numbers = append(numbers, d)
	}
	fmt.Printf("Done!\n")
	return numbers
}

// sums all the numbers in the array
func sumNumbers(numbers *[]int) (result int) {
	fmt.Printf("Summing the array...\n%d\n", numbers)
	for _, d := range *numbers {
		result += d
	}
	fmt.Printf("Done!\n")
	return result
}

func main() {
	// read the input
	rawInput, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fixInput(&rawInput)
	replaceNumbers(&rawInput)
	numbers := getNumbers(&rawInput)
	result := sumNumbers(&numbers)
	// print out the result
	fmt.Printf("Here is the result: %d\n", result)
}
