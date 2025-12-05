package day2

import (
	"testing"
)

func TestDay2Part1Example(t *testing.T) {
	result := day2Part1("test_input.txt")
	expectedResult := 1227775554
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay2Part2Example(t *testing.T) {
	result := day2Part2("test_input.txt")
	expectedResult := 4174379265
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay2Part1(t *testing.T) {
	result := day2Part1("input.txt")
	expectedResult := 17077011375
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}

func TestDay2Part2(t *testing.T) {
	result := day2Part2("input.txt")
	expectedResult := 36037497037
	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, expectedResult)
	}
}
