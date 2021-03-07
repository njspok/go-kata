package bowling

import "errors"

var ErrGameFinished = errors.New("game finished")

func NewBowling() *Bowling {
	return &Bowling{
		total:      0,
		rollNumber: 0,
	}
}

type Bowling struct {
	total      uint
	rollNumber uint
	// todo save frames
	// todo save bonuses
}

func (b *Bowling) Total() uint {
	return b.total
}

func (b *Bowling) Roll(count uint) error {
	if b.rollNumber == 20 {
		return ErrGameFinished
	}

	b.total += count
	b.rollNumber++

	return nil
}

func (b *Bowling) IsFinished() bool {
	return b.rollNumber == 20
}
