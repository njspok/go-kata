package lamp

func NewBrokenLamp(attempts uint) *BrokenLamp {
	return &BrokenLamp{
		counter:  0,
		attempts: attempts,
		lamp:     NewWorkLamp(),
	}
}

type BrokenLamp struct {
	attempts uint
	counter  uint
	lamp     *WorkLamp
}

func (b *BrokenLamp) IsLighted() bool {
	return b.lamp.IsLighted()
}

func (b *BrokenLamp) Switch() {
	b.counter++
	switch {
	case b.counter == b.attempts:
		b.lamp.Switch()
	case b.counter > b.attempts:
		b.lamp.Switch()
		b.counter = 1
	}
}
