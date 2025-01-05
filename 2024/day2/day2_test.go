package day2

import (
	"testing"
)

const FILENAME string = "input.txt"
const TEST_FILENAME string = "test_input.txt"

func TestDay2Part1Example(t *testing.T) {
	result := day2Part1(TEST_FILENAME)
	expectedResult := 2
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay2Part2Example(t *testing.T) {
	result := day2Part2(TEST_FILENAME)
	expectedResult := 4
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay2Part1(t *testing.T) {
	result := day2Part1(FILENAME)
	expectedResult := 383
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay2Part2(t *testing.T) {
	result := day2Part2(FILENAME)
	expectedResult := 436
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
