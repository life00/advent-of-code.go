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
// returns game ID and true if possible, or 0 and false if not
// second value is for counting how many games it was possible in total
func checkGame(game string) (int, bool) {
	// get the game ID
	id := getGameID(game)
	// get the array of rounds
	rounds := getRounds(game)
	// loop through each round
	for _, r := range rounds {
		// check if its not possible
		if !checkRound(r) {
			fmt.Printf("Game %d is not possible.\n", id)
			// return 0 if not possible
			return 0, false
		}
	}
	fmt.Printf("Game %d is possible.\n", id)
	// return ID if possible
	return id, true
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
// true: possible
// false: not possible
func checkRound(round string) bool {
	// separate the round into individual cube values
	cubes := strings.Split(round, ", ")
	// get cube map
	cubeMap := getMap(cubes)
	// loop through the cubeMap
	for c, n := range cubeMap {
		// check if the number of cubes is not above the limit
		// if its above the limit return false
		// otherwise true
		switch c {
		case "red":
			if n > maxRed {
				return false
			}
		case "green":
			if n > maxGreen {
				return false
			}
		case "blue":
			if n > maxBlue {
				return false
			}
		}
	}
	return true
}

// gets the individual values in the round
func getMap(cubes []string) (cubeMap map[string]int) {
	// initialize the map (set to nothing)
	cubeMap = map[string]int{}
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

	// variable to sum all possible games IDs
	var idSum int
	// array of games that are possible
	var possibleGames []bool

	// loop through all the games
	for _, v := range games {
		// if the game is impossible it will return 0, 0
		// otherwise it will return id and 1
		id, possible := checkGame(v)
		// summing IDs
		idSum = idSum + id
		// appending the status
		possibleGames = append(possibleGames, possible)
	}

	// number of possible games
	var possible int

	// calculating number of possible games
	for _, p := range possibleGames {
		if p {
			possible += 1
		}
	}
	fmt.Printf("%d / %d games were possible.\n", possible, len(games))
	fmt.Printf("Sum of all possible game IDs: %d\n", idSum)
}
