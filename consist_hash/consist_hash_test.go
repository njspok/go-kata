package consist_hash

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerRing(t *testing.T) {
	ring := NewServerRing()
	ring.Add("server1")
	ring.Add("server2")
	ring.Add("server3")

	s, err := ring.Get("hello")

	require.NoError(t, err)
	require.EqualValues(t, "server3", s)
}

func TestHash(t *testing.T) {
	t.Run("no servers", func(t *testing.T) {
		_, err := getServer(nil, 100)
		require.EqualError(t, err, "no servers")
	})

	t.Run("found servers", func(t *testing.T) {
		tests := []struct {
			servers []ServerNo
			key     uint32
			result  ServerNo
		}{
			{servers: []ServerNo{10}, key: 0, result: 10},
			{servers: []ServerNo{10}, key: 20, result: 10},
			{servers: []ServerNo{10}, key: math.MaxUint32, result: 10},

			{servers: []ServerNo{10, 20}, key: 0, result: 10},
			{servers: []ServerNo{10, 20}, key: 15, result: 20},
			{servers: []ServerNo{10, 20}, key: 25, result: 10},
			{servers: []ServerNo{10, 20}, key: math.MaxUint32, result: 10},
		}

		for _, tt := range tests {
			t.Run("", func(t *testing.T) {
				server, err := getServer(tt.servers, tt.key)
				require.NoError(t, err)
				require.EqualValues(t, int(tt.result), int(server))
			})
		}
	})
}
