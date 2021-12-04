package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/rayhaanbhikha/aoc-2021/template/language"
)

var languageTemplate string
var day int

func init() {
	flag.StringVar(&languageTemplate, "lang", "go", "Specify language template to generate")
	flag.IntVar(&day, "day", 0, "Advent of code calenday day")

	flag.Parse()

	if day == 0 {
		log.Fatal(errors.New("day 0 does not exist"))
	}

	fmt.Println("Day: ", day)
	fmt.Println("Language: ", languageTemplate)
}

func main() {
	lang, err := language.NewLanguage(languageTemplate)
	if err != nil {
		log.Fatal(err)
	}

	rootDirName := fmt.Sprintf("day%d-%s", day, lang)

	if _, err := os.Stat(rootDirName); err == nil {
		log.Fatal(fmt.Errorf("%s already exists", rootDirName))
	}

	if err = os.Mkdir(rootDirName, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	inputData, err := fetchInput(day)
	if err != nil {
		log.Fatal(err)
	}

	if err := lang.Create(rootDirName, inputData); err != nil {
		log.Fatal(err)
	}
}

func fetchInput(day int) ([]byte, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day)

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
