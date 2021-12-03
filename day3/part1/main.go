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

	gammaRateVals := make([]string, 0)
	epsilonRateVals := make([]string, 0)

	for i := 0; i < len(inputs[0]); i++ {
		sum := 0
		for j := 0; j < len(inputs); j++ {
			num, _ := strconv.Atoi(string(inputs[j][i]))
			sum += num
		}
		if sum > (len(inputs) / 2) {
			gammaRateVals = append(gammaRateVals, "1")
			epsilonRateVals = append(epsilonRateVals, "0")
		} else {
			gammaRateVals = append(gammaRateVals, "0")
			epsilonRateVals = append(epsilonRateVals, "1")
		}
	}

	gammaRate, _ := strconv.ParseInt(strings.Join(gammaRateVals, ""), 2, 64)
	epsilonRate, _ := strconv.ParseInt(strings.Join(epsilonRateVals, ""), 2, 64)

	fmt.Println(epsilonRate * gammaRate)
}
