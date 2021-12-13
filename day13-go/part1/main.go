package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func printGrid(grid [][]string) {
	fmt.Println("")
	for _, row := range grid {
		for _, item := range row {
			fmt.Print(item, " ")
		}
		fmt.Println("")
	}
}

type TransparentPaper struct {
	grid [][]string
}

func (t *TransparentPaper) foldX(col int) {
	grid1 := make([][]string, 0)
	grid2 := make([][]string, 0)
	for _, row := range t.grid {
		grid1 = append(grid1, row[:col])
		gridRow2 := row[col+1:]
		for i, j := 0, len(gridRow2)-1; i < j; i, j = i+1, j-1 {
			gridRow2[i], gridRow2[j] = gridRow2[j], gridRow2[i]
		}

		grid2 = append(grid2, gridRow2)
	}

	t.mergeGrids(grid1, grid2)
}

func (t *TransparentPaper) foldY(row int) {
	grid1 := t.grid[:row]
	grid2 := t.grid[row+1:]

	for i, j := 0, len(grid2)-1; i < j; i, j = i+1, j-1 {
		grid2[i], grid2[j] = grid2[j], grid2[i]
	}

	t.mergeGrids(grid1, grid2)
}

func (t *TransparentPaper) mergeGrids(grid1, grid2 [][]string) {
	for rowIdx, row := range grid2 {
		for colIdx, col := range row {
			if col == "#" {
				grid1[rowIdx][colIdx] = col
			}
		}
	}
	t.grid = grid1
}

func (t *TransparentPaper) computeDots() int {
	dots := 0
	for _, row := range t.grid {
		for _, col := range row {
			if col == "#" {
				dots++
			}
		}
	}
	return dots
}

func (t *TransparentPaper) print() {
	printGrid(t.grid)
}

func NewTransparentPaperFromCoords(coords []string) *TransparentPaper {
	maxX := 0
	maxY := 0

	generatedCoords := make(map[string]struct{})

	for _, coord := range coords {
		res := strings.Split(coord, ",")
		x, _ := strconv.Atoi(res[0])
		y, _ := strconv.Atoi(res[1])
		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}

		key := fmt.Sprintf("%d:%d", x, y)
		generatedCoords[key] = struct{}{}
	}

	grid := make([][]string, 0)

	for j := 0; j <= maxY; j++ {
		gridRow := make([]string, 0)
		for i := 0; i <= maxX; i++ {
			key := fmt.Sprintf(`%d:%d`, i, j)
			if _, ok := generatedCoords[key]; ok {
				gridRow = append(gridRow, "#")
			} else {
				gridRow = append(gridRow, ".")
			}
		}
		grid = append(grid, gridRow)
	}

	return &TransparentPaper{grid: grid}
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	coords := strings.Split(inputs[0], "\n")
	folds := strings.Split(inputs[1], "\n")

	tp := NewTransparentPaperFromCoords(coords)

	for _, fold := range folds[:1] {
		fold = strings.TrimPrefix(fold, "fold along ")
		res := strings.Split(fold, "=")
		foldAxis := res[0]
		foldIndex, _ := strconv.Atoi(res[1])
		if foldAxis == "x" {
			tp.foldX(foldIndex)
		} else {
			tp.foldY(foldIndex)
		}
	}

	// tp.foldY(7)
	fmt.Println(tp.computeDots())
}
