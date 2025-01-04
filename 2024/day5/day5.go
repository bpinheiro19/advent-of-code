package day5

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strings"
)

func createHelperMap(filename string) map[int][]int {

	list := utils.GetStringListFromFile(filename)

	helperMap := make(map[int][]int)

	for _, str := range list {

		if strings.Contains(str, "|") {

			splitList := strings.Split(str, "|")

			num1 := utils.StringToInt(splitList[0])
			num2 := utils.StringToInt(splitList[1])

			helperMap[num2] = append(helperMap[num2], num1)
		}
	}
	return helperMap
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

func insertIntoArray(arr []int, i int, e int) []int {
	return append(arr[:i], append([]int{e}, arr[i:]...)...)
}

func removeFromArray(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func day5(filename string) (int, int) {
	result1 := 0
	result2 := 0

	list := utils.GetStringListFromFile(filename)

	helperMap := createHelperMap(filename)

	for _, str := range list {

		if strings.Contains(str, ",") {
			splitList := strings.Split(str, ",")

			numList := make([]int, 0)
			for i := range splitList {

				n := utils.StringToInt(splitList[i])

				numList = append(numList, n)
			}

			if !checkOrderViolation(numList, helperMap) {
				result1 += numList[(len(numList) / 2)]
			} else {
				result2 += day5Part2(numList, helperMap)
			}

		}

	}

	return result1, result2
}

func day5Part2(numList []int, helperMap map[int][]int) int {
	result := 0

	for i := 0; i < len(numList); i++ {
		for n := i + 1; n < len(numList); n++ {
			if isInArray(helperMap[numList[i]], numList[n]) {
				e := numList[n]
				removeFromArray(numList, n)
				insertIntoArray(numList, i, e)
				i = -1
				n = 0
				break
			}
		}
	}
	result += numList[(len(numList) / 2)]

	return result
}

func Run(filename string) {
	part1, part2 := day5(filename)
	fmt.Println("Day5 Part1 Result:", part1)
	fmt.Println("Day5 Part2 Result:", part2)
}
