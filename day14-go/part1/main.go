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
	s := strings.Split(polymerTemplate, "")
	for i := 0; i < len(s)-1; i++ {
		sChar1 := s[i]
		sChar2 := s[i+1]
		key := fmt.Sprintf("%s%s", sChar1, sChar2)
		if elementToAdd, ok := rules[key]; ok {
			remainingS := append([]string{elementToAdd}, s[i+1:]...)
			s = append(s[:i+1], remainingS...)
			i++
		}
	}
	return strings.Join(s, "")
}

func formatInsertionRules(rules string) map[string]string {
	formattedRules := make(map[string]string)
	for _, line := range strings.Split(rules, "\n") {
		res := strings.Split(line, " -> ")
		formattedRules[res[0]] = res[1]
	}
	return formattedRules
}
