package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type BingoBoardNums struct {
	val    int
	marked bool
}

type BingoBoard struct {
	size   int
	grid   [][]*BingoBoardNums
	hasWon bool
}

func NewBingoBoard() *BingoBoard {
	return &BingoBoard{
		size:   5,
		grid:   make([][]*BingoBoardNums, 0),
		hasWon: false,
	}
}

func (b *BingoBoard) addRow(nums []int) {
	formattedNums := make([]*BingoBoardNums, 0, len(nums))
	for _, num := range nums {
		formattedNums = append(formattedNums, &BingoBoardNums{val: num})
	}
	b.grid = append(b.grid, formattedNums)
}

func (b *BingoBoard) mark(numToMark int) {
	for i, row := range b.grid {
		for j, num := range row {
			if num.val == numToMark {
				b.grid[i][j].marked = true
			}
		}
	}
	b.checkIfWon()
}

func (b *BingoBoard) checkIfWon() {
	// check rows
	for i := 0; i < b.size; i++ {
		if b.checkRow(i) || b.checkCol(i) {
			b.hasWon = true
			return
		}
	}
}

func (b *BingoBoard) checkRow(rowIndex int) bool {
	for _, num := range b.grid[rowIndex] {
		if !num.marked {
			return false
		}
	}
	return true
}

func (b *BingoBoard) checkCol(colIndex int) bool {
	for i := 0; i < b.size; i++ {
		if !b.grid[i][colIndex].marked {
			return false
		}
	}
	return true
}

func (b *BingoBoard) sumOfUnMarkedItems() int {
	sum := 0
	for _, row := range b.grid {
		for _, num := range row {
			if !num.marked {
				sum += num.val
			}
		}
	}
	return sum
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	bingoNums := make([]int, 0)
	for _, num := range strings.Split(inputs[0], ",") {
		val, _ := strconv.Atoi(num)
		bingoNums = append(bingoNums, val)
	}

	bingoBoards := make([]*BingoBoard, 0)

	for _, row := range inputs[1:] {
		nums := make([]int, 0)
		for _, num := range strings.Split(row, " ") {
			if num == "" {
				continue
			}
			val, _ := strconv.Atoi(num)
			nums = append(nums, val)
		}
		if len(nums) == 0 {
			bingoBoards = append(bingoBoards, NewBingoBoard())
		} else {
			bingoBoards[len(bingoBoards)-1].addRow(nums)
		}
	}

	for _, bingoNum := range bingoNums {
		for _, bingoBoard := range bingoBoards {
			bingoBoard.mark(bingoNum)
			if bingoBoard.hasWon {
				fmt.Println(bingoBoard.sumOfUnMarkedItems() * bingoNum)
				return
			}
		}
	}
}
