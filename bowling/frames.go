package bowling

func NewFrames() *Frames {
	var frames []Framer
	for i := 0; i < 9; i++ {
		frames = append(frames, NewFrame(uint(i)+1))
	}

	// 10 final frame
	frames = append(frames, NewFinalFrame())

	return &Frames{
		frames: frames,
	}
}

type Frames struct {
	frames []Framer
}

func (fs *Frames) SetFrame(f Framer) {
	fs.frames[f.Number()-1] = f
}

func (fs *Frames) Frame(number uint) Framer {
	// todo check range
	return fs.frames[number-1]
}

func (fs *Frames) SetFinalFrame(f Framer) {
	// todo little lie
	fs.SetFrame(f)
}

func (fs *Frames) FinalFrame() Framer {
	return fs.Frame(10)
}

func (fs *Frames) Len() int {
	return len(fs.frames)
}
