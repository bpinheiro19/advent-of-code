package day10

import (
	"testing"
)

const FILENAME string = "input.txt"
const TEST_FILENAME string = "test_input.txt"

func TestDay10Part1Example(t *testing.T) {
	result := day10Part1(TEST_FILENAME)
	expectedResult := 36
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay10Part2Example(t *testing.T) {
	result := day10Part2(TEST_FILENAME)
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay10Part1(t *testing.T) {
	result := day10Part1(FILENAME)
	expectedResult := 820
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay10Part2(t *testing.T) {
	result := day10Part2(FILENAME)
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
