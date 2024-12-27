package main

import (
	"fmt"
	"log"
	"os"
)

func getInputFromFile(file string) []byte {

	input, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return input
}

func createByteArray() [][]byte {

	input := getInputFromFile("input.txt")

	byteArr := make([][]byte, 0)
	lineArr := make([]byte, 0)

	for _, ele := range input {

		if isNewLine(ele) {
			byteArr = append(byteArr, lineArr)
			lineArr = make([]byte, 0)

		} else {
			lineArr = append(lineArr, ele)
		}

	}
	byteArr = append(byteArr, lineArr)

	return byteArr
}

func day4Part1() int {
	result := 0
	byteArr := createByteArray()

	lenI := len(byteArr)
	for i := 0; i < lenI; i++ {
		lenJ := len(byteArr[i])
		for j := 0; j < lenJ; j++ {

			if isX(byteArr[i][j]) {

				for a := -1; a < 2; a++ {
					for b := -1; b < 2; b++ {

						if (inBounds(i+a, j+b, lenI, lenJ) && isM(byteArr[i+a][j+b])) &&
							(inBounds(i+(a*2), j+(b*2), lenI, lenJ) && isA(byteArr[i+(a*2)][j+(b*2)])) &&
							(inBounds(i+(a*3), j+(b*3), lenI, lenJ) && isS(byteArr[i+(a*3)][j+(b*3)])) {
							result++
						}
					}
				}
			}
		}
	}

	return result
}

func inBounds(i int, j int, limitI int, limitJ int) bool {
	if i >= 0 && i < limitI && j >= 0 && j < limitJ {
		return true
	}
	return false
}

/*
Possible solutions
M.M M.S S.M S.S
.A. .A. .A. .A.
S.S M.S S.M M.M

-1,-1  -1,1

	0,0

1,-1    1,1
*/
func day4Part2() int {
	result := 0

	byteArr := createByteArray()

	lenI := len(byteArr)
	for i := 0; i < lenI; i++ {
		lenJ := len(byteArr[i])
		for j := 0; j < lenJ; j++ {

			if isA(byteArr[i][j]) {

				if (inBounds(i-1, j-1, lenI, lenJ) && isM(byteArr[i-1][j-1])) &&
					(inBounds(i-1, j+1, lenI, lenJ) && isM(byteArr[i-1][j+1])) &&
					(inBounds(i+1, j-1, lenI, lenJ) && isS(byteArr[i+1][j-1])) &&
					(inBounds(i+1, j+1, lenI, lenJ) && isS(byteArr[i+1][j+1])) {
					result++
				}

				if (inBounds(i-1, j-1, lenI, lenJ) && isM(byteArr[i-1][j-1])) &&
					(inBounds(i-1, j+1, lenI, lenJ) && isS(byteArr[i-1][j+1])) &&
					(inBounds(i+1, j-1, lenI, lenJ) && isM(byteArr[i+1][j-1])) &&
					(inBounds(i+1, j+1, lenI, lenJ) && isS(byteArr[i+1][j+1])) {
					result++
				}

				if (inBounds(i-1, j-1, lenI, lenJ) && isS(byteArr[i-1][j-1])) &&
					(inBounds(i-1, j+1, lenI, lenJ) && isM(byteArr[i-1][j+1])) &&
					(inBounds(i+1, j-1, lenI, lenJ) && isS(byteArr[i+1][j-1])) &&
					(inBounds(i+1, j+1, lenI, lenJ) && isM(byteArr[i+1][j+1])) {
					result++
				}

				if (inBounds(i-1, j-1, lenI, lenJ) && isS(byteArr[i-1][j-1])) &&
					(inBounds(i-1, j+1, lenI, lenJ) && isS(byteArr[i-1][j+1])) &&
					(inBounds(i+1, j-1, lenI, lenJ) && isM(byteArr[i+1][j-1])) &&
					(inBounds(i+1, j+1, lenI, lenJ) && isM(byteArr[i+1][j+1])) {
					result++
				}

			}
		}
	}

	return result
}

/*
UTF-8 Letters in bytes
X - 0x58
M - 0x4D
A - 0x41
S - 0x53
*/

func isNewLine(b byte) bool {
	if b == 10 {
		return true
	}
	return false
}

func isX(b byte) bool {
	if b == 88 {
		return true
	}
	return false
}

func isA(b byte) bool {
	if b == 65 {
		return true
	}
	return false
}

func isM(b byte) bool {
	if b == 77 {
		return true
	}
	return false
}

func isS(b byte) bool {
	if b == 83 {
		return true
	}
	return false
}

func main() {

	fmt.Println("Day4 Part1 Result:", day4Part1())

	fmt.Println("Day4 Part2 Result:", day4Part2())

}
