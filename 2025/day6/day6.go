package day6

import (
	"fmt"
	"strings"

	"github.com/bpinheiro19/advent-of-code/utils"
)

type Math struct {
	operations []Operation
}

type Operation struct {
	nums     []int
	operator byte
}

func (math *Math) createOperations(filename string) {
	byteArr := utils.CreateByte2DArray(filename)
	math.operations = make([]Operation, 0)

	for i := range byteArr {

		ind := 0
		str := strings.Split(strings.TrimSpace(string(byteArr[i])), " ")

		for _, s := range str {

			if s != "" {
				if i == 0 {
					operation := &Operation{}
					operation.nums = make([]int, 0)
					math.operations = append(math.operations, *operation)
				}

				if s == "+" {
					math.operations[ind].operator = 43

				} else if s == "*" {
					math.operations[ind].operator = 42

				} else {
					math.operations[ind].nums = append(math.operations[ind].nums, utils.StringToInt(s))
				}
				ind++
			}
		}
	}
}

func (math *Math) resolveProblem() int {
	result := 0
	for _, op := range math.operations {
		if op.operator == 43 {
			result += op.sumOperation()
		} else {
			result += op.multiplyOperation()
		}
	}
	return result
}

func (operation *Operation) sumOperation() int {
	sum := 0
	for _, b := range operation.nums {
		sum += b
	}
	return sum
}

func (operation *Operation) multiplyOperation() int {
	sum := 1
	for _, b := range operation.nums {
		sum *= b
	}
	return sum
}

func day6Part1(filename string) int {
	math := &Math{}
	math.createOperations(filename)
	return math.resolveProblem()
}

func day6Part2(filename string) int {
	result := 0
	return result
}

func Run(filename string) {
	fmt.Println("Day6 Part1 Result:", day6Part1(filename))
	fmt.Println("Day6 Part2 Result:", day6Part2(filename))
}
