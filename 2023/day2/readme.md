# [Day 2](https://adventofcode.com/2023/day/2)

## Part 1

### Summary

I have the following input with several different games and rounds within them:

```
Game 1: 1 red, 2 green, 3 blue; 5 blue, 3 red
Game 2: 4 green; 10 blue, 2 red, 7 green; 3 blue, 1 red, 15 green
Game 3: 3 green, 15 blue, 4 red; 3 blue, 8 red
```

There is a limit of how many red, green, and blue cubes could exist at once. The maximum limits of cubes are: `12` red, `13` green, and `14` blue. This means that some of the rounds, and thus the games would be impossible. For example games 2 and 3 would be impossible in the above example.

It is necessary to determine which of the games would be impossible, then sum all the IDs of games that are possible.

### Solution

- declare and set constants representing maximum values of cubes
- read and save the input file
- split the input into an array of separate games
- loop through the array
- check if the game is possible, with return as ID of the game (if not possible then it is 0)
  - get game ID
  - split the game input string into array of rounds
  - loop through rounds, with return as bool of whether the round is possible or not
    - separate the round input into individual cube values
    - create a cube map
    - check if the values are within the limits using a switch statement
    - return true if within the limit, otherwise false
  - return ID if the game is possible, return 0 if not
- sum the returns and output the result

## Part 2

### Summary

There is the same input, but now it is necessary to find the minimum limit of red, green, and blue cubes in a game (i.e. max number of each color cubes). Then multiply them by each other, and find a sum together with other games.

For example for the game 1:

```
Game 1: 1 red, 2 green, 3 blue; 5 blue, 3 red
```

It would be necessary to have at least 3 red, 2 green, and 5 blue. $3\times 2 \times 5=30$, and sum it with powers of other games.

### Solution

Majority of previous functionality can be reused like parsing, but I have rewritten the solution:

- read and save the input file
- split the input into an array of separate games
- loop through the array
- get the power of the game which is the product of maximum color values
  - get game ID
  - split the game input string into array of rounds
  - loop through rounds, get color values
    - separate the round input into individual cube values
    - create a cube map
    - return color values
  - find the maximum color value in color arrays
  - calculate the product of maximum color values
  - return the result
- sum the results of powers from all the games

## Refactor

In order to make both parts work I have refactored the code to do the following:

- read the input
- split the input into individual gamest (and loop through everything afterwards)
- get ID for logging
- get max cube number for each color in a game
- check if the game is possible
- calculate the power of the game
- print out the results

## Error handling

To not complicate things I did not attempt to create logic or regex that would automatically check if a string is valid or not. Instead I let `strconv.Atoi` fail on parsing, print the error, and continue. I believe that returning the error up the functions to stop the process is also unnecessary.

This approach does not change anything really because if it fails, the maximum values for the game would still all be 0. Thus it will not impact the end calculation at all. I have tested this by running the program on an image as input file. It may impact the result only if some part of the game string was successfully parsed, but this is an expected behavior.

The user is informed about the problem in the logs. I believe that making a fatal error here because of failed parsing is unnecessary.

## Testing

To test the program just run `go test -v` while in this directory. I have implemented the most essential tests for all functions except for main. Please ignore `strconv.Atoi` error messages, those are a part of testing of error handling.
