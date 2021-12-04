package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

type Language int

func (l Language) String() string {
	switch l {
	case NODE:
		return "node"
	case GOLANG:
		return "go"
	}
	return "unknown"
}

const (
	GOLANG Language = iota
	NODE
)

var codeType Language
var day string

func init() {
	codeType = GOLANG

	if len(os.Args) < 1 {
		log.Fatal("Missing day")
	}

	day = os.Args[1]

	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "node":
			codeType = NODE
		case "go":
			codeType = GOLANG
		}
	}
}

func main() {
	rootDirName := fmt.Sprintf("day%s-%s", day, codeType)

	inputData, err := fetchInput(day)
	if err != nil {
		log.Fatal(err)
	}

	switch codeType {
	case NODE:
		if err := createForNode(rootDirName, inputData); err != nil {
			log.Fatal(err)
		}
	case GOLANG:
		if err := createForGolang(rootDirName, inputData); err != nil {
			log.Fatal(err)
		}
	}
}

func createForNode(rootDirName string, inputData []byte) error {
	defaultFileData := loadNodeJSFile()

	err := os.Mkdir(rootDirName, os.ModePerm)
	if err != nil {
		return err
	}

	files := []string{
		path.Join(rootDirName, "part1.js"),
		path.Join(rootDirName, "part2.js"),
	}

	for _, file := range files {
		if err := os.WriteFile(file, defaultFileData, os.ModePerm); err != nil {
			return err
		}
	}

	return writeInputData(path.Join(rootDirName, "input"), inputData)
}

func createForGolang(rootDirName string, inputData []byte) error {
	defaultFileData := loadDefaultGoFile()

	directories := []string{
		path.Join(rootDirName, "part1"),
		path.Join(rootDirName, "part2"),
	}

	for _, dir := range directories {
		if err := createDirectory(dir, defaultFileData); err != nil {
			return err
		}
	}

	return writeInputData(path.Join(rootDirName, "input"), inputData)
}

func createDirectory(dirName string, data []byte) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := path.Join(dirName, "main.go")

	return os.WriteFile(filePath, data, os.ModePerm)
}

func loadDefaultGoFile() []byte {
	return []byte(`package main
import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println(inputs)
}
	`)
}

func loadNodeJSFile() []byte {
	return []byte(`const { readFileSync } = require("fs");

const data = readFileSync('./input', { encoding: 'utf-8' }).trim().split('\n');
	`)
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

func writeInputData(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, os.ModePerm)
}
