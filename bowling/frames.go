package bowling

func NewFrames() *Frames {
	return &Frames{}
}

type Frames struct {
	frames []Framer
}

func (fs *Frames) Add(f Framer) {
	fs.frames = append(fs.frames, f)
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
	return fs.Frame(uint(fs.Len()))
}

func (fs *Frames) Len() int {
	return len(fs.frames)
}
