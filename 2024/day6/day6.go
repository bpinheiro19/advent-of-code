package day6

import (
	"fmt"
	"slices"

	"github.com/bpinheiro19/advent-of-code/utils"
)

/*
Byte Char
46   .
35   #
94   ^
62   >
60   <
118  v
*/

/*
Dir  Int
UP    0
DOWN  1
LEFT  2
RIGHT 3
*/

type GuardDirection int

const (
	UP GuardDirection = iota
	DOWN
	LEFT
	RIGHT
)

const MAX_ATTEMPTS int = 10000

func charIsObstacle(b byte) bool {
	return b == 35
}

func rotateChar90(char byte) byte {
	var rotatedChar byte = 0
	switch char {
	case 94:
		rotatedChar = 62
	case 62:
		rotatedChar = 118
	case 118:
		rotatedChar = 60
	case 60:
		rotatedChar = 94
	}
	return rotatedChar
}

func getDir(char byte) GuardDirection {
	direction := UP
	switch char {
	case 94:
		direction = UP
	case 118:
		direction = DOWN
	case 60:
		direction = LEFT
	case 62:
		direction = RIGHT
	}
	return direction
}

func getDxDy(dir GuardDirection) (int, int) {
	dx, dy := 0, 0
	switch dir {
	case 0:
		dx, dy = -1, 0
	case 1:
		dx, dy = 1, 0
	case 2:
		dx, dy = 0, -1
	case 3:
		dx, dy = 0, 1
	}
	return dx, dy
}

func charIsGuard(c byte) bool {
	return c == 94 || c == 62 || c == 60 || c == 118
}

func getGuardPosition(list []string) (int, int) {
	for i := 0; i < len(list); i++ {
		for n := 0; n < len(list[i]); n++ {
			if charIsGuard(list[i][n]) {
				return i, n
			}
		}
	}
	return -1, -1
}

func findTotalMoves(list []string) int {
	moves := 0
	for i := 0; i < len(list); i++ {
		for n := 0; n < len(list[i]); n++ {
			if list[i][n] == 88 {
				moves++
			}
		}
	}
	return moves
}

func moveGuard(list []string) int {
	attempts := 0

	gx, gy := getGuardPosition(list)
	dir := getDir(list[gx][gy])
	dx, dy := getDxDy(dir)

	for !utils.OutOfBounds(gx+dx, gy+dy, len(list), len(list[gx])) && attempts < MAX_ATTEMPTS {

		if charIsObstacle(list[gx+dx][gy+dy]) {

			guard := rotateChar90(list[gx][gy])
			dir = getDir(list[gx][gy])
			dx, dy = getDxDy(dir)

			list[gx] = replaceByteInString([]byte(list[gx]), gy, guard)

		} else {

			guard := list[gx][gy]

			list[gx] = markPathCrossedWithX([]byte(list[gx]), gy)
			list[gx+dx] = replaceByteInString([]byte(list[gx+dx]), gy+dy, guard)

			gx, gy = gx+dx, gy+dy
		}
		attempts++
	}

	list[gx] = markPathCrossedWithX([]byte(list[gx]), gy)

	return attempts
}

func markPathCrossedWithX(str []byte, i int) string {
	return string(slices.Replace(str, i, i+1, 88))
}

func replaceByteInString(str []byte, i int, val byte) string {
	return string(slices.Replace(str, i, i+1, val))
}

func day6Part1(filename string) int {
	list := utils.GetStringListFromFile(filename)
	moveGuard(list)
	return findTotalMoves(list)
}

func day6Part2(filename string) int {
	positions := 0

	list := utils.GetStringListFromFile(filename)

	for i := 0; i < len(list); i++ {
		for n := 0; n < len(list[i]); n++ {

			if !charIsGuard(list[i][n]) && !charIsObstacle(list[i][n]) {

				newList := make([]string, len(list))
				copy(newList, list)

				newList[i] = replaceByteInString([]byte(list[i]), n, 35)

				attempts := moveGuard(newList)

				if attempts == MAX_ATTEMPTS {
					positions++
				}
			}
		}
	}

	return positions
}

func Run(filename string) {
	fmt.Println("Day6 Part1 Result:", day6Part1(filename))
	fmt.Println("Day6 Part2 Result:", day6Part2(filename))
}
