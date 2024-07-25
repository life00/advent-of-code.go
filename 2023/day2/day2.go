package main

import (
	"fmt"
	"os"
	"strings"
)

func checkGame(raw *string) int {
	fmt.Printf("Checking: %s\n", *raw)
	// ...
	return 0
}

func main() {
	// read the input file
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	// split it into an array based on newline character
	games := strings.Split(string(bytes), "\n")
	// remove the last faulty element in the array
	if len(games) > 0 {
		games = games[:len(games)-1]
	}

	// variable to sum all possible games
	var sum int

	for _, v := range games {
		sum += checkGame(&v)
	}

	fmt.Println(sum)
}
