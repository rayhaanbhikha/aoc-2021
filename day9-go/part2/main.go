package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type CaveAreaLocation struct {
	row, col int
	caveArea [][]int
}

func (c *CaveAreaLocation) computeBasinSize() int {
	currentSize := 0
	visitedLocations := make(map[string]struct{})
	queue := findNeighbours(c.col, c.row, c.caveArea)

	for len(queue) != 0 {
		newLocation := queue[0]
		queue = queue[1:]
		key := fmt.Sprintf("%d:%d", newLocation.col, newLocation.row)
		if _, ok := visitedLocations[key]; ok {
			continue
		}
		visitedLocations[key] = struct{}{}
		if c.caveArea[newLocation.row][newLocation.col] == 9 {
			continue
		}
		currentSize++
		queue = append(queue, findNeighbours(newLocation.col, newLocation.row, c.caveArea)...)
	}
	return currentSize
}

func (c *CaveAreaLocation) isLowest() (int, bool) {
	currentVal := c.caveArea[c.row][c.col]
	neighbourCoords := findNeighbours(c.col, c.row, c.caveArea)

	for _, location := range neighbourCoords {
		if c.caveArea[location.row][location.col] <= currentVal {
			return 0, false
		}
	}

	return 1 + currentVal, true
}

func findNeighbours(col, row int, caveArea [][]int) []struct{ row, col int } {

	maxCol := len(caveArea[0]) - 1
	maxRow := len(caveArea) - 1

	isTopLeft := row == 0 && col == 0
	isTopRight := row == 0 && col == maxCol
	isTopRow := row == 0
	isBottomRow := row == maxRow
	isLeftCol := col == 0
	isRightCol := col == maxCol
	isBottomLeft := row == maxRow && col == 0
	isBottomRight := row == maxRow && col == maxCol

	var neighbourCoords [][2]int // [x, y]

	switch {
	case isTopRight:
		neighbourCoords = [][2]int{{-1, 0}, {0, 1}}
	case isTopLeft:
		neighbourCoords = [][2]int{{1, 0}, {0, 1}}
	case isBottomRight:
		neighbourCoords = [][2]int{{-1, 0}, {0, -1}}
	case isBottomLeft:
		neighbourCoords = [][2]int{{1, 0}, {0, -1}}
	case isTopRow:
		neighbourCoords = [][2]int{{-1, 0}, {1, 0}, {0, 1}}
	case isBottomRow:
		neighbourCoords = [][2]int{{-1, 0}, {1, 0}, {0, -1}}
	case isLeftCol:
		neighbourCoords = [][2]int{{1, 0}, {0, -1}, {0, 1}}
	case isRightCol:
		neighbourCoords = [][2]int{{-1, 0}, {0, -1}, {0, 1}}
	default:
		neighbourCoords = [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	}

	computedNeighbourIndexes := make([]struct{ row, col int }, 0)
	for _, location := range neighbourCoords {
		computedNeighbourIndexes = append(computedNeighbourIndexes, struct {
			row int
			col int
		}{
			col: location[0] + col,
			row: location[1] + row,
		})
	}
	return computedNeighbourIndexes
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	rows := strings.Split(strings.TrimSpace(string(data)), "\n")

	caveArea := make([][]int, 0)
	for _, row := range rows {
		caveAreaRow := make([]int, 0)
		for _, col := range strings.Split(strings.TrimSpace(row), "") {
			height, _ := strconv.Atoi(col)
			caveAreaRow = append(caveAreaRow, height)
		}
		caveArea = append(caveArea, caveAreaRow)
	}

	basinSizes := make([]int, 0)

	for rowIndex, row := range caveArea {
		for colIndex := range row {
			caveAreaLocation := &CaveAreaLocation{row: rowIndex, col: colIndex, caveArea: caveArea}
			if _, ok := caveAreaLocation.isLowest(); ok {
				basinSizes = append(basinSizes, caveAreaLocation.computeBasinSize())
			}
		}
	}

	sort.Ints(basinSizes)
	n := len(basinSizes) - 1
	fmt.Println(basinSizes[n] * basinSizes[n-1] * basinSizes[n-2])
}
