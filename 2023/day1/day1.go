package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func replaceNumbers(rawInput *[]byte) {
	// temporarily save it as other variable
	input := string(*rawInput)
	// define strings and ints
	strings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	ints := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	// regex condition for loop
	regLoop, err := regexp.Compile("one|two|three|four|five|six|seven|eight|nine")
	if err != nil {
		fmt.Println(err)
	}
	for {
		// check for matches in input
		if match := regLoop.FindString(input); match != "" {
			// number that was matched
			var num string
			// identify the number that was matched
			for i, v := range strings {
				if v == match {
					num = ints[i]
				}
			}
			// replacing the match
			regReplace, err := regexp.Compile(match)
			if err != nil {
				fmt.Println(err)
			}
			// the following code replaces only the first match and saves as input variable
			flag := false
			input = regReplace.ReplaceAllStringFunc(input, func(s string) string {
				if flag {
					return s
				}
				flag = true
				return regReplace.ReplaceAllString(s, num)
			})
		} else {
			// when there is no more matches to substitute
			break
		}
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
