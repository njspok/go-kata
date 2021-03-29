package bowling

import "errors"

var ErrGameFinished = errors.New("game finished")

func NewBowling() *Bowling {
	frames := newClassicFrames()
	iterator := NewIterator(frames)
	scorer := NewScorer()

	bowling := &Bowling{
		scorer:   scorer,
		frames:   frames,
		iterator: iterator,
	}

	iterator.SetOnComplete(scorer.Earn)
	iterator.SetOnRoll(scorer.Roll)

	return bowling
}

type Framer interface {
	Number() uint
	Roll(count uint)
	IsComplete() bool
	SetScore(v uint)
}

type Bowling struct {
	scorer   *Scorer
	frames   *Frames
	iterator *Iterator
	// todo save bonuses
}

func (b *Bowling) Total() uint {
	return b.scorer.Total()
}

func (b *Bowling) Roll(count uint) error {
	if b.IsFinished() {
		return ErrGameFinished
	}

	b.iterator.Roll(count)

	return nil
}

func (b *Bowling) IsFinished() bool {
	return b.frames.FinalFrame().IsComplete()
}

func (b *Bowling) Frames() *Frames {
	return b.frames
}

func newClassicFrames() *Frames {
	f := NewFrames()

	for i := 0; i < 9; i++ {
		f.Add(NewFrame(uint(i) + 1))
	}

	// 10 final frame
	f.Add(NewFinalFrame())

	return f
}
