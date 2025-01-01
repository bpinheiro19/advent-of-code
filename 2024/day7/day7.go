package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getStringListFromFile() []string {

	input, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	return strings.Fields(string(input))
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func createIntList() [][]int {
	list := getStringListFromFile()

	newList := make([][]int, 0)
	tempList := make([]int, 0)
	for _, str := range list {

		if str[len(str)-1] == ':' {
			if len(tempList) > 0 {
				newList = append(newList, tempList)
			}
			tempList = make([]int, 0)
			tempList = append(tempList, stringToInt(str[0:len(str)-1]))

		} else {
			tempList = append(tempList, stringToInt(str))
		}

	}
	newList = append(newList, tempList)

	if len(newList) == 0 {
		log.Fatal("Unable to create int list")
	}

	return newList
}

func day7Part1() int {
	result := 0

	list := createIntList()

	for i := 0; i < len(list); i++ {
		total := list[i][0]

		if canBeSolved(list[i], total) {
			result += total
		}
	}

	return result
}

func tryToSolve(a int, b []int, total int) int {

	if len(b) == 0 {
		return a
	}

	sum := tryToSolve(a+b[0], b[1:], total)
	sum2 := tryToSolve(a*b[0], b[1:], total)

	if sum == total {
		return sum
	}

	if sum2 == total {
		return sum2
	}

	return 0
}

func canBeSolved(list []int, total int) bool {
	return tryToSolve(list[1], list[2:], total) == total
}

func day7Part2() int {
	result := 0

	return result
}

func main() {

	fmt.Println("Day7 Part1 Result:", day7Part1())
	fmt.Println("Day7 Part2 Result:", day7Part2())

}
