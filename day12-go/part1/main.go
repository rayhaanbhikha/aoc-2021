package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Node struct {
	isBigCave  bool
	label      string
	neighbours []*Node
}

func NewNode(label string) *Node {
	isBigCave := true
	if strings.ToLower(label) == label {
		isBigCave = false
	}
	return &Node{label: label, isBigCave: isBigCave}
}

func (n *Node) addNeighbour(neighbour *Node) {
	n.neighbours = append(n.neighbours, neighbour)
}

func (n *Node) isEnd() bool {
	return n.label == "end"
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	nodes := make(map[string]*Node)

	for _, connection := range inputs {
		res := strings.Split(connection, "-")
		node1Label := strings.TrimSpace(res[0])
		node2Label := strings.TrimSpace(res[1])

		node1, ok := nodes[node1Label]
		if !ok {
			node1 = NewNode(node1Label)
			nodes[node1Label] = node1
		}

		node2, ok := nodes[node2Label]
		if !ok {
			node2 = NewNode(node2Label)
			nodes[node2Label] = node2
		}

		node1.addNeighbour(node2)
		node2.addNeighbour(node1)
	}

	startNode := nodes["start"]

	paths := computePaths(startNode, make(map[string]struct{}))

	fmt.Println(nodes, paths)
}

func computePaths(node *Node, path map[string]struct{}) int {
	if node.isEnd() {
		return 1
	}

	if _, visited := path[node.label]; visited {
		return 0
	}

	if !node.isBigCave {
		path[node.label] = struct{}{}
	}

	paths := 0

	for _, neighbour := range node.neighbours {
		paths += computePaths(neighbour, path)
	}

	delete(path, node.label)

	return paths
}
