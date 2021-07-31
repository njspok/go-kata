package wait_queue

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWaitQueue(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		q := NewWaitQueue()
		require.NotNil(t, q)
	})
	t.Run("run", func(t *testing.T) {
		queue := NewWaitQueue()

		result := queue.Run("do-it", func() Result {
			return true
		})

		require.Equal(t, true, result)
	})
	t.Run("run different parallel", func(t *testing.T) {
		queue := NewWaitQueue()

		result := make(chan bool)

		go func() {
			queue.Run("first", func() Result {
				result <- true
				return nil
			})
		}()

		go func() {
			queue.Run("second", func() Result {
				result <- true
				return nil
			})
		}()

		require.True(t, <-result)
		require.True(t, <-result)
	})
	t.Run("run same parallel", func(t *testing.T) {
		queue := NewWaitQueue()

		result := make(chan Result)

		var counter int32

		do := func() {
			result <- queue.Run("first", func() Result {
				atomic.AddInt32(&counter, 1)
				time.Sleep(1 * time.Second)
				return true
			})
		}

		const threads = 7000

		for i := 0; i < threads; i++ {
			go do()
		}

		for i := 0; i < threads; i++ {
			require.Equal(t, true, <-result)
		}

		require.Equal(t, int32(1), counter)
		require.Nil(t, queue.Load("first"))
	})
}
