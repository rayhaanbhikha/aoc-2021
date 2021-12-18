package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Grid struct {
	grid           [][]int
	maxRow, maxCol int
}

func NewGrid(rawInput []string, scale int) *Grid {
	grid := make([][]int, 0)
	for _, row := range rawInput {
		gridRow := make([]int, 0)
		for _, rawNum := range strings.Split(row, "") {
			num, _ := strconv.Atoi(rawNum)
			gridRow = append(gridRow, num)
		}
		newGridRow := gridRow
		for i := 0; i < scale-1; i++ {
			for _, riskLevel := range gridRow {
				val := riskLevel + 1 + i
				if val > 9 {
					val %= 10
					val++
				}
				newGridRow = append(newGridRow, val)
			}
		}
		grid = append(grid, newGridRow)
	}

	rowsToAppend := make([][]int, len(grid)*(scale-1))

	for rowIdx, row := range grid {

		for i := 0; i < scale-1; i++ {
			newRow := make([]int, 0, len(row))
			for _, riskLevel := range row {
				val := riskLevel + 1 + i
				if val > 9 {
					val %= 10
					val++
				}
				newRow = append(newRow, val)
			}
			// fmt.Println(i, newRow)
			rowsToAppend[rowIdx+len(grid)*i] = newRow
			// grid = append(grid, newRow)
			// 	rowsToAppend = append(rowsToAppend, newRow)
			// 	// grid = append(grid, newRow)
		}
	}

	grid = append(grid, rowsToAppend...)
	return &Grid{grid: grid, maxRow: len(grid) - 1, maxCol: len(grid[0]) - 1}
}

func (g *Grid) getVal(row, col int) int {
	return g.grid[row][col]
}

func (g *Grid) print() {
	fmt.Println("")
	for _, row := range g.grid {
		for _, riskLevel := range row {
			fmt.Print(riskLevel)
		}
		fmt.Println("")
	}
}
