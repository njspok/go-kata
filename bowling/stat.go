package bowling

func NewStat() *Stat {
	frames := make(Frames, 9)
	for i := 0; i < 9; i++ {
		frames[i].Number = uint(i) + 1
		frames[i].Status = FirstRollStatus
	}

	return &Stat{
		frames: frames,
		final: FinalFrame{
			First:  0,
			Second: 0,
			Third:  0,
			Status: FirstRollStatus,
		},
	}
}

type Status string

const (
	OpenedStatus     Status = "opened"
	StrikeStatus     Status = "strike"
	SpareStatus      Status = "spare"
	FirstRollStatus  Status = "first"
	SecondRollStatus Status = "second"
	ThirdRollStatus  Status = "third"
)

type Frame struct {
	Number uint
	First  uint
	Second uint
	Score  uint
	Status Status
}

type Frames []Frame

type FinalFrame struct {
	First  uint
	Second uint
	Third  uint
	Score  uint
	Status Status
}

func (f FinalFrame) IsComplete() bool {
	return f.Status == OpenedStatus ||
		f.Status == StrikeStatus ||
		f.Status == SpareStatus
}

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
