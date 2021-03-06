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

	oxygenRate, _ := strconv.ParseInt(oxygenGeneratorRating, 2, 64)
	co2ScrubberRate, _ := strconv.ParseInt(co2ScrubberRating, 2, 64)

	fmt.Println(co2ScrubberRate * oxygenRate)
}

func findOxygenGeneratorRating(inputs []string) string {
	return compute(inputs, func(oneBits, zeroBits int) rune {
		if oneBits >= zeroBits {
			return '1'
		}
		return '0'
	})
}

func findCO2ScrubberRating(inputs []string) string {
	return compute(inputs, func(oneBits, zeroBits int) rune {
		if zeroBits <= oneBits {
			return '0'
		}
		return '1'
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
