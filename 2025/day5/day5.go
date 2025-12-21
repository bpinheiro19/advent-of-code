package day5

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bpinheiro19/advent-of-code/utils"
)

type Database struct {
	ingredientRange []Range
	freshItems      int
}

type Range struct {
	min int
	max int
}

func (database *Database) createRange(min int, max int) {
	rg := &Range{min, max}
	database.ingredientRange = append(database.ingredientRange, *rg)
}

func (database *Database) createDatabase(input [][]byte) [][]byte {
	database.ingredientRange = make([]Range, 0)

	for i, b := range input {
		if len(b) == 0 {
			return input[i+1:]
		}
		str := strings.Split(string(b), "-")
		database.createRange(int(utils.StringToInt(str[0])), int(utils.StringToInt(str[1])))
	}
	return [][]byte{}
}

func (database *Database) sortDatabase() {
	sort.Slice(database.ingredientRange, func(i, j int) bool {
		return database.ingredientRange[i].min < database.ingredientRange[j].min
	})
}

func (database *Database) findFreshItems(input [][]byte) {
	for _, bId := range input {
		id := utils.ByteListToInt(bId)

		for i := 0; i+1 < len(database.ingredientRange); i++ {
			if id >= database.ingredientRange[i].min && id <= database.ingredientRange[i].max {
				database.freshItems++
				break
			}
		}
	}
}

func day5Part1(filename string) int {
	input := utils.CreateByte2DArray(filename)

	database := &Database{}
	input = database.createDatabase(input)
	database.sortDatabase()
	database.findFreshItems(input)

	return database.freshItems
}

func day5Part2(filename string) int {
	result := 0
	return result
}

func Run(filename string) {
	fmt.Println("Day5 Part1 Result:", day5Part1(filename))
	fmt.Println("Day5 Part2 Result:", day5Part2(filename))
}
