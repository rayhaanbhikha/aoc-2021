package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Grid struct {
	grid           [][]int
	maxCol, maxRow int
}

func NewGrid(rawInput []string) *Grid {
	grid := make([][]int, 0)
	for _, row := range rawInput {
		gridRow := make([]int, 0)
		for _, rawNum := range strings.Split(row, "") {
			num, _ := strconv.Atoi(rawNum)
			gridRow = append(gridRow, num)
		}
		grid = append(grid, gridRow)
	}
	return &Grid{grid: grid, maxCol: len(grid[0]) - 1, maxRow: len(grid) - 1}
}

func (g *Grid) getVal(row, col int) int {
	return g.grid[row][col]
}

func main() {
	data, _ := ioutil.ReadFile("../sample2")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println(inputs)
	g := NewGrid(inputs)
	memo := make(map[string]int)
	riskLevel := traverseGrid(0, 0, g, memo)
	fmt.Println(memo)
	fmt.Println(riskLevel)
}

func traverseGrid(row, col int, g *Grid, memo map[string]int) int {
	key := fmt.Sprintf("%d:%d", row, col)
	if val, ok := memo[key]; ok {
		return val
	}

	riskLevel := g.getVal(row, col)
	if row == g.maxRow && col == g.maxCol {
		return riskLevel
	}

	translations := [][2]int{{0, 1}, {1, 0}}
	result := math.MaxInt

	for _, translation := range translations {
		newCol := col + translation[0]
		newRow := row + translation[1]
		if newRow > g.maxRow || newCol > g.maxCol {
			continue
		}
		if res := traverseGrid(newRow, newCol, g, memo); res < result {
			result = res
		}
	}

	memo[key] = result + riskLevel

	return memo[key]
}
