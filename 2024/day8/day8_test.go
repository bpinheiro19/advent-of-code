package day8

import (
	"testing"
)

const FILENAME string = "input.txt"
const TEST_FILENAME string = "test_input.txt"

func TestDay8Part1Example(t *testing.T) {
	result := day8Part1(TEST_FILENAME)
	expectedResult := 14
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay8Part2Example(t *testing.T) {
	result := day8Part2(TEST_FILENAME)
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay8Part1(t *testing.T) {
	result := day8Part1(FILENAME)
	expectedResult := 228
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay8Part2(t *testing.T) {
	result := day8Part2(FILENAME)
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
