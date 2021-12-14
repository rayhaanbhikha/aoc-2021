package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	polymnerTemplate := inputs[0]
	pairInsertionRules := formatInsertionRules(inputs[1])

	for i := 0; i < 10; i++ {
		polymnerTemplate = implementStep(polymnerTemplate, pairInsertionRules)
	}
	min, max := computeScore(polymnerTemplate)

	fmt.Println(max - min)
}

func computeScore(polymerTemplate string) (int, int) {
	min := math.MaxInt
	max := 0

	occurrenceMap := make(map[rune]int)

	for _, char := range polymerTemplate {
		val := occurrenceMap[char]
		occurrenceMap[char] = val + 1
	}

	for _, v := range occurrenceMap {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max
}

func implementStep(polymerTemplate string, rules map[string]string) string {
	for i := 0; i < len(polymerTemplate)-1; i++ {
		key := polymerTemplate[i : i+2]
		if elementToAdd, ok := rules[key]; ok {
			polymerTemplate = polymerTemplate[:i+1] + elementToAdd + polymerTemplate[i+1:]
			i++
		}
	}
	return polymerTemplate
}

func formatInsertionRules(rules string) map[string]string {
	formattedRules := make(map[string]string)
	for _, line := range strings.Split(rules, "\n") {
		res := strings.Split(line, " -> ")
		formattedRules[res[0]] = res[1]
	}
	return formattedRules
}
