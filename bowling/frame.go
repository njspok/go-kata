package bowling

func NewFrame(number uint) Frame {
	return Frame{
		Number: number,
		First:  0,
		Second: 0,
		Score:  0,
		Status: FirstRollStatus,
	}
}

type Frame struct {
	Number uint
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

func (f *Frame) isStrike() bool {
	return f.First == 10
}

func (f *Frame) isSpare() bool {
	return (f.First + f.Second) == 10
}
