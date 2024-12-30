package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
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

func getStringListFromFile() []string {

	input, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	return strings.Fields(string(input))
}

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

func outOfBounds(i int, j int, limitI int, limitJ int) bool {
	return i < 0 || i >= limitI || j < 0 || j >= limitJ
}

func printBoard(list []string) {
	fmt.Println("###################")
	for _, str := range list {
		fmt.Println(str)
	}
	fmt.Println("###################")
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

	for !outOfBounds(gx+dx, gy+dy, len(list), len(list[gx])) && attempts < MAX_ATTEMPTS {

		if charIsObstacle(list[gx+dx][gy+dy]) {

			guard := rotateChar90(list[gx][gy])
			dir = getDir(list[gx][gy])
			dx, dy = getDxDy(dir)

			list[gx] = string(slices.Replace([]byte(list[gx]), gy, gy+1, guard))

		} else {

			guard := list[gx][gy]

			list[gx] = string(slices.Replace([]byte(list[gx]), gy, gy+1, 88))
			list[gx+dx] = string(slices.Replace([]byte(list[gx+dx]), gy+dy, gy+dy+1, guard))

			gx, gy = gx+dx, gy+dy

		}

		attempts++
	}

	list[gx] = string(slices.Replace([]byte(list[gx]), gy, gy+1, 88))

	return attempts
}

func day6Part1() int {
	list := getStringListFromFile()
	moveGuard(list)
	return findTotalMoves(list)
}

func day6Part2() int {
	positions := 0

	list := getStringListFromFile()

	for i := 0; i < len(list); i++ {
		for n := 0; n < len(list[i]); n++ {

			if !charIsGuard(list[i][n]) && !charIsObstacle(list[i][n]) {

				newList := make([]string, len(list))
				copy(newList, list)

				newList[i] = string(slices.Replace([]byte(list[i]), n, n+1, 35))

				attempts := moveGuard(newList)

				if attempts == MAX_ATTEMPTS {
					positions++
				}
			}
		}
	}

	return positions
}

func main() {

	fmt.Println("Day6 Part1 Result:", day6Part1())
	fmt.Println("Day6 Part2 Result:", day6Part2())

}
