package day4

import (
	"testing"
)

func TestDay4Part1Example(t *testing.T) {
	result := day4Part1("test_input.txt")
	expectedResult := 13
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay4Part2Example(t *testing.T) {
	result := day4Part2("test_input.txt")
	expectedResult := 43
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay4Part1(t *testing.T) {
	result := day4Part1("input.txt")
	expectedResult := 1478
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay4Part2(t *testing.T) {
	result := day4Part2("input.txt")
	expectedResult := 9120
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
