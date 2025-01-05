package day5

import (
	"testing"
)

const FILENAME string = "input.txt"
const TEST_FILENAME string = "test_input.txt"

func TestDay5Part1Example(t *testing.T) {
	result, _ := day5(TEST_FILENAME)
	expectedResult := 143
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay5Part2Example(t *testing.T) {
	_, result := day5(TEST_FILENAME)
	expectedResult := 123
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay5Part1(t *testing.T) {
	result, _ := day5(FILENAME)
	expectedResult := 5166
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay5Part2(t *testing.T) {
	_, result := day5(FILENAME)
	expectedResult := 4679
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
