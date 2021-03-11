package bowling

func NewFrame(number uint) *Frame {
	return &Frame{
		number: number,
		First:  0,
		Second: 0,
		Score:  0,
		Status: FirstRollStatus,
	}
}

// todo count score?
type Frame struct {
	number uint
	First  uint
	Second uint
	Score  uint
	Status Status
}

func (f *Frame) Roll(count uint) {
	switch f.Status {
	case FirstRollStatus:
		f.First = count
		if f.isStrike() {
			f.Status = StrikeStatus
		} else {
			f.Status = SecondRollStatus
		}
	case SecondRollStatus:
		f.Second = count
		if f.isSpare() {
			f.Status = SpareStatus
		} else {
			f.Status = OpenedStatus
		}
	default:
		return
	}
}

func (f *Frame) IsComplete() bool {
	return f.Status == OpenedStatus ||
		f.Status == StrikeStatus ||
		f.Status == SpareStatus
}

func (f *Frame) SetScore(v uint) {
	f.Score = v
}

func (f *Frame) Number() uint {
	return f.number
}

func (f *Frame) isStrike() bool {
	return f.First == 10
}

func (f *Frame) isSpare() bool {
	return (f.First + f.Second) == 10
}
