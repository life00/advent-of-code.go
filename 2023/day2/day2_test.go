package main

import (
	"fmt"
	"testing"
)

func TestGetMaxColor(t *testing.T) {
	tests := []struct {
		expected map[string]int
		input    string
	}{
		{
			map[string]int{
				"red":   5,
				"green": 7,
				"blue":  10,
			},
			"Game 1: 4 green; 10 blue, 2 red, 7 green; 3 blue, 1 red, 0 green; 8 blue, 5 red",
		},
		{
			map[string]int{
				"red":   0,
				"green": 3,
				"blue":  15,
			},
			"Game 2: 3 green, 15 blue",
		},
		{
			map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			},
			"Game 3: corrupted game input",
		},
		{
			map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			},
			"corrupted game input",
		},
	}
	for i, test := range tests {
		name := fmt.Sprintf("game %d", i+1)
		t.Run(name, func(t *testing.T) {
			result := GetMaxColor(test.input)
			if len(result) != len(test.expected) {
				t.Errorf("length of map maxColorMap did not match expected length\n")
			}
			for j := range result {
				if result[j] != test.expected[j] {
					t.Errorf("value of key %s in map maxColorMap did not match expected value\n", j)
				}
			}
		})
	}
}

func TestGetGameId(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"Game 1: 4 green; 10 blue, 2 red, 7 green; 3 blue, 1 red, 0 green; 8 blue, 5 red",
			1,
		},
		{
			"Game 2: 3 green, 15 blue",
			2,
		},
		{
			"Game 3: corrupted game input",
			3,
		},
		{
			"corrupted game input",
			0,
		},
	}
	for i, test := range tests {
		name := fmt.Sprintf("game %d", i+1)
		t.Run(name, func(t *testing.T) {
			result := GetGameId(test.input)
			if result != test.expected {
				t.Errorf("value of integer id did not match expected value\n")
			}
		})
	}
}

func TestGetRounds(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			"Game 1: 4 green; 10 blue, 2 red, 7 green; 3 blue, 1 red, 0 green; 8 blue, 5 red",
			[]string{"4 green", "10 blue, 2 red, 7 green", "3 blue, 1 red, 0 green", "8 blue, 5 red"},
		},
		{
			"Game 2: 3 green, 15 blue",
			[]string{"3 green, 15 blue"},
		},
		{
			"Game 3: corrupted game input",
			[]string{"corrupted game input"},
		},
		{
			"corrupted game input",
			[]string{""},
		},
	}
	for i, test := range tests {
		name := fmt.Sprintf("game %d", i+1)
		t.Run(name, func(t *testing.T) {
			result := GetRounds(test.input)
			if len(result) != len(test.expected) {
				t.Errorf("length of array rounds did not match expected length\n")
			}
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("value of index %d in array rounds did not match expected value\n", i)
				}
			}
		})
	}
}

func TestGetColors(t *testing.T) {
	tests := []struct {
		expected map[string]int
		input    string
	}{
		{
			map[string]int{"red": 0, "green": 4, "blue": 0},
			"4 green",
		},
		{
			map[string]int{"red": 1, "green": 0, "blue": 3},
			"3 blue, 1 red, 0 green",
		},
		{
			map[string]int{"red": 0, "green": 0, "blue": 0, "game input": 0},
			"corrupted game input",
		},
		{
			map[string]int{"red": 0, "green": 0, "blue": 0, "": 0},
			"", // corrupted game input
		},
	}
	for i, test := range tests {
		name := fmt.Sprintf("round %d", i+1)
		t.Run(name, func(t *testing.T) {
			result := GetColors(test.input)
			if len(result) != len(test.expected) {
				t.Errorf("length of map colorMap did not match expected length\n")
			}
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("value of key %s in map colorMap did not match expected value\n", i)
				}
			}
		})
	}
}

func TestGetPower(t *testing.T) {
	tests := []struct {
		input    map[string]int
		expected int
	}{
		{
			map[string]int{"red": 0, "green": 4, "blue": 0},
			4,
		},
		{
			map[string]int{"red": 1, "green": 0, "blue": 3},
			3,
		},
		{
			map[string]int{"red": 4, "green": 8, "blue": 3},
			96,
		},
		{
			map[string]int{"red": 0, "green": 0, "blue": 0, "corrupted value": 0},
			0,
		},
	}
	for i, test := range tests {
		name := fmt.Sprintf("game %d", i+1)
		t.Run(name, func(t *testing.T) {
			result := GetPower(test.input)
			if result != test.expected {
				t.Errorf("value of integer power did not match expected value\n")
			}
		})
	}
}

func TestIsGamePossible(t *testing.T) {
	tests := []struct {
		input    map[string]int
		expected bool
	}{
		{
			map[string]int{"red": 99, "green": 0, "blue": 0},
			false,
		},
		{
			map[string]int{"red": 0, "green": 99, "blue": 0},
			false,
		},
		{
			map[string]int{"red": 0, "green": 0, "blue": 99},
			false,
		},
		{
			map[string]int{"red": 0, "green": 0, "blue": 0},
			true,
		},
	}
	for i, test := range tests {
		name := fmt.Sprintf("game %d", i+1)
		t.Run(name, func(t *testing.T) {
			result := IsGamePossible(test.input)
			if result != test.expected {
				t.Errorf("value of bool possible did not match expected value\n")
			}
		})
	}
}
