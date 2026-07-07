package avg_cons

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestHumanReadable(t *testing.T) {
	nodes := []*Node{
		NewNode(0, 100),
		NewNode(1, 200),
		NewNode(2, 300),
		NewNode(3, 400),
	}

	nodes[0].AddNeighbor(nodes[1])
	nodes[0].AddNeighbor(nodes[2])
	nodes[1].AddNeighbor(nodes[3])

	wg := sync.WaitGroup{}
	for _, n := range nodes {
		wg.Go(n.Run)
	}

	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return
			case <-time.After(1 * time.Second):
				for _, n := range nodes {
					log.Println("node", n.ID(), n.Val())
				}
			}
		}
	}()

	wg.Wait()
	close(done)
}
