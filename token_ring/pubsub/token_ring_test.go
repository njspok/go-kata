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

	broker.Publish(Marker{
		To:  1,
		TTL: 100,
	})

	require.Equal(t, 100, counter)
}
