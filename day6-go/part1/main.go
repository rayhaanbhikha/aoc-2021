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

func (f *Fishes) add(fish *Fish) {
	f.fishes = append(f.fishes, fish)
}

func (f *Fishes) addNewFish() {
	newFish := &Fish{daysRemaining: 8}
	f.add(newFish)
}

func (f *Fishes) nextDay() {
	for _, fish := range f.fishes {
		fish.nextDay()
		if fish.shouldAddFish {
			f.addNewFish()
		}
	}
}

type Fish struct {
	daysRemaining int
	shouldAddFish bool
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

	fishes := &Fishes{}
	for _, daysRemaining := range inputs {
		daysRemainingN, _ := strconv.Atoi(daysRemaining)
		fishes.add(&Fish{daysRemaining: daysRemainingN})
	}

	for i := 0; i < 80; i++ {
		fishes.nextDay()
	}

	fmt.Println(len(fishes.fishes))
}
