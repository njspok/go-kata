package bowling

import "errors"

var ErrGameFinished = errors.New("game finished")

func NewBowling() *Bowling {
	return &Bowling{
		total:              0,
		currentFrameNumber: 1,
		frames:             NewStat(),
	}
}

type Bowling struct {
	total              uint
	currentFrameNumber uint
	frames             *Stat // todo rename
	// todo save bonuses
}

func (b *Bowling) Total() uint {
	return b.total
}

func (b *Bowling) Roll(count uint) error {
	if b.IsFinished() {
		return ErrGameFinished
	}

	b.total += count

	if b.currentFrameNumber == 10 {
		frame := b.frames.FinalFrame()
		switch frame.Status {
		case FirstRollStatus:
			frame.First = count
			frame.Status = SecondRollStatus
		case SecondRollStatus:
			frame.Second = count
			frame.Status = OpenedStatus
		}

		frame.Score = b.total
	} else {
		frame := b.frames.Frame(b.currentFrameNumber)
		switch frame.Status {
		case FirstRollStatus:
			frame.First = count
			frame.Status = SecondRollStatus
		case SecondRollStatus:
			frame.Second = count
			frame.Status = OpenedStatus

			b.currentFrameNumber++
		}

		frame.Score = b.total
	}

	return nil
}

func (b *Bowling) IsFinished() bool {
	return b.frames.FinalFrame().IsComplete()
}

func (b *Bowling) Stat() *Stat {
	return b.frames
}
