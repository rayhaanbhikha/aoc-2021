package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func computeTranslations(row, col, maxRow, maxCol int) []Translation {
	translations := make([]Translation, 0)

	isTopLeft := row == 0 && col == 0
	isTopRight := row == 0 && col == maxCol
	isTopRow := row == 0
	isBottomRow := row == maxRow
	isLeftCol := col == 0
	isRightCol := col == maxCol
	isBottomLeft := row == maxRow && col == 0
	isBottomRight := row == maxRow && col == maxCol

	switch {
	case isTopRight:
		translations = append(translations, S, SW, W)
	case isTopLeft:
		translations = append(translations, E, SE, S)
	case isBottomRight:
		translations = append(translations, N, NW, W)
	case isBottomLeft:
		translations = append(translations, N, NE, E)
	case isTopRow:
		translations = append(translations, E, SE, S, SW, W)
	case isBottomRow:
		translations = append(translations, E, NE, N, NW, W)
	case isLeftCol:
		translations = append(translations, N, NE, E, SE, S)
	case isRightCol:
		translations = append(translations, N, NW, W, SW, S)
	default:
		translations = append(translations, N, NE, E, SE, S, SW, W, NW)
	}

	return translations
}

type Translation [2]int

var (
	N  = Translation{0, -1}
	NE = Translation{1, -1}
	E  = Translation{1, 0}
	SE = Translation{1, 1}
	S  = Translation{0, 1}
	SW = Translation{-1, 1}
	W  = Translation{-1, 0}
	NW = Translation{-1, -1}
)

type Octopus struct {
	row, col    int
	energyLevel int
	hasFlashed  bool
	maxEnergy   int
}

func NewOctopus(row, col, energyLevel int) *Octopus {
	return &Octopus{row: row, col: col, energyLevel: energyLevel, maxEnergy: 9}
}

func (o *Octopus) reset() {
	o.hasFlashed = false
}

func (o *Octopus) incrementEnergyLevel() bool {
	if o.hasFlashed {
		return false
	}

	if o.energyLevel == o.maxEnergy {
		o.hasFlashed = true
		o.energyLevel = 0
		return true
	}

	o.energyLevel++
	return false
}

type OctopusGrid struct {
	grid       [][]*Octopus
	flashCount int
	debug      bool
}

func NewOctopusGrid(rawGrid []string) *OctopusGrid {
	grid := make([][]*Octopus, 0)
	for rowIdx, row := range rawGrid {
		gridRow := make([]*Octopus, 0)
		for colIdx, col := range strings.Split(strings.TrimSpace(row), "") {
			energyLevel, _ := strconv.Atoi(col)
			gridRow = append(gridRow, NewOctopus(rowIdx, colIdx, energyLevel))
		}
		grid = append(grid, gridRow)
	}
	return &OctopusGrid{grid: grid}
}

func (og *OctopusGrid) step() {
	// loop through octopuses and increment energy levels by one.
	flashedOctopi := make([]*Octopus, 0)
	for _, row := range og.grid {
		for _, oct := range row {
			hasFlashed := oct.incrementEnergyLevel()
			if hasFlashed {
				og.flashCount++
				flashedOctopi = append(flashedOctopi, oct)
			}
		}
	}

	for _, flashedOctopus := range flashedOctopi {
		og.incrementAdjacentOctopi(flashedOctopus)
	}

	if og.debug {
		og.print()
	}

	// reset all flashes before moving.
	for _, row := range og.grid {
		for _, oct := range row {
			oct.reset()
		}
	}
	// if any flash compute their neighbours and increment them by 1.
}

func (og *OctopusGrid) incrementAdjacentOctopi(octopus *Octopus) {
	// flashedOctopi := make([]*Octopus, 0)

	neighbours := og.findNeighbours(octopus)
	for _, neighbourOctopus := range neighbours {
		hasFlashed := neighbourOctopus.incrementEnergyLevel()
		if hasFlashed {
			og.flashCount++
			// flashedOctopi = append(flashedOctopi, neighbourOctopus)
			og.incrementAdjacentOctopi(neighbourOctopus)
		}
	}

}

func (og *OctopusGrid) incrementOctopushEnergyLevels(octopi []*Octopus) {
	for _, octopus := range octopi {
		if octopus.hasFlashed {
			continue
		}
		hasFlashed := octopus.incrementEnergyLevel()
		if hasFlashed {
			og.flashCount++
			neighbours := og.findNeighbours(octopus)
			og.incrementOctopushEnergyLevels(neighbours)
		}
	}
}

func (og *OctopusGrid) print() {
	fmt.Println("")
	for _, row := range og.grid {
		for _, octopus := range row {
			if octopus.hasFlashed {
				fmt.Printf("\033[32m%d\033[0m ", octopus.energyLevel)
			} else {
				fmt.Print(octopus.energyLevel, " ")
			}
		}
		fmt.Println("")
	}
}

func (og *OctopusGrid) findNeighbours(octopus *Octopus) []*Octopus {
	maxCol := len(og.grid[0]) - 1
	maxRow := len(og.grid) - 1
	translations := computeTranslations(octopus.row, octopus.col, maxRow, maxCol)
	neighbours := make([]*Octopus, 0)
	for _, translation := range translations {
		newRow := octopus.row + translation[1]
		newCol := octopus.col + translation[0]
		octupus := og.grid[newRow][newCol]
		neighbours = append(neighbours, octupus)
	}

	return neighbours
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	octopusGrid := NewOctopusGrid(inputs)

	for i := 0; i < 100; i++ {
		octopusGrid.step()
	}

	fmt.Println(octopusGrid.flashCount)
}