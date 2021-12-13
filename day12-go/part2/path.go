package main

type Path struct {
	nodesVisited          map[string]int
	smallCaveVisitedTwice bool
}

func NewPath() *Path {
	return &Path{
		nodesVisited: make(map[string]int),
	}
}

func (p *Path) markNodeAsVisited(node *Node) {
	if node.isBigCave {
		return
	}

	visitedOccurrence, ok := p.nodesVisited[node.label]
	if ok && visitedOccurrence == 1 {
		p.smallCaveVisitedTwice = true
	}

	p.nodesVisited[node.label] = visitedOccurrence + 1
}

func (p *Path) hasBeenVisited(node *Node) bool {
	if node.isBigCave {
		return false
	}

	_, ok := p.nodesVisited[node.label]
	if ok && (node.isStart() || p.smallCaveVisitedTwice) {
		return true
	}

	return false
}

func (p *Path) removeNode(node *Node) {
	if node.isBigCave {
		return
	}

	occurrence, ok := p.nodesVisited[node.label]
	if !ok {
		return
	}

	if occurrence == 2 {
		p.nodesVisited[node.label]--
		p.smallCaveVisitedTwice = false
		return
	}

	delete(p.nodesVisited, node.label)
}
