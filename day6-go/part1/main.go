package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

	fishes := make([]*Fish, 0)
	for _, daysRemaining := range inputs {
		daysRemainingN, _ := strconv.Atoi(daysRemaining)
		fishes = append(fishes, &Fish{daysRemaining: daysRemainingN})
	}

	for i := 0; i < 256; i++ {
		for _, fish := range fishes {
			fish.nextDay()
			if fish.shouldAddFish {
				fishes = append(fishes, &Fish{daysRemaining: 8})
			}
		}
	}

	fmt.Println(len(fishes))
}
