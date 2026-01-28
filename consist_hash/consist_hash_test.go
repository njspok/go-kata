package consist_hash

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerRingPropertyBased(t *testing.T) {
	t.Run("return same server", func(t *testing.T) {
		ring := NewServerRing()

		require.NoError(t, ring.Add("server1"))
		require.NoError(t, ring.Add("server2"))
		require.NoError(t, ring.Add("server3"))

		server, err := ring.Get("hello")
		require.NoError(t, err)

		for range 100 {
			res, err := ring.Get("hello")
			require.NoError(t, err)
			require.Equal(t, server, res)
		}
	})

	t.Run("return server from server list", func(t *testing.T) {
		ring := NewServerRing()

		servers := []ServerName{"server1", "server2", "server3"}
		for _, server := range servers {
			require.NoError(t, ring.Add(server))
		}

		for i := range 100 {
			key := fmt.Sprintf("hello_no_%d", i)
			server, err := ring.Get(key)
			require.NoError(t, err)
			require.Contains(t, servers, server)
		}
	})

	t.Run("keys distributions by server correctly", func(t *testing.T) {
		// Arrange
		ring := NewServerRing()

		require.NoError(t, ring.Add("server1"))
		require.NoError(t, ring.Add("server2"))
		require.NoError(t, ring.Add("server3"))
		require.NoError(t, ring.Add("server4"))
		require.NoError(t, ring.Add("server5"))

		distribution := make(map[ServerName]int)

		// Act
		for range 100 {
			key := strconv.Itoa(rand.Intn(math.MaxInt))
			server, err := ring.Get(key)
			require.NoError(t, err)
			distribution[server]++
		}

		// Assert
		total := 0
		for key, count := range distribution {
			require.NotZero(t, count, key)
			total += count
		}
		require.Equal(t, 100, total)
	})

	t.Run("rebalanced keys after add new server", func(t *testing.T) {
		// Arrange
		ring := NewServerRing()

		require.NoError(t, ring.Add("server1"))
		require.NoError(t, ring.Add("server2"))
		require.NoError(t, ring.Add("server3"))

		keys := randomKeys()

		makeKeyDistribution := func() map[ServerName]int {
			result := make(map[ServerName]int)
			for _, key := range keys {
				server, err := ring.Get(key)
				require.NoError(t, err)
				result[server]++
			}
			return result
		}

		beforeDistribution := makeKeyDistribution()

		// Act
		require.NoError(t, ring.Add("server4"))
		require.NoError(t, ring.Add("server5"))

		// Assert
		requireNoZeroElements := func(t *testing.T, d map[ServerName]int) {
			t.Helper()
			for key, count := range d {
				require.NotZero(t, count, key)
			}
		}

		requireEqualTotal := func(t *testing.T, expect int, d map[ServerName]int) {
			t.Helper()
			var actual int
			for _, count := range d {
				actual += count
			}
			require.Equal(t, expect, actual)
		}

		requireNoZeroElements(t, beforeDistribution)
		requireEqualTotal(t, len(keys), beforeDistribution)

		afterDistribution := makeKeyDistribution()
		requireNoZeroElements(t, afterDistribution)
		requireEqualTotal(t, len(keys), afterDistribution)

		for server := range beforeDistribution {
			require.GreaterOrEqual(t, beforeDistribution[server], afterDistribution[server])
		}

		require.Len(t, beforeDistribution, 3)
		require.Len(t, afterDistribution, 5)
	})
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

func randomKeys() []string {
	const len = 100
	keys := make([]string, 0, len)
	for range len {
		keys = append(keys, strconv.Itoa(rand.Intn(math.MaxInt)))
	}
	return keys
}
