package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

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

		num1, err := strconv.Atoi(line[i])

		if err != nil {
			log.Fatal(err)
		}

		num2, err := strconv.Atoi(line[i+1])

		if err != nil {
			log.Fatal(err)
		}

		if abs(num2-num1) > 3 {
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

func day2Part1() int {

	result := 0

	list := getListFromFile("input.txt")

	for _, line := range list {

		safe := isReportSafe(line)

		if safe {
			result++
		}
	}

	return result
}

func day2Part2() int {

	result := 0

	list := getListFromFile("input.txt")

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

func main() {

	fmt.Println("Day2 Part1 Result:", day2Part1())

	fmt.Println("Day2 Part2 Result:", day2Part2())

}
