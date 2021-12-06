package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Fishes struct {
	fishes []*Fish
}

func (f *Fishes) load(input []string) {
	fishesSeen := make(map[int]*Fish)
	for _, daysRemaining := range input {
		daysRemainingN, _ := strconv.Atoi(daysRemaining)
		if fish, ok := fishesSeen[daysRemainingN]; ok {
			fish.incrementCopy()
		} else {
			newFish := NewFish(daysRemainingN, 1)
			fishesSeen[daysRemainingN] = newFish
			f.add(newFish)
		}
	}
}

func (f *Fishes) add(fish *Fish) {
	f.fishes = append(f.fishes, fish)
}

func (f *Fishes) nextDay() {
	fishesToSpawn := 0
	for _, fish := range f.fishes {
		fish.nextDay()
		if fish.shouldAddFish {
			fishesToSpawn += fish.copies
		}
	}
	if fishesToSpawn > 0 {
		f.add(NewFish(8, fishesToSpawn))
	}
}

type Fish struct {
	daysRemaining int
	shouldAddFish bool
	copies        int
}

func NewFish(daysRemaining, copies int) *Fish {
	return &Fish{daysRemaining: daysRemaining, copies: copies}
}

func (f *Fish) incrementCopy() {
	f.copies++
}

func (f *Fish) nextDay() {
	if f.daysRemaining == 0 {
		f.daysRemaining = 6
		f.shouldAddFish = true
	} else {
		f.daysRemaining -= 1
		f.shouldAddFish = false
	}
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), ",")

	days := 256

	fishes := &Fishes{}
	fishes.load(inputs)

	for i := 0; i < days; i++ {
		fishes.nextDay()
	}

	sum := 0
	for _, fish := range fishes.fishes {
		sum += fish.copies
	}

	fmt.Println(sum)
}
