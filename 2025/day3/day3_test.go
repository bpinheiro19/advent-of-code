package day3

import (
	"testing"
)

func TestDay3Part1Example(t *testing.T) {
	result := day3Part1("test_input.txt")
	expectedResult := 357
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay3Part2Example(t *testing.T) {
	result := day3Part2("test_input.txt")
	expectedResult := 3121910778619
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay3Part1(t *testing.T) {
	result := day3Part1("input.txt")
	expectedResult := 17554
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay3Part2(t *testing.T) {
	result := day3Part2("input.txt")
	expectedResult := 175053592950232
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
