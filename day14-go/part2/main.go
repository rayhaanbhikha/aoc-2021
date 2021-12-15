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

	pairs := make(map[string]*Pair)

	for i := 0; i < len(polymnerTemplate)-1; i++ {
		indexToUse := 0
		if i == len(polymnerTemplate)-2 {
			indexToUse = -1
		}
		pair := NewPair(polymnerTemplate[i:i+2], indexToUse, 1)
		pairs[pair.val] = pair
	}

	// printPairs(pairs)
	// fmt.Println("=====\n")

	for i := 0; i < 40; i++ {
		pairs = implementStep(pairs, pairInsertionRules)
		// printPairs(pairs)
		// min, max := computeScore(pairs)
		// fmt.Println(max - min)
		// fmt.Println("=====")
	}

	min, max := computeScore(pairs)

	fmt.Println(max - min)
}

func printPairs(pairs map[string]*Pair) {
	for k, v := range pairs {
		fmt.Println(k, v)
	}
}

func implementStep(pairs map[string]*Pair, rules map[string]string) map[string]*Pair {
	newPairs := make(map[string]*Pair)
	for pair, v := range pairs {
		if elementToAdd, ok := rules[v.val]; ok {
			newPair1 := string(pair[0]) + string(elementToAdd)
			newPair2 := string(elementToAdd) + string(pair[1])

			key1 := fmt.Sprintf("%s:%d", newPair1, 0)
			if val, ok := newPairs[key1]; ok {
				val.incrementCountBy(v.count)
			} else {
				newPairs[key1] = NewPair(newPair1, 0, v.count)
			}

			key2 := fmt.Sprintf("%s:%d", newPair2, v.indexToUse)
			if val, ok := newPairs[key2]; ok {
				val.incrementCountBy(v.count)
			} else {
				newPairs[key2] = NewPair(newPair2, v.indexToUse, v.count)
			}
		}
	}

	return newPairs
}

type Pair struct {
	val        string
	count      int
	indexToUse int
}

func NewPair(val string, indexToUse, count int) *Pair {
	return &Pair{
		val:        val,
		count:      count,
		indexToUse: indexToUse,
	}
}

func (p *Pair) incrementCountBy(n int) {
	p.count += n
}

func (p *Pair) String() string {
	return fmt.Sprintf("\tpair: %s, indexToUse: %d, count: %d", p.val, p.indexToUse, p.count)
}

func computeScore(pairs map[string]*Pair) (int, int) {
	min := math.MaxInt
	max := 0

	occurrenceMap := make(map[string]int)

	for _, pair := range pairs {
		v1 := string(pair.val[0])
		v2 := string(pair.val[1])
		occurrenceMap[v1] = occurrenceMap[v1] + pair.count
		if pair.indexToUse == -1 {
			occurrenceMap[v2] = occurrenceMap[v2] + pair.count
		}
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

func formatInsertionRules(rules string) map[string]string {
	formattedRules := make(map[string]string)
	for _, line := range strings.Split(rules, "\n") {
		res := strings.Split(line, " -> ")
		formattedRules[res[0]] = res[1]
	}
	return formattedRules
}
