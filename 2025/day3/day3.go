package day3

import (
	"fmt"

	"github.com/bpinheiro19/advent-of-code/utils"
)

func day3Part1(filename string) int {
	result := 0

	strInput := utils.GetStringListFromFile(filename)

	for _, s := range strInput {

		a, b := 0, 0
		for i := 0; i < len(s); i++ {

			if int(s[i]-48) > a && i < len(s)-1 {
				a = int(s[i] - 48)
				b = 0

			} else if int(s[i]-48) > b {
				b = int(s[i] - 48)
			}
		}
		result += a*10 + b
	}
	return result
}

func day3Part2(filename string) int {
	result := 0
	return result
}

func Run(filename string) {
	fmt.Println("Day3 Part1 Result:", day3Part1(filename))
	fmt.Println("Day3 Part2 Result:", day3Part2(filename))
}
