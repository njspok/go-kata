package bowling

func NewIterator(frames *Frames) *Iterator {
	return &Iterator{
		frames:  frames,
		current: 1,
	}
}

type Iterator struct {
	frames     *Frames
	current    uint
	onComplete func(Framer)
	onRoll     func(uint)
}

func (i *Iterator) Roll(count uint) {
	i.onRoll(count)
	frame := i.frames.Frame(i.current)
	frame.Roll(count)
	if frame.IsComplete() {
		i.onComplete(frame)
		i.current++
	}
}

func (i *Iterator) SetOnComplete(cb func(Framer)) {
	i.onComplete = cb
}

func (i *Iterator) SetOnRoll(cb func(uint)) {
	i.onRoll = cb
}
