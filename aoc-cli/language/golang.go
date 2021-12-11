package language

import (
	"os"
	"path"
)

type GoLanguage struct {
	basePath string
}

func NewGoLanguage(basePath string, inputData []byte) error {
	language := &GoLanguage{basePath}
	return language.create(inputData)
}

func (g *GoLanguage) create(inputData []byte) error {

	if _, err := os.Stat(g.basePath); err != nil {
		return err
	}

	if err := g.createInputFile(inputData); err != nil {
		return err
	}

	if err := g.createSampleFile(); err != nil {
		return err
	}

	directories := []string{
		path.Join(g.basePath, "part1"),
		path.Join(g.basePath, "part2"),
	}

	for _, dir := range directories {
		if err := g.createDirectory(dir, g.template()); err != nil {
			return err
		}
	}
	return nil
}

func (g *GoLanguage) createDirectory(dirName string, data []byte) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := path.Join(dirName, "main.go")

	return os.WriteFile(filePath, data, os.ModePerm)
}

func (g *GoLanguage) createInputFile(inputData []byte) error {
	return os.WriteFile(path.Join(g.basePath, "input"), inputData, os.ModePerm)
}

func (g *GoLanguage) createSampleFile() error {
	return os.WriteFile(path.Join(g.basePath, "sample"), []byte(""), os.ModePerm)
}

func (g *GoLanguage) template() []byte {
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
