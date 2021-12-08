package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("../input")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")
	vals := 0
	for _, line := range input {
		result := strings.Split(line, "|")
		fourDigits := strings.Split(strings.TrimSpace(result[1]), " ")
		vals += findNums(fourDigits)
	}
	fmt.Println(vals)
}

func findNums(signalPatterns []string) int {
	num := 0
	for _, pattern := range signalPatterns {
		switch len(pattern) {
		case 2, 3, 4, 7:
			num++
		default:
			continue
		}
	}
	return num
}
