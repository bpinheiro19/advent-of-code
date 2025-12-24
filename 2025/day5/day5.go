package day5

import (
	"fmt"
	"slices"
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

func (database *Database) createDatabase(filename string) [][]byte {
	database.ingredientRange = make([]Range, 0)
	input := utils.CreateByte2DArray(filename)

	for i, b := range input {
		if len(b) == 0 {
			return input[i+1:]
		}
		str := strings.Split(string(b), "-")
		database.createRange(int(utils.StringToInt(str[0])), int(utils.StringToInt(str[1])))
	}
	return input
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

func (database *Database) countTotalFreshItems() {
	for i := 0; i < len(database.ingredientRange)-1; i++ {

		if database.ingredientRange[i].max+1 >= database.ingredientRange[i+1].min {

			database.ingredientRange[i+1].min = database.ingredientRange[i].min

			if database.ingredientRange[i].max > database.ingredientRange[i+1].max {
				database.ingredientRange[i+1].max = database.ingredientRange[i].max
			}
			database.ingredientRange = slices.Delete(database.ingredientRange, i, i+1)
			i--
		}
	}

	for i := 0; i < len(database.ingredientRange); i++ {
		database.freshItems += database.ingredientRange[i].max - database.ingredientRange[i].min + 1
	}
}

func day5Part1(filename string) int {
	database := &Database{}
	input := database.createDatabase(filename)
	database.sortDatabase()
	database.findFreshItems(input)
	return database.freshItems
}

func day5Part2(filename string) int {
	database := &Database{}
	database.createDatabase(filename)
	database.sortDatabase()
	database.countTotalFreshItems()
	return database.freshItems
}

func Run(filename string) {
	fmt.Println("Day5 Part1 Result:", day5Part1(filename))
	fmt.Println("Day5 Part2 Result:", day5Part2(filename))
}
