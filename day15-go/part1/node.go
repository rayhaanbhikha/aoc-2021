package main

import "fmt"

type Node struct {
	key        string
	row        int
	col        int
	val        int
	neighbours []*Node
}

func NewNode(row, col, val int) *Node {
	return &Node{
		key:        fmt.Sprintf("%d:%d", row, col),
		row:        row,
		col:        col,
		val:        val,
		neighbours: make([]*Node, 0),
	}
}

func (n *Node) addNeighbour(neighbour *Node) {
	n.neighbours = append(n.neighbours, neighbour)
}
