package consist_hash

import (
	"errors"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func get(servers []uint32, key uint32) (uint32, error) {
	if len(servers) == 0 {
		return 0, errors.New("no servers")
	}

	var res uint32
	found := false

	for _, n := range servers {
		if n >= key {
			res = n
			found = true
			break
		}
	}

	if !found {
		res = servers[0]
	}

	return res, nil
}

func Test(t *testing.T) {
	t.Run("no servers", func(t *testing.T) {
		_, err := get(nil, 100)
		require.EqualError(t, err, "no servers")
	})

	t.Run("found servers", func(t *testing.T) {
		tests := []struct {
			servers []uint32
			key     uint32
			result  uint32
		}{
			{servers: []uint32{10}, key: 0, result: 10},
			{servers: []uint32{10}, key: 20, result: 10},
			{servers: []uint32{10}, key: math.MaxUint32, result: 10},

			{servers: []uint32{10, 20}, key: 0, result: 10},
			{servers: []uint32{10, 20}, key: 15, result: 20},
			{servers: []uint32{10, 20}, key: 25, result: 10},
			{servers: []uint32{10, 20}, key: math.MaxUint32, result: 10},
		}

		for _, tt := range tests {
			t.Run("", func(t *testing.T) {
				server, err := get(tt.servers, tt.key)
				require.NoError(t, err)
				require.EqualValues(t, int(tt.result), int(server))
			})
		}
	})
}
