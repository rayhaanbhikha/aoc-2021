package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func NewGrid(rawInput []string) [][]int {
	grid := make([][]int, 0)
	for _, row := range rawInput {
		gridRow := make([]int, 0)
		for _, rawNum := range strings.Split(row, "") {
			num, _ := strconv.Atoi(rawNum)
			gridRow = append(gridRow, num)
		}
		grid = append(grid, gridRow)
	}
	return grid
}

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

func NewNodesFromGrid(grid [][]int) Nodes {
	nodes := make(Nodes)
	for rowIdx, row := range grid {
		for colIdx := range row {
			nodes.addNode(rowIdx, colIdx, math.MaxInt)
		}
	}
	return nodes
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	grid := NewGrid(inputs)

	nodes := NewNodesFromGrid(grid)

	// rootNode := nodes.rootNode()

	// connect nodes
	for _, node := range nodes {
		translations := computeTranslations(node.row, node.col, len(grid)-1, len(grid[0])-1)
		for _, translation := range translations {
			neighbourNode := nodes.getNode(node.row+translation[1], node.col+translation[0])
			if neighbourNode == nil {
				panic(fmt.Errorf("Node does not exist!. Trying to find at %v from %v", translation, node))
			}
			node.addNeighbour(neighbourNode)
		}
	}

	visitedNodes := dijkstrasAlgo(nodes, grid)
	fmt.Println(visitedNodes[fmt.Sprintf("%d:%d", len(grid)-1, len(grid[0])-1)])
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

func dijkstrasAlgo(nodes Nodes, grid [][]int) map[string]*Node {
	visitedNodes := make(map[string]*Node)
	nodes.rootNode().val = 0
	// needs to priority heap.
	queue := &NodeHeap{nodes.rootNode()}
	heap.Init(queue)
	for queue.Len() != 0 {
		currentNode := heap.Pop(queue).(*Node)
		if _, ok := visitedNodes[currentNode.key]; ok {
			continue
		} else {
			visitedNodes[currentNode.key] = currentNode
		}

		for _, neighbour := range currentNode.neighbours {
			neighbourRiskLevel := grid[neighbour.row][neighbour.col]
			if distance := neighbourRiskLevel + currentNode.val; distance < neighbour.val {
				neighbour.val = distance
			}
			heap.Push(queue, neighbour)
		}

		// queue = append(queue, currentNode.neighbours...)
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
