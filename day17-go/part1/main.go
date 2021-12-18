package main

import (
	"fmt"
	"math"
)

type Probe struct {
	minX, maxX int
	minY, maxY int
}

func NewProbe(minX, maxX, minY, maxY int) *Probe {
	return &Probe{
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,
	}
}

func (p *Probe) testTrajectory(xVelocity, yVelocity int) (int, bool) {
	posX, posY := 0, 0
	maxYReached := math.MinInt
	for {

		if posY > maxYReached {
			maxYReached = posY
		}

		if p.inTargetArea(posX, posY) {
			return maxYReached, true
		}

		if p.isOutsideTargetArea(posX, posY) {
			return 0, false
		}

		posX += xVelocity
		posY += yVelocity

		if xVelocity > 0 {
			xVelocity--
		} else if xVelocity < 0 {
			xVelocity++
		}

		yVelocity--
	}
}

func (p *Probe) inTargetArea(x, y int) bool {
	return x >= p.minX && x <= p.maxX && y >= p.minY && y <= p.maxY
}

func (p *Probe) isOutsideTargetArea(x, y int) bool {
	return y < p.minY || x > p.maxX
}

func main() {
	// maxX, minX, minY, maxY := 20, 30, -10, -5
	minX, maxX, minY, maxY := 179, 201, -109, -63
	probe := NewProbe(minX, maxX, minY, maxY)

	maxHeightReached := math.MinInt

	for i := 1; i <= minX; i++ {
		for j := 1; j <= minX; j++ {
			height, ok := probe.testTrajectory(i, j)
			if ok && height > maxHeightReached {
				maxHeightReached = height
			}
		}
	}
	fmt.Println(maxHeightReached)
}
