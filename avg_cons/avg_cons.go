package avg_cons

import (
	"sync"
	"time"
)

func NewNode(number int, val float64) *Node {
	n := &Node{
		number: number,
		val:    val,
	}
	return n
}

type Node struct {
	neighbors []*Node
	number    int
	val       float64
	mu        sync.RWMutex
}

func (n *Node) AddNeighbor(ng *Node) {
	n.neighbors = append(n.neighbors, ng)
}

func (n *Node) Exchange(v float64) float64 {
	old := n.getVal()
	n.updateVal(v)
	return old
}

func (n *Node) Run() {
	for range time.Tick(time.Second) {
		n.interactWithNeighbors()
	}
}

func (n *Node) interactWithNeighbors() {
	for _, ng := range n.neighbors {
		val := ng.Exchange(n.val)
		n.updateVal(val)
	}
}

func (n *Node) updateVal(v float64) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.val = (n.val + v) / 2
}

func (n *Node) getVal() float64 {
	n.mu.RLock()
	defer n.mu.RUnlock()

	return n.val
}
