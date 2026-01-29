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
	// Helpers

	// распределение ключей по серверам server => count
	type distribution map[ServerName]int

	requiredAllKeysDistributed := func(t *testing.T, d distribution) {
		t.Helper()
		for key, count := range d {
			require.NotZero(t, count, key)
		}
	}

	requiredNumberOfKeysEqual := func(t *testing.T, expect int, d distribution) {
		t.Helper()
		var actual int
		for _, count := range d {
			actual += count
		}
		require.Equal(t, expect, actual)
	}

	requireCompatibility := func(t *testing.T, before, after distribution) {
		t.Helper()
		for server := range before {
			require.GreaterOrEqual(t, before[server], after[server])
		}
	}

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

		distribution := make(distribution)

		keys := randomKeys()

		// Act
		for _, key := range keys {
			server, err := ring.Get(key)
			require.NoError(t, err)
			distribution[server]++
		}

		// Assert
		requiredAllKeysDistributed(t, distribution)
		requiredNumberOfKeysEqual(t, len(keys), distribution)
	})

	t.Run("rebalanced keys after add new server", func(t *testing.T) {
		// Arrange
		ring := NewServerRing()

		require.NoError(t, ring.Add("server1"))
		require.NoError(t, ring.Add("server2"))
		require.NoError(t, ring.Add("server3"))

		keys := randomKeys()

		makeKeyDistribution := func() distribution {
			result := make(map[ServerName]int)
			for _, key := range keys {
				server, err := ring.Get(key)
				require.NoError(t, err)
				result[server]++
			}
			return result
		}

		distributionBefore := makeKeyDistribution()

		// Act
		require.NoError(t, ring.Add("server4"))
		require.NoError(t, ring.Add("server5"))

		// Assert
		requiredAllKeysDistributed(t, distributionBefore)
		requiredNumberOfKeysEqual(t, len(keys), distributionBefore)

		distributionAfter := makeKeyDistribution()
		requiredAllKeysDistributed(t, distributionAfter)
		requiredNumberOfKeysEqual(t, len(keys), distributionAfter)

		requireCompatibility(t, distributionBefore, distributionAfter)

		require.Len(t, distributionBefore, 3)
		require.Len(t, distributionAfter, 5)
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
	const count = 1000
	keys := make([]string, 0, count)
	for range count {
		keys = append(keys, strconv.Itoa(rand.Intn(math.MaxInt)))
	}
	return keys
}
