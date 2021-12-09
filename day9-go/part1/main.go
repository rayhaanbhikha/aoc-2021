package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type CaveAreaLocation struct {
	row, col int
	caveArea [][]int
}

func (c *CaveAreaLocation) isLowest() (int, bool) {

	currentVal := c.caveArea[c.row][c.col]
	neighbourCoords := c.findNeighbours()

	for _, location := range neighbourCoords {
		col := location[0] + c.col
		row := location[1] + c.row
		if c.caveArea[row][col] <= currentVal {
			return 0, false
		}
	}

	return 1 + currentVal, true
}

func (c *CaveAreaLocation) findNeighbours() [][2]int {

	maxCol := len(c.caveArea[0]) - 1
	maxRow := len(c.caveArea) - 1

	isTopLeft := c.row == 0 && c.col == 0
	isTopRight := c.row == 0 && c.col == maxCol
	isTopRow := c.row == 0
	isBottomRow := c.row == maxRow
	isLeftCol := c.col == 0
	isRightCol := c.col == maxCol
	isBottomLeft := c.row == maxRow && c.col == 0
	isBottomRight := c.row == maxRow && c.col == maxCol

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

	return neighbourCoords
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

	totalRisk := 0

	for rowIndex, row := range caveArea {
		for colIndex := range row {
			caveAreaLocation := &CaveAreaLocation{row: rowIndex, col: colIndex, caveArea: caveArea}
			if risk, ok := caveAreaLocation.isLowest(); ok {
				totalRisk += risk
			}
		}
	}

	fmt.Println(totalRisk)
}
