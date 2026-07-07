package avg_cons

import (
	"sync"
	"testing"
)

func TestHumanReadable(t *testing.T) {
	node1 := NewNode(1, 100)
	node2 := NewNode(2, 200)
	node3 := NewNode(3, 300)
	node4 := NewNode(4, 400)

	node1.AddNeighbor(node2)
	node1.AddNeighbor(node3)
	node2.AddNeighbor(node4)

	wg := sync.WaitGroup{}

	wg.Go(node1.Run)
	wg.Go(node2.Run)
	wg.Go(node3.Run)
	wg.Go(node4.Run)

	wg.Wait()
}
