package day9

import (
	"aoc2024/utils"
	"fmt"
)

func createDiskBlock(filename string) ([]int, map[int]int) {
	fileId := 0

	diskList := utils.GetByteListFromFile(filename)
	fileList := make([]int, 0)
	helperMap := make(map[int]int)

	for i := 0; i < len(diskList); i++ {

		len := utils.ByteToInt(diskList[i])
		if i%2 == 0 {
			for j := 0; j < len; j++ {
				fileList = append(fileList, fileId)
			}

			helperMap[fileId] = len
			fileId++

		} else {
			for j := 0; j < len; j++ {
				fileList = append(fileList, -1)
			}
		}
	}
	return fileList, helperMap
}

func findRightMostFileIndex(fileList []int) int {
	for i := len(fileList) - 1; i >= 0; i-- {
		if fileList[i] != -1 {
			return i
		}
	}
	return -1
}

func shiftBlocksLeftPart1(fileList []int) []int {

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

	fileList, _ := createDiskBlock(filename)

	fileList = shiftBlocksLeftPart1(fileList)

	return findCheckSum(fileList)
}

func findFreeSpace(fileList []int, size int) (int, int) {
	beg, end := -1, -1
	for i := 0; i < len(fileList); i++ {
		if fileList[i] == -1 {
			if beg == -1 {
				beg = i
				end = i
			} else {
				end++
			}

		} else {
			if beg != -1 && end != -1 && end-beg+1 >= size {
				return beg, end
			} else {
				beg, end = -1, -1
			}
		}
	}
	return -1, -1
}

func findFileIdIndex(fileList []int, fileId int) int {
	for i := len(fileList) - 1; i >= 0; i-- {
		if fileList[i] == fileId {
			return i
		}
	}
	return -1
}

func shiftBlocksLeftPart2(fileList []int, helperMap map[int]int) []int {

	for i := len(helperMap) - 1; i >= 0; i-- {

		fileId, fileSize := i, helperMap[i]
		beg, end := findFreeSpace(fileList, fileSize)
		index := findFileIdIndex(fileList, fileId)

		if beg != -1 && end != -1 && beg < index {

			for j := 0; j < fileSize; j++ {
				fileList[beg+j] = fileId
				fileList[index-j] = -1
			}

		}
	}

	return fileList
}

func day9Part2(filename string) int {
	fileList, helperMap := createDiskBlock(filename)

	fileList = shiftBlocksLeftPart2(fileList, helperMap)

	return findCheckSum(fileList)
}

func Run(filename string) {
	fmt.Println("Day8 Part1 Result:", day9Part1(filename))
	fmt.Println("Day8 Part2 Result:", day9Part2(filename))
}
