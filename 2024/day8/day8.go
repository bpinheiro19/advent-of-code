package day8

import (
	"fmt"

	"github.com/bpinheiro19/advent-of-code/utils"
)

func createCoordinatesMapMap(list []string) map[byte][][]int {

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
	coordinatesMap := createCoordinatesMapMap(list)

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

				newX1, newX2, newY1, newY2 := getAntiNodeXYCoordinates(coordinatesArr[n], x, y, 1)

				if utils.InBounds(newX1, newY1, listSize, listSize) && list[newX1][newY1] != '#' {
					list[newX1] = utils.ReplaceByteInString([]byte(list[newX1]), newY1, 35)
				}

				if utils.InBounds(newX2, newY2, listSize, listSize) && list[newX2][newY2] != '#' {
					list[newX2] = utils.ReplaceByteInString([]byte(list[newX2]), newY2, 35)
				}
			}
		}
	}
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

func calculateDistanceBetweenNodes(x, y, x2, y2 int) (int, int) {
	return utils.Abs(x - x2), utils.Abs(y - y2)
}

func getAntiNodeXYCoordinates(coordinatesArr []int, x, y int, multiplier int) (int, int, int, int) {
	x2, y2 := coordinatesArr[0], coordinatesArr[1]

	dX, dY := calculateDistanceBetweenNodes(x, y, x2, y2)
	dX, dY = dX*multiplier, dY*multiplier

	newX1, newX2, newY1, newY2 := 0, 0, 0, 0

	if x > x2 {
		newX1, newX2 = x2-dX, x+dX
	} else {
		newX1, newX2 = x2+dX, x-dX
	}

	if y > y2 {
		newY1, newY2 = y2-dY, y+dY
	} else {
		newY1, newY2 = y2+dY, y-dY
	}

	return newX1, newX2, newY1, newY2
}

func createAntiNodesPart2(list []string, coordinatesArr [][]int) {
	arrSize := len(coordinatesArr)
	listSize := len(list)

	for i := 0; i < arrSize; i++ {

		x, y := coordinatesArr[i][0], coordinatesArr[i][1]

		for n := 0; n < arrSize; n++ {

			if i != n {

				counter := 1
				newX1, newX2, newY1, newY2 := getAntiNodeXYCoordinates(coordinatesArr[n], x, y, counter)

				for utils.InBounds(newX1, newY1, listSize, listSize) || utils.InBounds(newX2, newY2, listSize, listSize) {

					if utils.InBounds(newX1, newY1, listSize, listSize) && list[newX1][newY1] != '#' {
						list[newX1] = utils.ReplaceByteInString([]byte(list[newX1]), newY1, 35)
					}

					if utils.InBounds(newX2, newY2, listSize, listSize) && list[newX2][newY2] != '#' {
						list[newX2] = utils.ReplaceByteInString([]byte(list[newX2]), newY2, 35)
					}

					counter++
					newX1, newX2, newY1, newY2 = getAntiNodeXYCoordinates(coordinatesArr[n], x, y, counter)
				}
			}
		}
	}
}

func markEachMissingAntenaAsAntiNode(list []string, coordinatesMap map[byte][][]int) {
	for i := 0; i < len(list); i++ {
		for n := 0; n < len(list[i]); n++ {
			if list[i][n] != '.' && list[i][n] != '#' && len(coordinatesMap[list[i][n]]) > 1 {
				list[i] = utils.ReplaceByteInString([]byte(list[i]), n, 35)
			}
		}
	}
}

func day8Part2(filename string) int {

	list := utils.GetStringListFromFile(filename)
	coordinatesMap := createCoordinatesMapMap(list)

	for _, coordinatesArr := range coordinatesMap {
		createAntiNodesPart2(list, coordinatesArr)
	}

	markEachMissingAntenaAsAntiNode(list, coordinatesMap)

	return findTotalAntiNodes(list)
}

func Run(filename string) {
	fmt.Println("Day8 Part1 Result:", day8Part1(filename))
	fmt.Println("Day8 Part2 Result:", day8Part2(filename))
}
