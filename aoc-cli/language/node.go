package language

import (
	"os"
	"path"
)

type NodeLanguage struct {
	basePath string
}

func NewNodeLanguage(basePath string, inputData []byte) error {
	language := &NodeLanguage{basePath}
	return language.create(inputData)
}

func (n *NodeLanguage) create(inputData []byte) error {

	if _, err := os.Stat(n.basePath); err != nil {
		return err
	}

	if err := n.createInputFile(inputData); err != nil {
		return err
	}

	if err := n.createSampleFile(); err != nil {
		return err
	}

	filePaths := []string{
		path.Join(n.basePath, "part1.js"),
		path.Join(n.basePath, "part2.js"),
	}

	for _, filePath := range filePaths {
		if err := os.WriteFile(filePath, n.template(), os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func (n *NodeLanguage) createInputFile(inputData []byte) error {
	return os.WriteFile(path.Join(n.basePath, "input"), inputData, os.ModePerm)
}

func (n *NodeLanguage) createSampleFile() error {
	return os.WriteFile(path.Join(n.basePath, "sample"), []byte(""), os.ModePerm)
}

func (n *NodeLanguage) template() []byte {
	return []byte(`const { readFileSync } = require("fs");

const data = readFileSync('./input', { encoding: 'utf-8' }).trim().split('\n');
	`)
}
