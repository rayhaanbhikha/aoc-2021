package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Grid struct {
	cucumbers      map[string]*Cucumber
	maxRow, maxCol int
}

func NewGrid(maxRow, maxCol int) *Grid {
	return &Grid{
		cucumbers: make(map[string]*Cucumber),
		maxRow:    maxRow,
		maxCol:    maxCol,
	}
}

func (g *Grid) addCucumber(cucumber *Cucumber) {
	key := fmt.Sprintf("%d:%d", cucumber.row, cucumber.col)
	g.cucumbers[key] = cucumber
}

func (g *Grid) takeStep() bool {
	eastFacingHeard := make([]*Cucumber, 0)
	southFacingHeard := make([]*Cucumber, 0)

	for _, cucumber := range g.cucumbers {
		switch cucumber.direction {
		case East:
			eastFacingHeard = append(eastFacingHeard, cucumber)
		case South:
			southFacingHeard = append(southFacingHeard, cucumber)
		}
	}

	hasMoved := false

	// ################## East ##################

	cucumbersToMove := make([]*Cucumber, 0)

	for _, cucumber := range eastFacingHeard {
		nextRow, nextCol := cucumber.nextStep(g.maxRow, g.maxCol)
		if _, ok := g.get_cucumber(nextRow, nextCol); !ok {
			cucumbersToMove = append(cucumbersToMove, cucumber)
		}
	}

	if len(cucumbersToMove) > 0 {
		hasMoved = true
	}

	for _, c := range cucumbersToMove {
		delete(g.cucumbers, g.generateKey(c))
		c.takeStep(g.maxRow, g.maxCol)
		g.addCucumber(c)
	}

	// ################## South ##################

	cucumbersToMove = make([]*Cucumber, 0)

	for _, cucumber := range southFacingHeard {
		nextRow, nextCol := cucumber.nextStep(g.maxRow, g.maxCol)
		if _, ok := g.get_cucumber(nextRow, nextCol); !ok {
			cucumbersToMove = append(cucumbersToMove, cucumber)
		}
	}

	if len(cucumbersToMove) > 0 {
		hasMoved = true
	}

	for _, c := range cucumbersToMove {
		delete(g.cucumbers, g.generateKey(c))
		c.takeStep(g.maxRow, g.maxCol)
		g.addCucumber(c)
	}

	return hasMoved
}

func (g *Grid) generateKey(c *Cucumber) string {
	return fmt.Sprintf("%d:%d", c.row, c.col)
}

func (g *Grid) get_cucumber(row, col int) (*Cucumber, bool) {
	key := fmt.Sprintf("%d:%d", row, col)
	val, ok := g.cucumbers[key]
	return val, ok
}

func (g *Grid) String() string {
	m := "\n"
	for i := 0; i < g.maxRow; i++ {
		for j := 0; j < g.maxRow; j++ {
			val, ok := g.get_cucumber(i, j)
			if ok {
				m += string(val.rawDirection)
			} else {
				m += "."
			}
		}
		m += "\n"
	}
	return m
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	grid := NewGrid(len(inputs), len(inputs[0]))

	for rowIndex, row := range inputs {
		for colIndex, directionChar := range row {
			cucumber, err := NewCucumber(rowIndex, colIndex, directionChar)
			if err != nil {
				continue
			}
			grid.addCucumber(cucumber)
		}
	}

	fmt.Println(grid)

	i := 0
	for {
		i++
		hasMoved := grid.takeStep()
		if !hasMoved {
			break
		}
		// fmt.Println("Step: ", i)
		// fmt.Println(grid)
	}
	fmt.Println(i)
}
