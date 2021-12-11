package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Number struct {
	signal         []string
	signalCharsMap map[string]struct{}
	hasVal         bool
	val            string
}

func NewNumber(signal string) *Number {
	signalsArr := strings.Split(signal, "")
	signalsCharMap := make(map[string]struct{})
	for _, char := range signalsArr {
		signalsCharMap[char] = struct{}{}
	}

	return &Number{
		signal:         signalsArr,
		signalCharsMap: signalsCharMap,
	}
}

func (n *Number) setVal(val string) {
	n.val = val
	n.hasVal = true
}

func (n *Number) signalLen() int {
	return len(n.signal)
}

func (n *Number) getVal(signal string) (string, bool) {
	if len(signal) != len(n.signal) {
		return "", false
	}

	for _, s := range signal {
		if _, ok := n.signalCharsMap[string(s)]; !ok {
			return "", false
		}
	}
	return n.val, true
}

func (n *Number) has(otherNum *Number) bool {
	for char := range otherNum.signalCharsMap {
		if _, ok := n.signalCharsMap[char]; !ok {
			return false
		}
	}
	return true
}

func (n *Number) intersects(otherNum *Number) int {
	charMapCopy := make(map[string]struct{})

	for char := range otherNum.signalCharsMap {
		charMapCopy[char] = struct{}{}
	}

	for char := range n.signalCharsMap {
		charMapCopy[char] = struct{}{}
	}

	return len(charMapCopy)
}

func main() {
	data, _ := ioutil.ReadFile("../sample2")
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	total := 0

	for _, line := range input {
		result := strings.Split(line, "|")
		signals := strings.Split(strings.TrimSpace(result[0]), " ")
		fourDigits := strings.Split(strings.TrimSpace(result[1]), " ")
		total += decodeFourDigits(signals, fourDigits)
	}

	fmt.Println(total)
}

func decodeFourDigits(signals, fourDigitsSignal []string) int {
	uniqueNumbers := parseSignals(signals)

	val := make([]string, 0)
	for _, digitSignal := range fourDigitsSignal {
		for _, uniqueNum := range uniqueNumbers {
			parseNum, ok := uniqueNum.getVal(digitSignal)
			if ok {
				val = append(val, parseNum)
				break
			}
		}
	}
	finalNum, _ := strconv.Atoi(strings.Join(val, ""))
	return finalNum
}

func parseSignals(signals []string) map[string]*Number {
	uniqueNums, remainingNums := generateNumbers(signals)

	filterLengthFive(uniqueNums, remainingNums["5"])
	filterLengthSix(uniqueNums, remainingNums["6"])

	return uniqueNums
}

func filterLengthFive(uniqueNums map[string]*Number, nums []*Number) {
	for _, num := range nums {
		switch {
		case num.has(uniqueNums["1"]):
			num.setVal("3")
		case num.intersects(uniqueNums["4"]) == 7:
			num.setVal("2")
		default:
			num.setVal("5")
		}
		uniqueNums[num.val] = num
	}
}

func filterLengthSix(uniqueNums map[string]*Number, nums []*Number) {
	for _, num := range nums {
		switch {
		case num.intersects(uniqueNums["7"]) == 7:
			num.setVal("6")
		case num.intersects(uniqueNums["4"]) == 7:
			num.setVal("0")
		default:
			num.setVal("9")
		}
		uniqueNums[num.val] = num
	}
}

func generateNumbers(signals []string) (uniqueNums map[string]*Number, remainingNumsByLen map[string][]*Number) {
	uniqueNums = make(map[string]*Number)
	remainingNumsByLen = make(map[string][]*Number)

	for _, signal := range signals {
		num := NewNumber(signal)
		switch len(signal) {
		case 2:
			num.setVal("1")
		case 4:
			num.setVal("4")
		case 3:
			num.setVal("7")
		case 7:
			num.setVal("8")
		}
		if num.hasVal {
			uniqueNums[num.val] = num
		} else {
			key := fmt.Sprintf("%d", num.signalLen())
			remainingNumsByLen[key] = append(remainingNumsByLen[key], num)
		}
	}

	return
}
