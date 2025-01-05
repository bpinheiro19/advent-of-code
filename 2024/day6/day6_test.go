package day6

import (
	"testing"
)

const FILENAME string = "input.txt"
const TEST_FILENAME string = "test_input.txt"

func TestDay6Part1Example(t *testing.T) {
	result := day6Part1(TEST_FILENAME)
	expectedResult := 41
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay6Part2Example(t *testing.T) {
	result := day6Part2(TEST_FILENAME)
	expectedResult := 6
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay6Part1(t *testing.T) {
	result := day6Part1(FILENAME)
	expectedResult := 4454
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay6Part2(t *testing.T) {
	result := day6Part2(FILENAME)
	expectedResult := 1503
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
