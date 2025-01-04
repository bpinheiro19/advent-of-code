package day2

import (
	"aoc2024/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getListFromFile(file string) [][]string {

	list := make([][]string, 0)

	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("open file error: %v", err)
	}

	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := strings.Fields(sc.Text())
		list = append(list, line)
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}

	return list
}

func isReportSafe(line []string) bool {

	safe := true
	increasing := true
	decreasing := true

	for i := 0; i < len(line)-1; i++ {

		num1 := utils.StringToInt(line[i])
		num2 := utils.StringToInt(string(line[i+1]))

		if utils.Abs(num2-num1) > 3 {
			safe = false
			break
		}

		if num1 == num2 {
			safe = false
			break

		} else if num1 < num2 {
			decreasing = false

			if !increasing {
				safe = false
				break
			}

		} else if num1 > num2 {
			increasing = false

			if !decreasing {
				safe = false
				break
			}
		}
	}
	return safe
}

func day2Part1(filename string) int {
	result := 0
	list := getListFromFile(filename)

	for _, line := range list {

		safe := isReportSafe(line)

		if safe {
			result++
		}
	}

	return result
}

func day2Part2(filename string) int {
	result := 0
	list := getListFromFile(filename)

	for _, line := range list {

		safe := isReportSafe(line)

		if safe {
			result++
		} else if tryProblemDampener(line) {
			result++
		}
	}

	return result
}

func tryProblemDampener(line []string) bool {
	safe := false
	for i := 0; i < len(line); i++ {
		newLine := make([]string, len(line))
		copy(newLine, line)
		safe = isReportSafe(append(newLine[:i], newLine[i+1:]...))
		if safe {
			break
		}
	}
	return safe
}

func Run(filename string) {
	fmt.Println("Day2 Part1 Result:", day2Part1(filename))
	fmt.Println("Day2 Part2 Result:", day2Part2(filename))
}
