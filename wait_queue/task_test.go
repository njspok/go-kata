package wait_queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTask(t *testing.T) {
	const count = 10

	task := NewTask(func() Result {
		return 101
	})

	results := make(chan Result, count)

	for i := 0; i < count; i++ {
		go func() {
			task.Wait()
			results <- task.Result
		}()
	}

	task.Run()
	task.Wait()

	require.Equal(t, 101, task.Result)
	for i := 0; i < count; i++ {
		require.Equal(t, 101, <-results)
	}
}
