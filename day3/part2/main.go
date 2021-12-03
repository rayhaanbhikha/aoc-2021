package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	oxygenGeneratorRating := findOxygenGeneratorRating(inputs)
	co2ScrubberRating := findCO2ScrubberRating(inputs)

	gammaRate, _ := strconv.ParseInt(oxygenGeneratorRating, 2, 64)
	epsilonRate, _ := strconv.ParseInt(co2ScrubberRating, 2, 64)

	fmt.Println(epsilonRate * gammaRate)
}

func findOxygenGeneratorRating(inputs []string) string {
	for i := 0; i < len(inputs[0]); i++ {
		zeroBits := 0
		oneBits := 0
		for j := 0; j < len(inputs); j++ {
			num, _ := strconv.Atoi(string(inputs[j][i]))
			if num == 0 {
				zeroBits++
			} else {
				oneBits++
			}
		}
		val := "0"
		if oneBits >= zeroBits {
			val = "1"
		}
		filteredInputs := make([]string, 0)
		for _, input := range inputs {
			if val == string(input[i]) {
				filteredInputs = append(filteredInputs, input)
			}
		}
		inputs = filteredInputs
		if len(filteredInputs) == 1 {
			return filteredInputs[0]
		}
	}

	return inputs[0]
}

func findCO2ScrubberRating(inputs []string) string {
	for i := 0; i < len(inputs[0]); i++ {
		zeroBits := 0
		oneBits := 0
		for j := 0; j < len(inputs); j++ {
			num, _ := strconv.Atoi(string(inputs[j][i]))
			if num == 0 {
				zeroBits++
			} else {
				oneBits++
			}
		}

		val := "1"
		if zeroBits <= oneBits {
			val = "0"
		}
		filteredInputs := make([]string, 0)
		for _, input := range inputs {
			if val == string(input[i]) {
				filteredInputs = append(filteredInputs, input)
			}
		}
		inputs = filteredInputs
		if len(filteredInputs) == 1 {
			return filteredInputs[0]
		}
	}

	return inputs[0]
}
