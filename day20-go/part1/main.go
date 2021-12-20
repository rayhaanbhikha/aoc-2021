package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var surroundPixelTranslations = [][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {0, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

type Image struct {
	algorithm string
	data      [][]int
}

func NewImage(algorithm string, data []string) *Image {
	imageData := make([][]int, 0)
	for _, row := range data {
		imageDataRow := make([]int, 0)
		for _, pixel := range strings.Split(row, "") {
			pixelVal := 0
			if pixel == "#" {
				pixelVal = 1
			}
			imageDataRow = append(imageDataRow, pixelVal)
		}
		imageData = append(imageData, imageDataRow)
	}
	return &Image{algorithm: algorithm, data: imageData}
}

func (i *Image) getAlgorithmValue(index int) int {
	// index will never be greater than len of algo.
	res := string(i.algorithm[index])
	if res == "#" {
		return 1
	}
	return 0
}

func (i *Image) enhance() {
	i.expandImage()
	i.enhanceImage()
}

func (i *Image) enhanceImage() {
	for rowIdx, row := range i.data {
		for colIdx, currentPixel := range row {
			data := i.getPixelData(rowIdx, colIdx)
			newVal := i.getAlgorithmValue(data)
			if newVal != currentPixel {
				// TODO: to flag change.
				i.setPixelValue(rowIdx, colIdx, newVal)
			}
		}
	}
}

func (i *Image) setPixelValue(row, col, newVal int) {
	i.data[row][col] = newVal
}

func (i *Image) getPixelData(row, col int) int {
	pixValue := 0b0
	for _, translation := range surroundPixelTranslations {
		newCol := col + translation[0]
		newRow := row + translation[1]
		val := i.getPixelValue(newRow, newCol)
		pixValue <<= 1
		pixValue |= val
	}
	return pixValue
}

func (i *Image) getPixelValue(row, col int) int {
	minRow, minCol := 0, 0
	maxRow := len(i.data) - 1
	maxCol := len(i.data[0]) - 1
	if row < minRow || row > maxRow {
		return 0 // default to '.'
	}
	if col < minCol || col > maxCol {
		return 0
	}

	return i.data[row][col]
}

func (i *Image) print() {
	fmt.Println("")
	for _, row := range i.data {
		for _, pixel := range row {
			if pixel == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (i *Image) expandImage() {
	size := 3
	padding := make([]int, size)
	emptyZeroRowTop := make([]int, len(i.data[0]))
	emptyZeroRowBottom := make([]int, len(i.data[0]))

	for index := 0; index < 3; index++ {
		i.data = append(i.data, emptyZeroRowBottom)
		i.data = append([][]int{emptyZeroRowTop}, i.data...)
	}

	for rowIdx := range i.data {
		i.data[rowIdx] = append(i.data[rowIdx], padding...)
		i.data[rowIdx] = append(padding, i.data[rowIdx]...)
	}
}

func main() {
	data, _ := ioutil.ReadFile("../sample")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	image := NewImage(inputs[0], strings.Split(inputs[1], "\n"))
	image.print()
	image.enhance()
	// image.expandImage()
	// image.enhanceImage()
	image.print()
}
