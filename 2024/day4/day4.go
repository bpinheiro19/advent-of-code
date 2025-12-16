package day4

import (
	"fmt"

	"github.com/bpinheiro19/advent-of-code/utils"
)

func day4Part1(filename string) int {
	result := 0
	byteArr := utils.CreateByte2DArray(filename)

	lenI := len(byteArr)
	for i := 0; i < lenI; i++ {
		lenJ := len(byteArr[i])
		for j := 0; j < lenJ; j++ {

			if isX(byteArr[i][j]) {

				for a := -1; a < 2; a++ {
					for b := -1; b < 2; b++ {

						if (utils.InBounds(i+a, j+b, lenI, lenJ) && isM(byteArr[i+a][j+b])) &&
							(utils.InBounds(i+(a*2), j+(b*2), lenI, lenJ) && isA(byteArr[i+(a*2)][j+(b*2)])) &&
							(utils.InBounds(i+(a*3), j+(b*3), lenI, lenJ) && isS(byteArr[i+(a*3)][j+(b*3)])) {
							result++
						}
					}
				}
			}
		}
	}
	return result
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

func day4Part2(filename string) int {
	result := 0
	byteArr := utils.CreateByte2DArray(filename)

	lenI := len(byteArr)
	for i := 0; i < lenI; i++ {
		lenJ := len(byteArr[i])
		for j := 0; j < lenJ; j++ {

			if isA(byteArr[i][j]) {

				if (utils.InBounds(i-1, j-1, lenI, lenJ) && isM(byteArr[i-1][j-1])) &&
					(utils.InBounds(i-1, j+1, lenI, lenJ) && isM(byteArr[i-1][j+1])) &&
					(utils.InBounds(i+1, j-1, lenI, lenJ) && isS(byteArr[i+1][j-1])) &&
					(utils.InBounds(i+1, j+1, lenI, lenJ) && isS(byteArr[i+1][j+1])) {
					result++
				}

				if (utils.InBounds(i-1, j-1, lenI, lenJ) && isM(byteArr[i-1][j-1])) &&
					(utils.InBounds(i-1, j+1, lenI, lenJ) && isS(byteArr[i-1][j+1])) &&
					(utils.InBounds(i+1, j-1, lenI, lenJ) && isM(byteArr[i+1][j-1])) &&
					(utils.InBounds(i+1, j+1, lenI, lenJ) && isS(byteArr[i+1][j+1])) {
					result++
				}

				if (utils.InBounds(i-1, j-1, lenI, lenJ) && isS(byteArr[i-1][j-1])) &&
					(utils.InBounds(i-1, j+1, lenI, lenJ) && isM(byteArr[i-1][j+1])) &&
					(utils.InBounds(i+1, j-1, lenI, lenJ) && isS(byteArr[i+1][j-1])) &&
					(utils.InBounds(i+1, j+1, lenI, lenJ) && isM(byteArr[i+1][j+1])) {
					result++
				}

				if (utils.InBounds(i-1, j-1, lenI, lenJ) && isS(byteArr[i-1][j-1])) &&
					(utils.InBounds(i-1, j+1, lenI, lenJ) && isS(byteArr[i-1][j+1])) &&
					(utils.InBounds(i+1, j-1, lenI, lenJ) && isM(byteArr[i+1][j-1])) &&
					(utils.InBounds(i+1, j+1, lenI, lenJ) && isM(byteArr[i+1][j+1])) {
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

func isX(b byte) bool {
	return b == 88
}

func isA(b byte) bool {
	return b == 65
}

func isM(b byte) bool {
	return b == 77
}

func isS(b byte) bool {
	return b == 83
}

func Run(filename string) {
	fmt.Println("Day4 Part1 Result:", day4Part1(filename))
	fmt.Println("Day4 Part2 Result:", day4Part2(filename))
}
