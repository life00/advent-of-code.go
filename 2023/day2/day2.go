package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

// checks if game is possible
// returns game ID if possible, or 0 if not
func checkGame(game string) int {
	// get the game ID
	id := getGameID(game)
	fmt.Printf("Checking game %d...\n", id)
	// get the array of rounds
	rounds := getRounds(game)
	// loop through each round
	for _, r := range rounds {
		// check if its possible
		if checkRound(r) {
			// return ID if possible
			return id
		}
	}
	// return 0 if not possible
	return 0
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

// check if the round is possible
func checkRound(round string) bool {
	cubes := strings.Split(round, ", ")
	cubeMap := getMap(cubes)
	fmt.Println(cubeMap)
	// ...
	return false
}

// gets the individual values in the round
func getMap(cubes []string) (cubeMap map[string]string) {
	// initialize the map (set to nothing)
	cubeMap = map[string]string{}
	// loop through cubes
	for _, c := range cubes {
		// split the cube value into number and color
		number, color, _ := strings.Cut(c, " ")
		// assign to the map
		cubeMap[color] = number
	}
	return cubeMap
}

func main() {
	// print out the maximum values
	fmt.Printf("Maximum number of cubes for the games: %d red, %d green, and %d blue.\n", maxRed, maxGreen, maxBlue)
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

	// loop through all the games
	for _, v := range games {
		// add the game ID
		// if the game is impossible it will return 0
		sum += checkGame(v)
	}

	// print the result
	fmt.Printf("Sum of all possible game IDs: %d\n", sum)
}
