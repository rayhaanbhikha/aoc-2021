package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Nodes map[string]*Node

func (n Nodes) addNode(rowIdx, colIdx, riskLevel int) {
	key := fmt.Sprintf("%d:%d", rowIdx, colIdx)
	n[key] = NewNode(rowIdx, colIdx, riskLevel)
}

func (n Nodes) getNode(rowIdx, colIdx int) *Node {
	key := fmt.Sprintf("%d:%d", rowIdx, colIdx)
	return n[key]
}

func (n Nodes) rootNode() *Node {
	return n.getNode(0, 0)
}

func NewNodesFromGrid(grid *Grid) Nodes {
	nodes := make(Nodes)
	for rowIdx := 0; rowIdx <= grid.maxRow; rowIdx++ {
		for colIdx := 0; colIdx <= grid.maxCol; colIdx++ {
			nodes.addNode(rowIdx, colIdx, math.MaxInt)
		}
	}
	return nodes
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	grid := NewGrid(inputs, 5)

	// grid.print()
	nodes := NewNodesFromGrid(grid)

	// connect nodes
	for _, node := range nodes {
		translations := computeTranslations(node.row, node.col, grid.maxRow, grid.maxCol)
		for _, translation := range translations {
			neighbourNode := nodes.getNode(node.row+translation[1], node.col+translation[0])
			if neighbourNode == nil {
				panic(fmt.Errorf("Node does not exist!. Trying to find at %v from %v", translation, node))
			}
			node.addNeighbour(neighbourNode)
		}
	}

	visitedNodes := dijkstrasAlgo(nodes, grid)
	fmt.Println(grid.maxCol, grid.maxRow)
	fmt.Println(visitedNodes[fmt.Sprintf("%d:%d", grid.maxRow, grid.maxCol)].val)
}

type NodeHeap []*Node

func (n NodeHeap) Len() int {
	return len(n)
}

func (n NodeHeap) Less(i, j int) bool {
	return n[i].val < n[j].val
}

func (n NodeHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *NodeHeap) Push(x interface{}) {
	*n = append(*n, x.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func dijkstrasAlgo(nodes Nodes, grid *Grid) map[string]*Node {
	visitedNodes := make(map[string]*Node)
	nodes.rootNode().val = 0
	queue := &NodeHeap{nodes.rootNode()}
	heap.Init(queue)
	for queue.Len() != 0 {
		currentNode := heap.Pop(queue).(*Node)
		// lastNodeKey := fmt.Sprintf("%d:%d", grid.maxRow, grid.maxCol)

		if _, ok := visitedNodes[currentNode.key]; ok {
			continue
		} else {
			visitedNodes[currentNode.key] = currentNode
		}

		// if currentNode.key == lastNodeKey {
		// 	break
		// }

		for _, neighbour := range currentNode.neighbours {
			neighbourRiskLevel := grid.getVal(neighbour.row, neighbour.col)
			if distance := neighbourRiskLevel + currentNode.val; distance < neighbour.val {
				neighbour.val = distance
			}
			heap.Push(queue, neighbour)
		}
	}
	return visitedNodes
}

func computeTranslations(row, col, maxRow, maxCol int) []Translation {
	translations := make([]Translation, 0)

	isTopLeft := row == 0 && col == 0
	isTopRight := row == 0 && col == maxCol
	isTopRow := row == 0
	isBottomRow := row == maxRow
	isLeftCol := col == 0
	isRightCol := col == maxCol
	isBottomLeft := row == maxRow && col == 0
	isBottomRight := row == maxRow && col == maxCol

	switch {
	case isTopRight:
		translations = append(translations, S, W)
	case isTopLeft:
		translations = append(translations, E, S)
	case isBottomRight:
		translations = append(translations, N, W)
	case isBottomLeft:
		translations = append(translations, N, E)
	case isTopRow:
		translations = append(translations, E, S, W)
	case isBottomRow:
		translations = append(translations, E, N, W)
	case isLeftCol:
		translations = append(translations, N, E, S)
	case isRightCol:
		translations = append(translations, N, W, S)
	default:
		translations = append(translations, N, E, S, W)
	}

	return translations
}

type Translation [2]int

var (
	N = Translation{0, -1}
	E = Translation{1, 0}
	S = Translation{0, 1}
	W = Translation{-1, 0}
)
