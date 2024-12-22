package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getStringFromFile(file string) []string {

	input, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return strings.Fields(string(input))
}

func checkRegEx(input string, regex string) []string {
	pattern := regexp.MustCompile(regex)
	return pattern.FindAllString(input, -1)
}

func day3Part1() int {
	result := 0

	list := getStringFromFile("input.txt")

	for _, line := range list {
		mulList := checkRegEx(line, `(mul\(\d*,\d*\))`)

		for _, mul := range mulList {
			sumList := checkRegEx(mul, `(\d*\,\d*)`)

			for _, val := range sumList {
				splitSumList := strings.Split(val, ",")

				num1, err := strconv.Atoi(splitSumList[0])

				if err != nil {
					log.Fatal(err)
				}

				num2, err := strconv.Atoi(splitSumList[1])

				if err != nil {
					log.Fatal(err)
				}

				result += num1 * num2
			}

		}

	}

	return result
}

func day3Part2() int {
	result := 0

	return result
}

func main() {

	fmt.Println("Day3 Part1 Result:", day3Part1())

	fmt.Println("Day3 Part2 Result:", day3Part2())

}
