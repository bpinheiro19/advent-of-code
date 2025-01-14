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

func day11Part2(filename string) int {
	return 0
}

func Run(filename string) {
	fmt.Println("Day11 Part1 Result:", day11Part1(filename))
	fmt.Println("Day11 Part2 Result:", day11Part2(filename))
}
