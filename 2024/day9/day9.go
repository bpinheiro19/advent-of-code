package day9

import (
	"aoc2024/utils"
	"fmt"
)

func createDiskBlock(filename string) []int {
	fileId := 0

	diskList := utils.GetByteListFromFile(filename)
	fileList := make([]int, 0)

	for i := 0; i < len(diskList); i++ {

		if i%2 == 0 {
			for j := 0; j < utils.ByteToInt(diskList[i]); j++ {
				fileList = append(fileList, fileId)
			}
			fileId++

		} else {
			for j := 0; j < utils.ByteToInt(diskList[i]); j++ {
				fileList = append(fileList, -1)
			}
		}
	}
	return fileList
}

func findRightMostFileIndex(fileList []int) int {
	for i := len(fileList) - 1; i >= 0; i-- {
		if fileList[i] != -1 {
			return i
		}
	}
	return -1
}

func shiftBlocksLeft(fileList []int) []int {

	index := findRightMostFileIndex(fileList)

	for i := 0; i <= index; i++ {
		for fileList[index] == -1 && i <= index {
			index--
		}

		if fileList[i] == -1 {
			fileList[i] = fileList[index]
			index--
		}
	}

	newList := make([]int, index+1)
	copy(newList, fileList)

	return newList
}

func findCheckSum(fileList []int) int {
	sum := 0
	for i := 0; i < len(fileList); i++ {
		if fileList[i] != -1 {
			sum += fileList[i] * i
		}
	}
	return sum
}

func day9Part1(filename string) int {

	fileList := createDiskBlock(filename)

	fileList = shiftBlocksLeft(fileList)

	return findCheckSum(fileList)
}

func day9Part2(filename string) int {
	result := 0
	return result
}

func Run(filename string) {
	fmt.Println("Day8 Part1 Result:", day9Part1(filename))
	fmt.Println("Day8 Part2 Result:", day9Part2(filename))
}
