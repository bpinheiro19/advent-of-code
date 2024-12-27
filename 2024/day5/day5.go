package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

func createHelperMap() map[int][]int {

	list := getStringListFromFile()

	helperMap := make(map[int][]int)

	for _, str := range list {

		if strings.Contains(str, "|") {

			splitList := strings.Split(str, "|")

			num1 := stringToInt(splitList[0])
			num2 := stringToInt(splitList[1])

			helperMap[num2] = append(helperMap[num2], num1)
		}
	}
	return helperMap
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func day5Part1() int {
	result := 0

	list := getStringListFromFile()

	helperMap := createHelperMap()

	for _, str := range list {

		if strings.Contains(str, ",") {
			splitList := strings.Split(str, ",")

			numList := make([]int, 0)
			for i := range splitList {

				n := stringToInt(splitList[i])

				numList = append(numList, n)
			}

			if !checkOrderViolation(numList, helperMap) {
				result += numList[(len(numList) / 2)]
			}

		}

	}

	return result
}

func checkOrderViolation(numList []int, helperMap map[int][]int) bool {

	for i := 0; i < len(numList); i++ {
		for n := i + 1; n < len(numList); n++ {
			if isInArray(helperMap[numList[i]], numList[n]) {
				return true
			}
		}
	}

	return false
}

func isInArray(arr []int, i int) bool {
	return slices.Index(arr, i) != -1
}

func day5Part2() int {
	result := 0

	return result
}

func main() {

	fmt.Println("Day5 Part1 Result:", day5Part1())

	fmt.Println("Day5 Part2 Result:", day5Part2())

}
