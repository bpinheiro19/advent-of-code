package day3

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

func multiplyValues(mul string) int {
	result := 0
	sumList := utils.CheckRegEx(mul, `(\d*\,\d*)`)

	for _, val := range sumList {
		splitSumList := strings.Split(val, ",")

		num1 := utils.StringToInt(splitSumList[0])
		num2 := utils.StringToInt(splitSumList[1])

		result += num1 * num2
	}

	return result
}

func day3Part1(filename string) int {
	result := 0

	list := utils.GetStringListFromFile(filename)

	for _, line := range list {
		mulList := utils.CheckRegEx(line, `(mul\(\d*,\d*\))`)

		for _, mul := range mulList {
			result += multiplyValues(mul)
		}

	}

	return result
}

func day3Part2(filename string) int {
	result := 0
	mulEnabled := true

	list := utils.GetStringListFromFile(filename)

	for _, line := range list {

		firstList := utils.CheckRegEx(line, `(mul\(\d*,\d*\))|(do\(\))|(don't\(\))`)

		for _, val := range firstList {

			if val == "do()" {
				mulEnabled = true
			} else if val == "don't()" {
				mulEnabled = false
			} else {

				if mulEnabled {
					result += multiplyValues(val)
				}
			}
		}
	}

	return result
}

func Run(filename string) {
	fmt.Println("Day3 Part1 Result:", day3Part1(filename))
	fmt.Println("Day3 Part2 Result:", day3Part2(filename))
}
