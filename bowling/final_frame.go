package bowling

func NewFinalFrame() *FinalFrame {
	return &FinalFrame{
		number:  10,
		fist:    0,
		second:  0,
		third:   0,
		score:   0,
		status:  FirstRollStatus,
		bonuses: 0,
	}
}

type FinalFrame struct {
	number  uint
	fist    uint
	second  uint
	third   uint
	score   uint
	status  Status
	bonuses uint
}

func (f *FinalFrame) Roll(count uint) {
	switch f.status {
	case FirstRollStatus:
		f.fist = count
		f.status = SecondRollStatus
	case SecondRollStatus:
		f.second = count
		switch {
		case f.isFirstStrike():
			f.status = ThirdRollStatus
		case f.isSpare():
			f.status = ThirdRollStatus
		default:
			f.status = FinalStatus
		}
	case ThirdRollStatus:
		f.third = count
		f.status = FinalStatus
	default:
		return
	}
}

func (f *FinalFrame) IsComplete() bool {
	return f.status == FinalStatus
}

func (f *FinalFrame) SetScore(v uint) {
	f.score = v
}

func (f *FinalFrame) Number() uint {
	return f.number
}

func (f *FinalFrame) Bonuses() uint {
	return f.bonuses
}

func (f *FinalFrame) First() uint {
	return f.fist
}

func (f *FinalFrame) Second() uint {
	return f.second
}

func (f *FinalFrame) Third() uint {
	return f.third
}

func (f *FinalFrame) Score() uint {
	return f.score
}

func (f *FinalFrame) Status() Status {
	return f.status
}

func (f *FinalFrame) isFirstStrike() bool {
	return f.fist == 10
}

func (f *FinalFrame) isSpare() bool {
	return (f.fist + f.second) == 10
}
