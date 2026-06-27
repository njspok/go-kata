package _func

func NewProc(do func(), in <-chan struct{}) chan struct{} {
	out := make(chan struct{})

	go func() {
		defer close(out)
		for range in {
			do()
			out <- struct{}{}
		}
	}()

	return out
}
