package day1

import (
	"aoc2024/utils"
	"fmt"
	"sort"
)

func createNumLists(filename string) ([]int, []int) {

	list1 := make([]int, 1)
	list2 := make([]int, 1)

	strInput := utils.GetStringListFromFile(filename)

	for i, v := range strInput {

		num := utils.StringToInt(v)

		if i%2 == 0 {
			list1 = append(list1, num)
		} else {
			list2 = append(list2, num)
		}
	}

	return list1, list2
}

func day1Part1(filename string) int {
	result := 0

	list1, list2 := createNumLists(filename)

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	for i, val := range list1 {
		result += utils.Abs(val - list2[i])
	}

	return result
}

func day1Part2(filename string) int {
	result := 0

	list1, list2 := createNumLists(filename)

	list2Map := make(map[int]int)

	for i := 0; i < 10; i++ {
		list2Map[i] = 0
	}

	for _, val := range list2 {
		list2Map[val]++
	}

	for _, val := range list1 {
		result += val * list2Map[val]
	}

	return result
}

func Run(filename string) {

	fmt.Println("Day1 Part1 Result:", day1Part1(filename))
	fmt.Println("Day1 Part2 Result:", day1Part2(filename))
}
