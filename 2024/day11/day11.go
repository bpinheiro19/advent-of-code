package day11

import (
	"aoc2024/utils"
	"fmt"
)

type Stone struct {
	val int
}

func blink(stoneList *[]*Stone) {
	for _, st := range *stoneList {
		if st.val == 0 {
			st.val = 1

		} else if len(utils.IntToString(st.val))%2 == 0 {
			bs := []byte(utils.IntToString(st.val))
			halfLen := len(bs) / 2

			st.val = utils.ByteListToInt(bs[halfLen:])
			stone := &Stone{utils.ByteListToInt(bs[:halfLen])}

			*stoneList = append(*stoneList, stone)
		} else {
			st.val *= 2024
		}
	}
}

func createStoneList(filename string) []*Stone {
	stringList := utils.GetStringListFromFile(filename)
	stoneList := make([]*Stone, 0)

	for _, str := range stringList {
		stone := &Stone{utils.StringToInt(str)}
		stoneList = append(stoneList, stone)
	}
	return stoneList
}

func day11Part1(filename string) int {

	stoneList := createStoneList(filename)

	for i := 0; i < 25; i++ {
		blink(&stoneList)
	}

	return len(stoneList)
}

func createStoneMap(filename string) map[int]int {

	stringList := utils.GetStringListFromFile(filename)

	stoneMap := make(map[int]int)

	for _, str := range stringList {
		stoneMap[utils.StringToInt(str)] += 1
	}
	return stoneMap
}

func copyMap(stoneMap map[int]int) map[int]int {
	newStoneMap := make(map[int]int, 0)

	for x, i := range stoneMap {
		newStoneMap[x] = i
	}
	return newStoneMap
}

func improvedBlink(stoneMap map[int]int) map[int]int {

	newStoneMap := copyMap(stoneMap)

	for key, val := range stoneMap {

		if key == 0 {
			newStoneMap[key] -= val
			newStoneMap[1] += val

		} else if len(utils.IntToString(key))%2 == 0 {

			newStoneMap[key] -= val

			str := utils.IntToString(key)
			halfLen := len(str) / 2

			newStoneMap[utils.StringToInt(str[:halfLen])] += val
			newStoneMap[utils.StringToInt(str[halfLen:])] += val

		} else {
			newStoneMap[key] -= val
			newStoneMap[key*2024] += val
		}

		if newStoneMap[key] <= 0 {
			delete(newStoneMap, key)
		}
	}
	return newStoneMap
}

func day11Part2(filename string) int {

	stoneMap := createStoneMap(filename)

	for i := 0; i < 75; i++ {
		stoneMap = improvedBlink(stoneMap)
	}

	stoneLength := 0
	for _, val := range stoneMap {
		stoneLength += val
	}

	return stoneLength
}

func Run(filename string) {
	fmt.Println("Day11 Part1 Result:", day11Part1(filename))
	fmt.Println("Day11 Part2 Result:", day11Part2(filename))
}
