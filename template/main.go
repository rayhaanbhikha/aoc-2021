package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

type Language int

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
	rootDirName := fmt.Sprintf("day%s", day)
	switch codeType {
	case NODE:
		if err := createForNode(rootDirName); err != nil {
			log.Fatal(err)
		}
	case GOLANG:
		if err := createForGolang(rootDirName); err != nil {
			log.Fatal(err)
		}
	}
}

func createForNode(rootDirName string) error {
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

	return nil
}

func createForGolang(rootDirName string) error {
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

	return nil
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
	inputs := strings.Split(string(data), "\n")

	fmt.Println(inputs)
}
	`)
}

func loadNodeJSFile() []byte {
	return []byte(`const { readFileSync } = require("fs");

const data = readFileSync('./input', { encoding: 'utf-8' }).trim().split('\n');
	`)
}
