package day10

import (
	"testing"
)

const FILENAME string = "input.txt"
const TEST_FILENAME string = "test_input.txt"

func TestDay10Part1Example(t *testing.T) {
	result, _ := day10(TEST_FILENAME)
	expectedResult := 36
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay10Part2Example(t *testing.T) {
	_, result := day10(TEST_FILENAME)
	expectedResult := 81
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay10Part1(t *testing.T) {
	result, _ := day10(FILENAME)
	expectedResult := 820
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay10Part2(t *testing.T) {
	_, result := day10(FILENAME)
	expectedResult := 1786
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
