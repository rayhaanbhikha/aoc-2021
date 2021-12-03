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

func compute(inputs []string, inputsFilter func(oneBits, zeroBits, bitIndex int, inputs []string) []string) string {
	for i := 0; i < len(inputs[0]); i++ {
		zeroBits := 0
		oneBits := 0
		for j := 0; j < len(inputs); j++ {
			if inputs[j][i] == '0' {
				zeroBits++
			} else {
				oneBits++
			}
		}
		inputs = inputsFilter(oneBits, zeroBits, i, inputs)
		if len(inputs) == 1 {
			return inputs[0]
		}
	}

	return ""
}

func findOxygenGeneratorRating(inputs []string) string {
	return compute(inputs, func(oneBits, zeroBits, bitIndex int, inputs []string) []string {
		val := "0"
		if oneBits >= zeroBits {
			val = "1"
		}

		filteredInputs := make([]string, 0)

		for _, input := range inputs {
			if val == string(input[bitIndex]) {
				filteredInputs = append(filteredInputs, input)
			}
		}

		return filteredInputs
	})
}

func findCO2ScrubberRating(inputs []string) string {
	return compute(inputs, func(oneBits, zeroBits, bitIndex int, inputs []string) []string {
		val := '1'
		if zeroBits <= oneBits {
			val = '0'
		}
		filteredInputs := make([]string, 0)
		for _, input := range inputs {
			s := rune(input[bitIndex])
			if s == val {
				filteredInputs = append(filteredInputs, input)
			}
		}

		return filteredInputs
	})
}
