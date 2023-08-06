package wait_queue

import "sync"

func NewTask[T any](do Action[T]) *Task[T] {
	t := &Task[T]{
		do: do,
	}
	t.wg.Add(1)
	return t
}

type Task[T any] struct {
	wg     sync.WaitGroup
	do     Action[T]
	Result T
}

func (t *Task[T]) Run() {
	go func() {
		t.Result = t.do()
		t.wg.Done()
	}()
}

func (t *Task[T]) Wait() {
	t.wg.Wait()
}
