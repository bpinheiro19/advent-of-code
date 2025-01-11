package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func GetStringListFromFile(filename string) []string {
	return strings.Fields(string(GetByteListFromFile(filename)))
}

func GetByteListFromFile(filename string) []byte {
	input, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return input
}

func InBounds(i int, j int, limitI int, limitJ int) bool {
	return i >= 0 && i < limitI && j >= 0 && j < limitJ
}

func OutOfBounds(i int, j int, limitI int, limitJ int) bool {
	return !InBounds(i, j, limitI, limitJ)
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func CheckRegEx(input string, regex string) []string {
	pattern := regexp.MustCompile(regex)
	return pattern.FindAllString(input, -1)
}

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func ConcatenateTwoInts(a, b int) int {
	return StringToInt(strings.Join([]string{strconv.Itoa(a), strconv.Itoa(b)}, ""))
}

func ReplaceByteInString(str []byte, i int, val byte) string {
	return string(slices.Replace(str, i, i+1, val))
}

func PrintBoard(list []string) {
	fmt.Println("###################")
	for _, str := range list {
		fmt.Println(str)
	}
	fmt.Println("###################")
}

func ByteToInt(b byte) int {
	return int(b - 48)
}

func ByteIsNewLine(b byte) bool {
	return b == 10
}
