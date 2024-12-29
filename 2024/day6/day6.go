package main

import (
	"fmt"
	"log"
	"os"
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

func getStringListFromFile() []string {

	input, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	return strings.Fields(string(input))
}

func hasObstacle(b byte) bool {
	if b == 35 {
		return true
	}
	return false
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

func getGuardPosition(list []string) (int, int) {
	for i := 0; i < len(list); i++ {
		for n := 0; n < len(list[i]); n++ {
			if list[i][n] == 94 || list[i][n] == 62 || list[i][n] == 60 || list[i][n] == 118 {
				return i, n
			}
		}
	}
	return -1, -1
}

func inBounds(i int, j int, limitI int, limitJ int) bool {
	if i >= 0 && i < limitI && j >= 0 && j < limitJ {
		return true
	}
	return false
}

func insertIntoArray(arr []byte, i int, e byte) []byte {
	return append(arr[:i], append([]byte{e}, arr[i:]...)...)
}

func removeFromArray(arr []byte, i int) []byte {
	return append(arr[:i], arr[i+1:]...)
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

func day6Part1() int {

	list := getStringListFromFile()

	gx, gy := getGuardPosition(list)
	dir := getDir(list[gx][gy])
	dx, dy := getDxDy(dir)

	for inBounds(gx+dx, gy+dy, len(list), len(list[gx])) {

		if hasObstacle(list[gx+dx][gy+dy]) {

			guard := rotateChar90(list[gx][gy])

			list[gx] = string(removeFromArray([]byte(list[gx]), gy))
			list[gx] = string(insertIntoArray([]byte(list[gx]), gy, guard))

		} else {
			guard := list[gx][gy]

			list[gx] = string(removeFromArray([]byte(list[gx]), gy))
			list[gx] = string(insertIntoArray([]byte(list[gx]), gy, 88))

			list[gx+dx] = string(removeFromArray([]byte(list[gx+dx]), gy+dy))
			list[gx+dx] = string(insertIntoArray([]byte(list[gx+dx]), gy+dy, guard))
		}

		gx, gy = getGuardPosition(list)
		dir = getDir(list[gx][gy])
		dx, dy = getDxDy(dir)
	}

	list[gx] = string(removeFromArray([]byte(list[gx]), gy))
	list[gx] = string(insertIntoArray([]byte(list[gx]), gy, 88))

	return findTotalMoves(list)
}

func day6Part2() int {
	result := 0

	return result
}

func main() {

	fmt.Println("Day6 Part1 Result:", day6Part1())
	fmt.Println("Day6 Part2 Result:", day6Part2())

}
