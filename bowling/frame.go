package bowling

func NewFrame(number uint) *Frame {
	return &Frame{
		number:  number,
		first:   0,
		second:  0,
		score:   0,
		status:  FirstRollStatus,
		bonuses: 0,
	}
}

// todo count score?
type Frame struct {
	number  uint
	first   uint
	second  uint
	score   uint
	status  Status
	bonuses uint
}

func (f *Frame) Roll(count uint) {
	switch f.status {
	case FirstRollStatus:
		f.first = count
		if f.isStrike() {
			f.status = StrikeStatus
		} else {
			f.status = SecondRollStatus
		}
	case SecondRollStatus:
		f.second = count
		if f.isSpare() {
			f.status = SpareStatus
		} else {
			f.status = OpenedStatus
		}
	default:
		return
	}
}

func (f *Frame) IsComplete() bool {
	return f.status == OpenedStatus ||
		f.status == StrikeStatus ||
		f.status == SpareStatus
}

func (f *Frame) SetScore(v uint) {
	f.score = v
}

func (f *Frame) Bonuses() uint {
	return f.bonuses
}

func (f *Frame) Number() uint {
	return f.number
}

func (f *Frame) First() uint {
	return f.first
}

func (f *Frame) Second() uint {
	return f.second
}

func (f *Frame) Score() uint {
	return f.score
}

func (f *Frame) Status() Status {
	return f.status
}

func (f *Frame) isStrike() bool {
	return f.first == 10
}

func (f *Frame) isSpare() bool {
	return (f.first + f.second) == 10
}
