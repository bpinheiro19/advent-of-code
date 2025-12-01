package day10

import (
	"fmt"

	"github.com/bpinheiro19/advent-of-code/utils"
)

func createIntList(filename string) [][]int {
	byteList := utils.GetByteListFromFile(filename)
	intList := make([][]int, 0)
	tmpList := make([]int, 0)

	for i := 0; i < len(byteList); i++ {
		if byteList[i] == 10 {
			intList = append(intList, tmpList)
			tmpList = make([]int, 0)
		} else {
			tmpList = append(tmpList, int(byteList[i]-48))
		}

	}
	intList = append(intList, tmpList)
	return intList
}

func createPositionList(intList [][]int) ([][]int, [][]int) {

	zeroPosList := make([][]int, 0)
	ninePosList := make([][]int, 0)

	for i := 0; i < len(intList); i++ {
		for j := 0; j < len(intList[i]); j++ {
			if intList[i][j] == 0 {
				zeroPosList = append(zeroPosList, []int{i, j})
			}
			if intList[i][j] == 9 {
				ninePosList = append(ninePosList, []int{i, j})
			}
		}
	}
	return zeroPosList, ninePosList
}

func findTrailHeadScore(intList [][]int, x, y, w, z int) int {

	score := 0
	val := intList[x][y]
	length := len(intList)

	tailheadStack := make(map[int][][]int)
	adjacentList := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	if x == w && y == z {
		return 1
	}

	for _, al := range adjacentList {
		dX, dY := al[0], al[1]

		if utils.InBounds(x+dX, y+dY, length, length) && intList[x+dX][y+dY] == val+1 {
			tailheadStack[val] = append(tailheadStack[val], []int{x + dX, y + dY})
		}
	}

	newList := tailheadStack[val]
	for i := 0; i < len(newList); i++ {
		score += findTrailHeadScore(intList, newList[i][0], newList[i][1], w, z)

	}

	return score
}

func day10(filename string) (int, int) {
	resultPart1 := 0
	resultPart2 := 0

	intList := createIntList(filename)
	zeroPosList, ninePosList := createPositionList(intList)

	for i := 0; i < len(zeroPosList); i++ {
		x, y := zeroPosList[i][0], zeroPosList[i][1]

		for j := 0; j < len(ninePosList); j++ {
			w, z := ninePosList[j][0], ninePosList[j][1]

			result := findTrailHeadScore(intList, x, y, w, z)

			if result > 0 {
				resultPart1++
			}

			resultPart2 += result
		}
	}
	return resultPart1, resultPart2
}

func Run(filename string) {
	part1, part2 := day10(filename)
	fmt.Println("Day10 Part1 Result:", part1)
	fmt.Println("Day10 Part2 Result:", part2)
}
