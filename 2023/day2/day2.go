package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// gets the power of a set of cubes in the game
// it is the product of maximum red, green and blue cube values
func getPower(game string) (power int) {
	// get the game ID
	id := getGameID(game)
	// get the array of rounds
	rounds := getRounds(game)
	// declare arrays of colors
	var reds, blues, greens []int
	// loop through each round
	for _, round := range rounds {
		// get the colors and append them
		r, g, b := getColors(round)
		reds = append(reds, r)
		blues = append(blues, b)
		greens = append(greens, g)
	}
	// find maximum values for each of the colors
	maxRed := slices.Max(reds)
	maxBlue := slices.Max(blues)
	maxGreen := slices.Max(greens)
	// calculating power
	power = maxRed * maxBlue * maxGreen
	fmt.Printf("Power for %d game is: %d\n", id, power)
	return power
}

// gets the game ID
func getGameID(game string) (id int) {
	// remove the prefix and suffix of the game data
	game, _ = strings.CutPrefix(game, "Game ")
	game, _, _ = strings.Cut(game, ":")
	// convert to integer
	id, err := strconv.Atoi(game)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

// gets separate rounds in the game
func getRounds(game string) []string {
	// remove Game ID part
	_, game, _ = strings.Cut(game, ": ")
	// create an array based on ; character
	return strings.Split(game, "; ")
}

// gets the color values of the round
func getColors(round string) (int, int, int) {
	// separate the round into individual cube values
	cubes := strings.Split(round, ", ")
	// get cube map
	cubeMap := getMap(cubes)
	// return the red, green and blue values
	return cubeMap["red"], cubeMap["green"], cubeMap["blue"]
}

// gets the individual values in the round
func getMap(cubes []string) (cubeMap map[string]int) {
	// initialize the map
	// HACK: it is all set to 1 by default to prevent situations when
	// uninitialized map is returned with value 0 which would break
	// the power calculation. 1 is a neutral value in power calculation.
	cubeMap = map[string]int{
		"red":   1,
		"green": 1,
		"blue":  1,
	}
	// loop through cubes
	for _, c := range cubes {
		// split the cube value into number and color
		number, color, _ := strings.Cut(c, " ")
		// assign to the map
		var err error
		cubeMap[color], err = strconv.Atoi(number)
		if err != nil {
			fmt.Println(err)
		}
	}
	return cubeMap
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

	// variable to sum all powers of max cube numbers
	var powerSum int

	// loop through all the games
	for _, g := range games {
		// if the game is impossible it will return 0, 0
		// otherwise it will return id and 1
		powerSum += getPower(g)
	}
	fmt.Printf("The sum of all powers is: %d\n", powerSum)
}
