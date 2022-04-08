package main

import "fmt"

type CucumberDirection int

const (
	East CucumberDirection = iota
	South
)

func (d CucumberDirection) getCoords() (int, int) {
	switch d {
	case East:
		return 1, 0
	case South:
		return 0, 1
	default:
		return 0, 0
	}
}

type Cucumber struct {
	row, col     int
	direction    CucumberDirection
	rawDirection rune
}

func NewCucumber(row, col int, direction rune) (*Cucumber, error) {
	var cucumberDirection CucumberDirection
	switch direction {
	case '>':
		cucumberDirection = East
	case 'v':
		cucumberDirection = South
	default:
		return nil, fmt.Errorf("Can't parse %c", direction)
	}

	return &Cucumber{
		row:          row,
		col:          col,
		direction:    cucumberDirection,
		rawDirection: direction,
	}, nil
}

func (c *Cucumber) nextStep(maxRow, maxCol int) (int, int) {
	switch c.direction {
	case South:
		return (c.row + 1) % maxRow, c.col
	case East:
		return c.row, (c.col + 1) % maxCol
	default:
		return c.row, c.col
	}
}

func (c *Cucumber) takeStep(maxRow, maxCol int) (int, int) {
	nextRow, nextCol := c.nextStep(maxRow, maxCol)
	c.row = nextRow
	c.col = nextCol
	return nextRow, nextCol
}
