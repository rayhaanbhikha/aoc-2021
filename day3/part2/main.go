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
	return compute(inputs, func(oneBits, zeroBits int) rune {
		val := '0'
		if oneBits >= zeroBits {
			val = '1'
		}
		return val
	})
}

func findCO2ScrubberRating(inputs []string) string {
	return compute(inputs, func(oneBits, zeroBits int) rune {
		val := '1'
		if zeroBits <= oneBits {
			val = '0'
		}
		return val
	})
}

func compute(inputs []string, computeVal func(oneBits, zeroBits int) rune) string {
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
		val := computeVal(oneBits, zeroBits)
		inputs = filterInputsByBit(inputs, i, val)
		if len(inputs) == 1 {
			return inputs[0]
		}
	}

	return ""
}

func filterInputsByBit(inputs []string, bitIndex int, val rune) []string {
	filteredInputs := make([]string, 0)

	for _, input := range inputs {
		if val == rune(input[bitIndex]) {
			filteredInputs = append(filteredInputs, input)
		}
	}

	return filteredInputs
}
