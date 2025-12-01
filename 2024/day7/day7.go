package day7

import (
	"fmt"
	"log"

	"github.com/bpinheiro19/advent-of-code/utils"
)

func createIntList(filename string) [][]int {
	list := utils.GetStringListFromFile(filename)

	newList := make([][]int, 0)
	tempList := make([]int, 0)
	for _, str := range list {

		if str[len(str)-1] == ':' {
			if len(tempList) > 0 {
				newList = append(newList, tempList)
			}
			tempList = make([]int, 0)
			tempList = append(tempList, utils.StringToInt(str[0:len(str)-1]))

		} else {
			tempList = append(tempList, utils.StringToInt(str))
		}

	}
	newList = append(newList, tempList)

	if len(newList) == 0 {
		log.Fatal("Unable to create int list")
	}

	return newList
}

func day7Part1(filename string) int {
	result := 0

	list := createIntList(filename)

	for i := 0; i < len(list); i++ {
		total := list[i][0]

		if canBeSolvedPart1(list[i], total) {
			result += total
		}
	}

	return result
}

func tryToSolvePart1(a int, b []int, total int) int {

	if len(b) == 0 {
		return a
	}

	if a > total {
		return 0
	}

	sum := tryToSolvePart1(a+b[0], b[1:], total)
	sum2 := tryToSolvePart1(a*b[0], b[1:], total)

	if sum == total {
		return sum
	}

	if sum2 == total {
		return sum2
	}

	return 0
}

func canBeSolvedPart1(list []int, total int) bool {
	return tryToSolvePart1(list[1], list[2:], total) == total
}

func day7Part2(filename string) int {
	result := 0

	list := createIntList(filename)

	for i := 0; i < len(list); i++ {
		total := list[i][0]

		if canBeSolvedPart2(list[i], total) {
			result += total
		}
	}

	return result
}

func tryToSolvePart2(a int, b []int, total int) int {

	if len(b) == 0 {
		return a
	}

	if a > total {
		return 0
	}

	sum := tryToSolvePart2(a+b[0], b[1:], total)
	sum2 := tryToSolvePart2(a*b[0], b[1:], total)
	sum3 := tryToSolvePart2(utils.ConcatenateTwoInts(a, b[0]), b[1:], total)

	if sum == total {
		return sum
	}

	if sum2 == total {
		return sum2
	}

	if sum3 == total {
		return sum3
	}

	return 0
}

func canBeSolvedPart2(list []int, total int) bool {
	return tryToSolvePart2(list[1], list[2:], total) == total
}

func Run(filename string) {
	fmt.Println("Day7 Part1 Result:", day7Part1(filename))
	fmt.Println("Day7 Part2 Result:", day7Part2(filename))
}
