package day8

import (
	"fmt"
	"slices"

	"aoc2024/utils"
)

func createHelperMap(list []string) map[byte][][]int {

	helperMap := make(map[byte][][]int)

	for i := 0; i < len(list); i++ {
		for n := 0; n < len(list[i]); n++ {

			if list[i][n] != '.' {
				helperMap[list[i][n]] = append(helperMap[list[i][n]], []int{i, n})
			}
		}
	}

	return helperMap
}

func day8Part1(filename string) int {

	list := utils.GetStringListFromFile(filename)
	coordinatesMap := createHelperMap(list)

	for _, coordinatesArr := range coordinatesMap {
		createAntiNodes(list, coordinatesArr)
	}

	return findTotalAntiNodes(list)
}

func createAntiNodes(list []string, coordinatesArr [][]int) {
	arrSize := len(coordinatesArr)
	listSize := len(list)

	for i := 0; i < arrSize; i++ {

		x, y := coordinatesArr[i][0], coordinatesArr[i][1]

		for n := 0; n < arrSize; n++ {

			if i != n {

				x2, y2 := coordinatesArr[n][0], coordinatesArr[n][1]

				newX, newY := utils.Abs(x-x2), utils.Abs(y-y2)

				newX1, newX2, newY1, newY2 := 0, 0, 0, 0

				if x > x2 {
					newX1, newX2 = x2-newX, x+newX
				} else {
					newX1, newX2 = x2+newX, x-newX
				}

				if y > y2 {
					newY1, newY2 = y2-newY, y+newY
				} else {
					newY1, newY2 = y2+newY, y-newY
				}

				if utils.InBounds(newX1, newY1, listSize, listSize) && list[newX1][newY1] != '#' {
					list[newX1] = replaceByteInString([]byte(list[newX1]), newY1, 35)
				}

				if utils.InBounds(newX2, newY2, listSize, listSize) && list[newX2][newY2] != '#' {
					list[newX2] = replaceByteInString([]byte(list[newX2]), newY2, 35)
				}
			}
		}
	}

}

func replaceByteInString(str []byte, i int, val byte) string {
	return string(slices.Replace(str, i, i+1, val))
}

func findTotalAntiNodes(list []string) int {
	antiNodes := 0
	for i := 0; i < len(list); i++ {
		for n := 0; n < len(list[i]); n++ {
			if list[i][n] == '#' {
				antiNodes++
			}
		}
	}
	return antiNodes
}

func day8Part2(filename string) int {
	result := 0
	return result
}

func Run(filename string) {
	fmt.Println("Day8 Part1 Result:", day8Part1(filename))
	fmt.Println("Day8 Part2 Result:", day8Part2(filename))
}
