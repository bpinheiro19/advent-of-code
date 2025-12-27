package day6

import (
	"testing"
)

func TestDay6Part1Example(t *testing.T) {
	result := day6Part1("test_input.txt")
	expectedResult := 4277556
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay6Part2Example(t *testing.T) {
	result := day6Part2("test_input.txt")
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay6Part1(t *testing.T) {
	result := day6Part1("input.txt")
	expectedResult := 5877594983578
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay6Part2(t *testing.T) {
	result := day6Part2("input.txt")
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
