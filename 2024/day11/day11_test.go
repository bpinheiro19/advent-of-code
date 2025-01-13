package day11

import (
	"testing"
)

const FILENAME string = "input.txt"
const TEST_FILENAME string = "test_input.txt"

func TestDay11Part1Example(t *testing.T) {
	result := day11Part1(TEST_FILENAME)
	expectedResult := 55312
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay11Part2Example(t *testing.T) {
	result := day11Part2(TEST_FILENAME)
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay11Part1(t *testing.T) {
	result := day11Part1(FILENAME)
	expectedResult := 194482
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay11Part2(t *testing.T) {
	result := day11Part2(FILENAME)
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
