package avg_cons

import (
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for range 10 {
			for _, n := range nodes {
				log.Println("node", n.ID(), n.Val())
			}
			time.Sleep(time.Second)
		}

		cancel()
	}()

	wg := sync.WaitGroup{}
	for _, n := range nodes {
		wg.Go(func() { n.Run(ctx) })
	}
	wg.Wait()
}
