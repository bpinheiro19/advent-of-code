package day7

import (
	"testing"
)

const FILENAME string = "input.txt"
const TEST_FILENAME string = "test_input.txt"

func TestDay7Part1Example(t *testing.T) {
	result := day7Part1(TEST_FILENAME)
	expectedResult := 3749
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay7Part2Example(t *testing.T) {
	result := day7Part2(TEST_FILENAME)
	expectedResult := 11387
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay7Part1(t *testing.T) {
	result := day7Part1(FILENAME)
	expectedResult := 14711933466277
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay7Part2(t *testing.T) {
	result := day7Part2(FILENAME)
	expectedResult := 286580387663654
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
