package avg_cons

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHumanReadable(t *testing.T) {
	interval := time.Second

	nodes := []*Node{
		NewNode(0, 100, interval),
		NewNode(1, 200, interval),
		NewNode(2, 300, interval),
		NewNode(3, 400, interval),
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

func Test(t *testing.T) {
	interval := time.Millisecond

	nodes := []*Node{
		NewNode(0, 100, interval),
		NewNode(1, 200, interval),
		NewNode(2, 300, interval),
		NewNode(3, 400, interval),
	}

	nodes[0].AddNeighbor(nodes[1])
	nodes[0].AddNeighbor(nodes[2])
	nodes[1].AddNeighbor(nodes[3])

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)

	wg := sync.WaitGroup{}
	for _, n := range nodes {
		wg.Go(func() { n.Run(ctx) })
	}
	wg.Wait()

	var avg float64
	for _, n := range nodes {
		avg += n.Val() / float64(len(nodes))
	}

	for _, n := range nodes {
		log.Println("node", n.ID(), n.Val())
		require.InDelta(t, avg, n.Val(), 1)
	}
}
