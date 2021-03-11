package bowling

import "errors"

var ErrGameFinished = errors.New("game finished")

func NewBowling() *Bowling {
	return &Bowling{
		total:              0,
		currentFrameNumber: 1,
		frames:             NewFrames(),
	}
}

type Framer interface {
	Number() uint
	Roll(count uint)
	IsComplete() bool
	SetScore(v uint)
}

type Bowling struct {
	total              uint // todo who count total?
	currentFrameNumber uint
	frames             *Frames
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

	frame := b.frames.Frame(b.currentFrameNumber)
	frame.Roll(count)
	if frame.IsComplete() {
		frame.SetScore(b.total)
		b.currentFrameNumber++
	}

	return nil
}

func (b *Bowling) IsFinished() bool {
	return b.frames.FinalFrame().IsComplete()
}

func (b *Bowling) Frames() *Frames {
	return b.frames
}
