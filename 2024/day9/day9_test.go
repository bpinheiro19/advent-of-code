package day9

import (
	"testing"
)

const FILENAME string = "input.txt"
const TEST_FILENAME string = "test_input.txt"

func TestDay9Part1Example(t *testing.T) {
	result := day9Part1(TEST_FILENAME)
	expectedResult := 1928
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay9Part2Example(t *testing.T) {
	result := day9Part2(TEST_FILENAME)
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay9Part1(t *testing.T) {
	result := day9Part1(FILENAME)
	expectedResult := 6346871685398
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay9Part2(t *testing.T) {
	result := day9Part2(FILENAME)
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
