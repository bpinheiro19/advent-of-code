package day5

import (
	"testing"
)

func TestDay5Part1Example(t *testing.T) {
	result := day5Part1("test_input.txt")
	expectedResult := 3
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay5Part2Example(t *testing.T) {
	result := day5Part2("test_input.txt")
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay5Part1(t *testing.T) {
	result := day5Part1("input.txt")
	expectedResult := 885
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay5Part2(t *testing.T) {
	result := day5Part2("input.txt")
	expectedResult := 0
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
