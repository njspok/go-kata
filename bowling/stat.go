package bowling

func NewStat() *Stat {
	var frames Frames
	for i := 0; i < 9; i++ {
		frames = append(frames, NewFrame(uint(i)+1))
	}

	return &Stat{
		frames: frames,
		final:  NewFinalFrame(),
	}
}

type Frames []Frame

type Stat struct {
	frames Frames
	final  FinalFrame
}

func (s *Stat) SetFrame(f Frame) {
	s.frames[f.Number-1] = f
}

func (s *Stat) Frame(number uint) *Frame {
	// todo check range
	return &s.frames[number-1]
}

func (s *Stat) SetFinalFrame(f FinalFrame) {
	s.final = f
}

func (s *Stat) FinalFrame() *FinalFrame {
	return &s.final
}
