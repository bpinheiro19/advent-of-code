package day1

import (
	"fmt"

	"github.com/bpinheiro19/advent-of-code/utils"
)

func day1Part1(filename string) int {
	result := 0
	dial := 50

	strInput := utils.GetStringListFromFile(filename)

	for _, v := range strInput {
		num := utils.StringToInt(v[1:])

		switch v[0] {
		case 'L':
			dial -= num
			for dial < 0 {
				dial += 100
				result++
			}

		case 'R':
			dial += num
			for dial > 99 {
				dial -= 100
				result++
			}
		}

		if dial == 0 {
			result++
		}
	}
	return result
}

func day1Part2(filename string) int {
	result := 0

	return result
}

func Run(filename string) {
	fmt.Println("Day1 Part1 Result:", day1Part1(filename))
	fmt.Println("Day1 Part2 Result:", day1Part2(filename))
}
