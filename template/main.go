package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/rayhaanbhikha/aoc-2021/template/language"
)

var codeType language.Language
var day string

func init() {
	codeType = language.GOLANG

	if len(os.Args) < 1 {
		log.Fatal("Missing day")
	}

	day = os.Args[1]

	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "node":
			codeType = language.NODE
		case "go":
			codeType = language.GOLANG
		default:
			log.Fatal("Language unknown")
		}
	}
}

func main() {
	rootDirName := fmt.Sprintf("day%s-%s", day, codeType)

	if _, err := os.Stat(rootDirName); err == nil {
		log.Fatal(fmt.Errorf("%s already exists", rootDirName))
	}

	err := os.Mkdir(rootDirName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	inputData, err := fetchInput(day)
	if err != nil {
		log.Fatal(err)
	}

	switch codeType {
	case language.NODE:
		if err := language.NewNodeLanguage(rootDirName, inputData); err != nil {
			log.Fatal(err)
		}
	case language.GOLANG:
		if err := language.NewGoLanguage(rootDirName, inputData); err != nil {
			log.Fatal(err)
		}
	}
}

func fetchInput(day string) ([]byte, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://adventofcode.com/2021/day/%s/input", day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", os.Getenv("SESSION")))
	response, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
