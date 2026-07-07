package avg_cons

import (
	"context"
	"sync"
	"time"
)

func NewNode(id int, val float64) *Node {
	n := &Node{
		id:  id,
		val: val,
	}
	return n
}

type Node struct {
	neighbors []*Node
	id        int
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

func (n *Node) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
			n.interactWithNeighbors()
		}
	}
}

func (n *Node) ID() int {
	return n.id
}

func (n *Node) Val() float64 {
	return n.getVal()
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
