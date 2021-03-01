package lamp

func NewCounterWrapper(lamp lamper) *CounterWrapper {
	return &CounterWrapper{
		lamp:    lamp,
		counter: 0,
	}
}

type CounterWrapper struct {
	lamp    lamper
	counter uint
}

func (w *CounterWrapper) Switch() {
	w.counter++
	w.lamp.Switch()
}

func (w *CounterWrapper) Count() uint {
	return w.counter
}

func (w *CounterWrapper) IsLighted() bool {
	return w.lamp.IsLighted()
}
