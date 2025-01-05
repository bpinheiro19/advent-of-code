package day3

import (
	"testing"
)

const FILENAME string = "input.txt"
const TEST_FILENAME string = "test_input.txt"
const TEST_FILENAME2 string = "test_input2.txt"

func TestDay3Part1Example(t *testing.T) {
	result := day3Part1(TEST_FILENAME)
	expectedResult := 161
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay3Part2Example(t *testing.T) {
	result := day3Part2(TEST_FILENAME2)
	expectedResult := 48
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay3Part1(t *testing.T) {
	result := day3Part1(FILENAME)
	expectedResult := 165225049
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay3Part2(t *testing.T) {
	result := day3Part2(FILENAME)
	expectedResult := 108830766
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
