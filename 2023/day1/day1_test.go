package main

import (
	"testing"
)

// unit tests

func TestFixInput(t *testing.T) {
	input := `WIkuseveniney2xl10tScRg1D
zntxAo5xz
reSNZYMggZu
qIfeOFfV26bdbtwonez`
	expected := `WIkusevenniney2xl10tScRg1D
zntxAo5xz
reSNZYMggZu
qIfeOFfV26bdbtwoonez`
	result := FixInput(input)
	if result != expected {
		t.Errorf("return value from FixInput did not match expected value\n")
	}
}

func TestReplaceNumbers(t *testing.T) {
	input := `WIkusevenniney2xl10tScRg1D
zntxAo5xz
reSNZYMggZu
qIfeOFfV26bdbtwoonez`
	expected := `WIku79y2xl10tScRg1D
zntxAo5xz
reSNZYMggZu
qIfeOFfV26bdb21z`
	result := ReplaceNumbers(input)
	if result != expected {
		t.Errorf("return value from ReplaceNumbers did not match expected value\n")
	}
}

func TestGetNumbers(t *testing.T) {
	input := `WIku79y2xl10tScRg1D
zntxAo5xz
reSNZYMggZu
qIfeOFfV26bdb21z
` // additional newline is necessary to account for last newline that is usually in input
	expected := []int{71, 55, 21}
	result := GetNumbers(input)
	if len(result) != len(expected) {
		t.Error()
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Error()
		}
	}
}

func TestSumNumbers(t *testing.T) {
	input := []int{71, 55, 21}
	expected := 147
	result := SumNumbers(input)
	if result != expected {
		t.Errorf("value of integer result did not match expected value\n")
	}
}

// fuzz testing

func FuzzFixInput(f *testing.F) {
	input := `WIkuseveniney2xl10tScRg1D
zntxAo5xz
reSNZYMggZu
qIfeOFfV26bdbtwonez`
	f.Add(input)
	f.Fuzz(func(t *testing.T, input string) {
		FixInput(input)
	})
}

func FuzzReplaceNumbers(f *testing.F) {
	input := `WIkusevenniney2xl10tScRg1D
zntxAo5xz
reSNZYMggZu
qIfeOFfV26bdbtwoonez`
	f.Add(input)
	f.Fuzz(func(t *testing.T, input string) {
		ReplaceNumbers(input)
	})
}

func FuzzGetNumbers(f *testing.F) {
	input := `WIku79y2xl10tScRg1D
zntxAo5xz
reSNZYMggZu
qIfeOFfV26bdb21z
` // additional newline is necessary to account for last newline that is usually in input
	f.Add(input)
	f.Fuzz(func(t *testing.T, input string) {
		GetNumbers(input)
	})
}

func FuzzSumNumbers(f *testing.F) {
	f.Add(71, 55, 21)
	f.Fuzz(func(t *testing.T, i1 int, i2 int, i3 int) {
		SumNumbers([]int{i1, i2, i3})
	})
}

// benchmarks

func BenchmarkFixInput(b *testing.B) {
	input := `WIkuseveniney2xl10tScRg1D
zntxAo5xz
reSNZYMggZu
qIfeOFfV26bdbtwonez`
	for i := 0; i < b.N; i++ {
		FixInput(input)
	}
}

func BenchmarkReplaceNumbers(b *testing.B) {
	input := `WIkusevenniney2xl10tScRg1D
zntxAo5xz
reSNZYMggZu
qIfeOFfV26bdbtwoonez`
	for i := 0; i < b.N; i++ {
		ReplaceNumbers(input)
	}
}

func BenchmarkGetNumbers(b *testing.B) {
	input := `WIku79y2xl10tScRg1D
zntxAo5xz
reSNZYMggZu
qIfeOFfV26bdb21z
` // additional newline is necessary to account for last newline that is usually in input
	for i := 0; i < b.N; i++ {
		GetNumbers(input)
	}
}

func BenchmarkSumNumbers(b *testing.B) {
	input := []int{71, 55, 21}
	for i := 0; i < b.N; i++ {
		SumNumbers(input)
	}
}
