package main

import (
	"testing"
)

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
