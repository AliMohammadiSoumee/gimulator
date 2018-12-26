package agent

import (
	"github.com/alidadar7676/gimulator/types"
)

const (
	inf = int(1e6)
)

type Node interface {
	Neighbor() []Node
	Equal(Node) bool
	Heuristic() int
	Hit(int, Node)
	HasPrice() bool
	GetPos() types.State
}

func Minimax(node Node, depth int, isMax bool) int {
	//log.Println(node.GetPos(), depth, isMax)

	if depth == 0 {
		h := node.Heuristic()
		node.Hit(h, nil)
		//log.Println("depth = 0 ", node.GetPos(), h)
		return h
	}

	if isMax {
		value := -inf
		var bestChild Node
		for _, child := range node.Neighbor() {
			mm := Minimax(child, depth-1, child.HasPrice())
			if value < mm {
				value = mm
				bestChild = child
			}
		}

		//log.Println(node.Neighbor())
		//log.Println(bestChild)
		if bestChild == nil {
			value = node.Heuristic()
		}
		//log.Println("max : ", node.GetPos(), bestChild.GetPos(), value)
		node.Hit(value, bestChild)
		return value
	}

	value := inf
	var bestChild Node
	for _, child := range node.Neighbor() {
		mm := Minimax(child, depth-1, !child.HasPrice())
		if value > mm {
			value = mm
			bestChild = child
		}
	}
	if bestChild == nil {
		value = node.Heuristic()
	}
	//log.Println("min : ", node.GetPos(), bestChild.GetPos(), value)
	node.Hit(value, bestChild)

	return value
}
