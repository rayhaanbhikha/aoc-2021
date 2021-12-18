package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Packet struct {
	offset       int
	version      int
	typeID       int
	lengthTypeID int
}

func (p *Packet) parse(binaryString string) int {
	p.parseVersion(binaryString)
	p.parseTypeIP(binaryString)

	switch p.typeID {
	case 4: // Literal
		num, _ := p.parseNumber(binaryString)
		return int(num)
	default:
		return p.parseOperator(binaryString)
	}
}

func (p *Packet) parseVersion(binaryString string) {
	version, _ := strconv.ParseInt(binaryString[p.offset:p.offset+3], 2, 64)
	p.version = int(version)
	p.offset += 3
	// fmt.Println("version: ", version)
}

func (p *Packet) parseTypeIP(binaryString string) {
	typeID, _ := strconv.ParseInt(binaryString[p.offset:p.offset+3], 2, 64)
	p.typeID = int(typeID)
	p.offset += 3
}

func (p *Packet) parseLengthTypeIP(binaryString string) {
	lengthTypeID, _ := strconv.ParseInt(binaryString[p.offset:p.offset+1], 2, 64)
	p.offset++
	p.lengthTypeID = int(lengthTypeID)
}

func (p *Packet) parseOperator(binaryString string) int {
	p.parseLengthTypeIP(binaryString)

	subPacketResults := make([]int, 0)

	if p.lengthTypeID == 0 {
		// next 15 bits is a number that == total length in bits of the sub packets contained in this packet.
		length, _ := strconv.ParseInt(binaryString[p.offset:p.offset+15], 2, 64)
		p.offset += 15
		maxSubPacketLength := p.offset + int(length)
		for p.offset < maxSubPacketLength {
			packet := &Packet{offset: p.offset}
			subPacketResults = append(subPacketResults, packet.parse(binaryString))
			p.offset = packet.offset
		}
	} else if p.lengthTypeID == 1 {
		length, _ := strconv.ParseInt(binaryString[p.offset:p.offset+11], 2, 64)
		p.offset += 11
		for i := 0; i < int(length); i++ {
			packet := &Packet{offset: p.offset}
			subPacketResults = append(subPacketResults, packet.parse(binaryString))
			p.offset = packet.offset
		}
	}
	return p.computeSubPacketResults(subPacketResults)
}

func (p *Packet) computeSubPacketResults(subPacketResults []int) int {
	switch p.typeID {
	case 0:
		sum := 0
		for _, res := range subPacketResults {
			sum += res
		}
		return sum
	case 1:
		sum := 1
		for _, res := range subPacketResults {
			sum *= res
		}
		return sum
	case 2:
		result := math.MaxInt
		for _, res := range subPacketResults {
			if res < result {
				result = res
			}
		}
		return result
	case 3:
		result := math.MinInt
		for _, res := range subPacketResults {
			if res > result {
				result = res
			}
		}
		return result
	case 5:
		sp1, sp2 := subPacketResults[0], subPacketResults[1]
		if sp1 > sp2 {
			return 1
		}
		return 0
	case 6:
		sp1, sp2 := subPacketResults[0], subPacketResults[1]
		if sp1 < sp2 {
			return 1
		}
		return 0
	case 7:
		sp1, sp2 := subPacketResults[0], subPacketResults[1]
		if sp1 == sp2 {
			return 1
		}
		return 0
	}
	return 0
}

func (p *Packet) parseNumber(binaryString string) (int64, error) {
	var numToParse strings.Builder
	for i := p.offset; ; i += 5 {
		p.offset += 5
		leadingBit := string(binaryString[i])
		valueBits := binaryString[i+1 : i+5]
		numToParse.WriteString(valueBits)
		if leadingBit == "0" {
			break
		}
	}

	return strconv.ParseInt(numToParse.String(), 2, 64)
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.TrimSpace(string(data))

	// inputs := "C200B40A82"
	// inputs := "04005AC33890"

	val, _ := hexToBin(inputs)
	p := &Packet{}
	fmt.Println(p.parse(val))
	// fmt.Println(val)
}

func hexToBin(hex string) (string, error) {
	binaryStrings := [16]string{
		"0000", "0001", "0010", "0011", "0100", "0101", "0110", "0111",
		"1000", "1001", "1010", "1011", "1100", "1101", "1110", "1111"}

	var b strings.Builder
	for _, char := range hex {
		index, err := strconv.ParseUint(string(char), 16, 64)
		if err != nil {
			return "", err
		}

		if index >= uint64(len(binaryStrings)) {
			return "", fmt.Errorf("can't parse index %b", index)
		}

		b.WriteString(binaryStrings[index])
	}

	// %016b indicates base 2, zero padded, with 16 characters
	return b.String(), nil
}
