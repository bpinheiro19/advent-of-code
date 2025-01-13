package day11

import (
	"aoc2024/utils"
	"fmt"
	"slices"
)

type Stone struct {
	val int
}

func blink(stoneList []*Stone) []*Stone {

	for i := 0; i < len(stoneList); i++ {

		if stoneList[i].val == 0 {
			stoneList[i].val = 1

		} else if len(utils.IntToString(stoneList[i].val))%2 == 0 {

			bs := []byte(utils.IntToString(stoneList[i].val))
			halfLen := len(bs) / 2

			stoneList[i].val = utils.ByteListToInt(bs[halfLen:])
			stone := &Stone{utils.ByteListToInt(bs[:halfLen])}

			stoneList = slices.Insert(stoneList, i, stone)
			i++
		} else {
			stoneList[i].val *= 2024
		}
	}
	return stoneList
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
		stoneList = blink(stoneList)
	}

	return len(stoneList)
}

func day11Part2(filename string) int {
	return 0
}

func Run(filename string) {
	fmt.Println("Day11 Part1 Result:", day11Part1(filename))
	fmt.Println("Day11 Part2 Result:", day11Part2(filename))
}
