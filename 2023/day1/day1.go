package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// read the input
	rawInput, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Println(err)
	}
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
