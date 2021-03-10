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

type Status string

const (
	OpenedStatus Status = "opened"
	StrikeStatus Status = "strike"
	SpareStatus  Status = "spare"

	FirstRollStatus  Status = "first"
	SecondRollStatus Status = "second"
	ThirdRollStatus  Status = "third"
	FinalStatus      Status = "final"
)

type Frames []Frame

func NewFinalFrame() FinalFrame {
	return FinalFrame{
		First:  0,
		Second: 0,
		Third:  0,
		Score:  0,
		Status: FirstRollStatus,
	}
}

type FinalFrame struct {
	First  uint
	Second uint
	Third  uint
	Score  uint
	Status Status
}

func (f FinalFrame) IsComplete() bool {
	return f.Status == FinalStatus
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
