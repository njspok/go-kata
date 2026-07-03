package pubsub

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	counter := 0

	broker := NewLightBroker()

	node1 := NewNode(1, func() {
		counter++
	}, broker)

	node2 := NewNode(2, func() {
		counter++
	}, broker)

	node3 := NewNode(3, func() {
		counter++
	}, broker)

	node1.SetNextNodeID(node2.ID())
	node2.SetNextNodeID(node3.ID())
	node3.SetNextNodeID(node1.ID())

	broker.Add(node1)
	broker.Add(node2)
	broker.Add(node3)

	const ttl = 10_000

	broker.Publish(Marker{
		To:  1,
		TTL: ttl,
	})

	require.Equal(t, ttl, counter)
}
