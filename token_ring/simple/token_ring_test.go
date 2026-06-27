package simple

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	ring := NewRing()

	counter := 0

	wg := sync.WaitGroup{}
	wg.Go(func() {
		for range 10000 {
			ring.TakeToken()
			counter++
			ring.PutToken()
		}
	})
	wg.Go(func() {
		for range 10000 {
			ring.TakeToken()
			counter++
			ring.PutToken()
		}
	})
	wg.Wait()

	require.Equal(t, 20000, counter)
}
