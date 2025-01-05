package day1

import (
	"testing"
)

func TestDay1Part1Example(t *testing.T) {
	result := day1Part1("test_input.txt")
	expectedResult := 11
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay1Part2Example(t *testing.T) {
	result := day1Part2("test_input.txt")
	expectedResult := 31
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay1Part1(t *testing.T) {
	result := day1Part1("input.txt")
	expectedResult := 2192892
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay1Part2(t *testing.T) {
	result := day1Part2("input.txt")
	expectedResult := 22962826
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
