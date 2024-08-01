package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// maximum limits for number of cubes
const (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

// gets the maximum values of colors in the game
func GetMaxColor(game string) (maxColorMap map[string]int) {
	// get the array of rounds
	rounds := GetRounds(game)
	// declare arrays of colors
	var reds, blues, greens []int
	// loop through each round
	for _, round := range rounds {
		// get the colors and append them
		colorMap := GetColors(round)
		reds = append(reds, colorMap["red"])
		blues = append(blues, colorMap["blue"])
		greens = append(greens, colorMap["green"])
	}
	maxColorMap = map[string]int{}
	// get maximum color values
	maxColorMap["red"] = slices.Max(reds)
	maxColorMap["green"] = slices.Max(greens)
	maxColorMap["blue"] = slices.Max(blues)
	// return maximum values
	return maxColorMap
}

// gets the game ID
func GetGameId(game string) (id int) {
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
func GetRounds(game string) []string {
	// remove Game ID part
	_, game, _ = strings.Cut(game, ": ")
	// create an array based on ; character
	return strings.Split(game, "; ")
}

// gets the color values of the round
func GetColors(round string) (colorMap map[string]int) {
	// separate the round into individual color values
	colors := strings.Split(round, ", ")
	// initialize the map
	// have them initialized as zero by default
	colorMap = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	// loop through colors
	for _, c := range colors {
		// split the color value into number and color
		number, color, _ := strings.Cut(c, " ")
		// assign to the map
		var err error
		colorMap[color], err = strconv.Atoi(number)
		if err != nil {
			fmt.Println(err)
		}
	}
	return colorMap
}

// calculates the power of max color values
func GetPower(maxColorMap map[string]int) int {
	// if all values are 0 which may happen if parsing failed
	// then return 0. This is necessary because of the logic after that
	if maxColorMap["red"] == 0 && maxColorMap["green"] == 0 && maxColorMap["blue"] == 0 {
		return 0
	}
	// a value of colors may be zero if there was no cubes of that color at all
	// to prevent power becoming zero, replace those colors that have 0 with 1
	if maxColorMap["red"] == 0 {
		maxColorMap["red"] = 1
	}
	if maxColorMap["green"] == 0 {
		maxColorMap["green"] = 1
	}
	if maxColorMap["blue"] == 0 {
		maxColorMap["blue"] = 1
	}
	// multiply and return
	return maxColorMap["red"] * maxColorMap["green"] * maxColorMap["blue"]
}

func IsGamePossible(maxColorMap map[string]int) bool {
	// check if the number of cubes of each color is not above the limit
	// if its above the limit return false
	// otherwise true
	if maxColorMap["red"] > MaxRed {
		return false
	}
	if maxColorMap["green"] > MaxGreen {
		return false
	}
	if maxColorMap["blue"] > MaxBlue {
		return false
	}
	return true
}

func main() {
	// print out the maximum values
	fmt.Printf("Maximum number of cubes for the games: %d red, %d green, and %d blue.\n", MaxRed, MaxGreen, MaxBlue)
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

	// variable to sum all powers of max color numbers
	var powerSum int
	// variable to sum all the possible games
	var possibleSum int
	// variable to sum all the ID's of possible games
	var possibleIdSum int

	// loop through all the games
	for _, g := range games {
		// get the game ID for logging
		id := GetGameId(g)
		fmt.Printf("Working on game %d...\n", id)
		// get the map of max color values in the game
		maxColorMap := GetMaxColor(g)
		fmt.Printf("Maximum colors for game %d are: %d red, %d green, and %d blue\n", id, maxColorMap["red"], maxColorMap["green"], maxColorMap["blue"])
		// check if game is possible
		if IsGamePossible(maxColorMap) {
			fmt.Printf("Game %d is possible\n", id)
			// add values to the total sum
			possibleSum += 1
			possibleIdSum += id
		} else {
			fmt.Printf("Game %d is not possible\n", id)
		}
		// get the power of the game and add it to the sum
		power := GetPower(maxColorMap)
		fmt.Printf("Power of game %d is: %d\n", id, power)
		powerSum += power
	}
	// printing the final results
	fmt.Printf("%d / %d games are possible.\n", possibleSum, len(games))
	fmt.Printf("Sum of all possible game ID's: %d\n", possibleIdSum)
	fmt.Printf("Sum of all game powers is: %d\n", powerSum)
}
