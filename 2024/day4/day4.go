package main

import (
	"fmt"
	"log"
	"os"
)

func getByteListFromFile(file string) []byte {

	input, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return input
}

func day4Part1() int {
	result := 0
	input := getByteListFromFile("input2.txt")

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

	for _, ele := range byteArr {
		fmt.Println(string(ele))
	}

	return result
}

func trySolve(byteArr [][]byte) {
}

func day4Part2() int {
	result := 0
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
