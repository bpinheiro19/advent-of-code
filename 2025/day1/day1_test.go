package day1

import (
	"testing"
)

func TestDay1Part1Example(t *testing.T) {
	result := day1Part1("test_input.txt")
	expectedResult := 3
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay1Part2Example(t *testing.T) {
	result := day1Part2("test_input.txt")
	expectedResult := 6
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay1Part1(t *testing.T) {
	result := day1Part1("input.txt")
	expectedResult := 1195
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay1Part2(t *testing.T) {
	result := day1Part2("input.txt")
	expectedResult := 6770
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
