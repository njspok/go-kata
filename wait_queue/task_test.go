package wait_queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTask(t *testing.T) {
	const count = 10

	task := NewTask(func() int {
		return 101
	})

	results := make(chan int, count)

	for range count {
		go func() {
			task.Wait()
			results <- task.Result
		}()
	}

	task.Run()
	task.Wait()

	require.Equal(t, 101, task.Result)
	for range count {
		require.Equal(t, 101, <-results)
	}
}
