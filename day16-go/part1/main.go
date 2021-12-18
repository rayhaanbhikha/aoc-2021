package main

import (
	"fmt"
	"io/ioutil"
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
		p.parseNumber(binaryString)
		// fmt.Println("num parse: ", num)
	default:
		p.parseOperator(binaryString)
	}

	return p.version
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

func (p *Packet) parseOperator(binaryString string) {
	p.parseLengthTypeIP(binaryString)

	if p.lengthTypeID == 0 {
		// next 15 bits is a number that == total length in bits of the sub packets contained in this packet.
		length, _ := strconv.ParseInt(binaryString[p.offset:p.offset+15], 2, 64)
		p.offset += 15
		maxSubPacketLength := p.offset + int(length)
		for p.offset < maxSubPacketLength {
			packet := &Packet{offset: p.offset}
			p.version += packet.parse(binaryString)
			p.offset = packet.offset
		}
	} else if p.lengthTypeID == 1 {
		length, _ := strconv.ParseInt(binaryString[p.offset:p.offset+11], 2, 64)
		p.offset += 11

		for i := 0; i < int(length); i++ {
			packet := &Packet{offset: p.offset}
			p.version += packet.parse(binaryString)
			p.offset = packet.offset
		}
	}
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
	// inputs := "D2FE28"
	// inputs := "EE00D40C823060"
	// inputs := "8A004A801A8002F478"
	// inputs := "620080001611562C8802118E34"
	// inputs := "C0015000016115A2E0802F182340"
	// inputs := "A0016C880162017C3686B18A3D4780"
	// inputs := "38006F45291200"
	// inputs := "EE00D40C823060"
	val, _ := hexToBin(inputs)
	p := &Packet{}
	fmt.Println(p.parse(val))
	fmt.Println(val)
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
