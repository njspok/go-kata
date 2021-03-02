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
	if b.hopelessly() {
		return
	}

	if b.working() {
		b.lamp.Switch()
		return
	}

	b.counter++
	switch {
	case b.counter == b.attempts:
		b.lamp.Switch()
	case b.counter > b.attempts:
		b.lamp.Switch()
		b.counter = 1
	}
}

func (b *BrokenLamp) hopelessly() bool {
	return b.attempts == 0
}

func (b *BrokenLamp) working() bool {
	return b.attempts == 1
}
