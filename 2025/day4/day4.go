package day4

import (
	"fmt"

	"github.com/bpinheiro19/advent-of-code/utils"
)

type Board struct {
	tiles [][]Tile
}

type Tile struct {
	value byte
	x     int
	y     int
}

func (t *Tile) isPaper() bool {
	return t.value == 64
}

func (t *Tile) hasLessThan4AdjacentTiles(b *Board) bool {
	adjNum := 0

	for i := t.x - 1; i <= t.x+1; i++ {
		for j := t.y - 1; j <= t.y+1; j++ {

			if i >= 0 && i < len(b.tiles) && j >= 0 && j < len(b.tiles[i]) {

				if i == t.x && j == t.y {
					continue
				}

				if b.tiles[i][j].isPaper() {
					adjNum++
				}
			}
		}
	}
	return adjNum < 4
}

func (b *Board) initBoard(byteArr [][]byte) {

	for i := 0; i < len(byteArr); i++ {

		tiles := make([]Tile, 0)
		for j := 0; j < len(byteArr[i]); j++ {

			tile := &Tile{byteArr[i][j], i, j}
			tiles = append(tiles, *tile)
		}
		b.tiles = append(b.tiles, tiles)
	}

	printBoard(b.tiles)
}

func day4Part1(filename string) int {
	result := 0

	byteArr := utils.CreateByte2DArray(filename)

	board := &Board{}
	board.initBoard(byteArr)

	for i := 0; i < len(board.tiles); i++ {
		for j := 0; j < len(board.tiles[i]); j++ {

			if board.tiles[i][j].isPaper() && board.tiles[i][j].hasLessThan4AdjacentTiles(board) {
				result++
			}
		}
	}
	return result
}

func day4Part2(filename string) int {
	result := 0
	return result
}

func Run(filename string) {
	fmt.Println("Day4 Part1 Result:", day4Part1(filename))
	fmt.Println("Day4 Part2 Result:", day4Part2(filename))
}
