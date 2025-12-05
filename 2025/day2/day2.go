package day2

import (
	"fmt"
	"strings"

	"github.com/bpinheiro19/advent-of-code/utils"
)

func day2Part1(filename string) int {
	result := 0
	strList := utils.GetStringListFromCSV(filename)

	for _, s := range strList {
		sList := strings.Split(s, "-")

		for i := utils.StringToInt(sList[0]); i <= utils.StringToInt(sList[1]); i++ {
			str := utils.IntToString(i)

			if str[:len(str)/2] == str[len(str)/2:] {
				result += i
			}
		}
	}
	return result
}

func day2Part2(filename string) int {
	result := 0
	strList := utils.GetStringListFromCSV(filename)

	for _, s := range strList {
		sList := strings.Split(s, "-")

		for i := utils.StringToInt(sList[0]); i <= utils.StringToInt(sList[1]); i++ {
			str := utils.IntToString(i)

			for j := 1; j <= len(str)/2; j++ {

				if len(str)%j == 0 {

					valid := true
					for a := 0; a < len(str)-j; a += j {
						if str[a:a+j] != str[a+j:a+j+j] {
							valid = false
							break
						}
					}
					if valid {
						result += i
						break
					}
				}
			}
		}
	}
	return result
}

func Run(filename string) {
	fmt.Println("Day2 Part1 Result:", day2Part1(filename))
	fmt.Println("Day2 Part2 Result:", day2Part2(filename))
}
