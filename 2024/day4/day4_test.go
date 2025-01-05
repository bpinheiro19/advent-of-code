package day4

import (
	"testing"
)

const FILENAME string = "input.txt"
const TEST_FILENAME string = "test_input.txt"

func TestDay4Part1Example(t *testing.T) {
	result := day4Part1(TEST_FILENAME)
	expectedResult := 18
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay4Part2Example(t *testing.T) {
	result := day4Part2(TEST_FILENAME)
	expectedResult := 9
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay4Part1(t *testing.T) {
	result := day4Part1(FILENAME)
	expectedResult := 2591
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay4Part2(t *testing.T) {
	result := day4Part2(FILENAME)
	expectedResult := 1880
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
