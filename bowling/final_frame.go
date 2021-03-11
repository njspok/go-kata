package bowling

func NewFinalFrame() *FinalFrame {
	return &FinalFrame{
		number: 10,
		First:  0,
		Second: 0,
		Third:  0,
		Score:  0,
		Status: FirstRollStatus,
	}
}

type FinalFrame struct {
	number uint
	First  uint
	Second uint
	Third  uint
	Score  uint
	Status Status
}

func (f *FinalFrame) Roll(count uint) {
	switch f.Status {
	case FirstRollStatus:
		f.First = count
		f.Status = SecondRollStatus
	case SecondRollStatus:
		f.Second = count
		switch {
		case f.isFirstStrike():
			f.Status = ThirdRollStatus
		case f.isSpare():
			f.Status = ThirdRollStatus
		default:
			f.Status = FinalStatus
		}
	case ThirdRollStatus:
		f.Third = count
		f.Status = FinalStatus
	default:
		return
	}
}

func (f *FinalFrame) IsComplete() bool {
	return f.Status == FinalStatus
}

func (f *FinalFrame) SetScore(v uint) {
	f.Score = v
}

func (f *FinalFrame) Number() uint {
	return f.number
}

func (f *FinalFrame) isFirstStrike() bool {
	return f.First == 10
}

func (f *FinalFrame) isSpare() bool {
	return (f.First + f.Second) == 10
}
