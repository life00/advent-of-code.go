package main

import (
	"fmt"
	"testing"
)

// unit tests

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

// fuzz testing

func FuzzGetMaxColor(f *testing.F) {
	input := "Game 1: 4 green; 10 blue, 2 red, 7 green; 3 blue, 1 red, 0 green; 8 blue, 5 red"
	f.Add(input)
	f.Fuzz(func(t *testing.T, input string) {
		GetMaxColor(input)
	})
}

func FuzzGetGameId(f *testing.F) {
	input := "Game 1: 4 green; 10 blue, 2 red, 7 green; 3 blue, 1 red, 0 green; 8 blue, 5 red"
	f.Add(input)
	f.Fuzz(func(t *testing.T, input string) {
		GetGameId(input)
	})
}

func FuzzGetRounds(f *testing.F) {
	input := "Game 1: 4 green; 10 blue, 2 red, 7 green; 3 blue, 1 red, 0 green; 8 blue, 5 red"
	f.Add(input)
	f.Fuzz(func(t *testing.T, input string) {
		GetRounds(input)
	})
}

func FuzzGetColors(f *testing.F) {
	input := "10 blue, 2 red, 7 green"
	f.Add(input)
	f.Fuzz(func(t *testing.T, input string) {
		GetColors(input)
	})
}

func FuzzGetPower(f *testing.F) {
	f.Add(2, 4, 3)
	f.Fuzz(func(t *testing.T, i1 int, i2 int, i3 int) {
		GetPower(map[string]int{"red": i1, "green": i2, "blue": i3})
	})
}

func FuzzIsGamePossible(f *testing.F) {
	f.Add(2, 4, 3)
	f.Fuzz(func(t *testing.T, i1 int, i2 int, i3 int) {
		IsGamePossible(map[string]int{"red": i1, "green": i2, "blue": i3})
	})
}

// benchmarks

func BenchmarkGetMaxColor(b *testing.B) {
	input := "Game 1: 4 green; 10 blue, 2 red, 7 green; 3 blue, 1 red, 0 green; 8 blue, 5 red"
	for i := 0; i < b.N; i++ {
		GetMaxColor(input)
	}
}

func BenchmarkGetGameId(b *testing.B) {
	input := "Game 1: 4 green; 10 blue, 2 red, 7 green; 3 blue, 1 red, 0 green; 8 blue, 5 red"
	for i := 0; i < b.N; i++ {
		GetGameId(input)
	}
}

func BenchmarkGetRounds(b *testing.B) {
	input := "Game 1: 4 green; 10 blue, 2 red, 7 green; 3 blue, 1 red, 0 green; 8 blue, 5 red"
	for i := 0; i < b.N; i++ {
		GetRounds(input)
	}
}

func BenchmarkGetColors(b *testing.B) {
	input := "10 blue, 2 red, 7 green"
	for i := 0; i < b.N; i++ {
		GetColors(input)
	}
}

func BenchmarkGetPower(b *testing.B) {
	input := map[string]int{"red": 2, "green": 4, "blue": 3}
	for i := 0; i < b.N; i++ {
		GetPower(input)
	}
}

func BenchmarkIsGamePossible(b *testing.B) {
	input := map[string]int{"red": 2, "green": 4, "blue": 3}
	for i := 0; i < b.N; i++ {
		IsGamePossible(input)
	}
}
