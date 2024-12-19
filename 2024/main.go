package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func getListsFromFile(file string) ([]int, []int) {

	list1 := make([]int, 1)
	list2 := make([]int, 1)

	input, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	strInput := strings.Fields(string(input))

	for i, v := range strInput {

		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		if i%2 == 0 {
			list1 = append(list1, num)
		} else {
			list2 = append(list2, num)
		}
	}

	return list1, list2
}

func day1Part1() int {
	result := 0

	list1, list2 := getListsFromFile("list.txt")

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	for i, val := range list1 {
		result += abs(val - list2[i])
	}

	return result
}

func day1Part2() int {
	result := 0

	list1, list2 := getListsFromFile("list.txt")

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

func main() {

	day1Part1Result := day1Part1()
	fmt.Println("Day1 Part1 Result:", day1Part1Result)

	day1Part2Result := day1Part2()
	fmt.Println("Day1 Part2 Result:", day1Part2Result)

}
