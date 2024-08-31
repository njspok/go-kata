package pomo

import (
	"errors"
	"time"
)

type IntervalType string

var (
	ErrAlreadyRunning = errors.New("already running")
	ErrNotRunning     = errors.New("not running")
	ErrIncorrectTime  = errors.New("incorrect time")
)

const (
	Work       IntervalType = "work"
	ShortRelax              = "short"
	LongRelax               = "long"
)

const (
	WorkDuration       = time.Minute * 25
	ShortRelaxDuration = time.Minute * 5
	LongRelaxDuration  = time.Minute * 15
	countIteration     = 4
)

func NewWorkInterval(start time.Time) *Interval {
	return &Interval{
		tp:    Work,
		start: start,
		end:   start.Add(WorkDuration),
	}
}

func NewShortRelaxInterval(start time.Time) *Interval {
	return &Interval{
		tp:    ShortRelax,
		start: start,
		end:   start.Add(ShortRelaxDuration),
	}
}

func NewLongRelaxInternal(start time.Time) *Interval {
	return &Interval{
		tp:    LongRelax,
		start: start,
		end:   start.Add(LongRelaxDuration),
	}
}

type Interval struct {
	tp    IntervalType
	start time.Time
	end   time.Time
}

func (i *Interval) Type() IntervalType {
	return i.tp
}

func (i *Interval) Start() time.Time {
	return i.start
}

func (i *Interval) End() time.Time {
	return i.end
}

func (i *Interval) IsEnd(t time.Time) bool {
	return t.After(i.end) || t.Equal(i.end)
}

func New() *Pomo {
	p := &Pomo{}
	p.Reset()
	return p
}

type Pomo struct {
	start time.Time
	isRun bool
}

func (p *Pomo) Start(t time.Time) error {
	if p.isRun {
		return ErrAlreadyRunning
	}

	p.start = t
	p.isRun = true

	return nil
}

func (p *Pomo) Reset() {
	p.start = time.Time{}
	p.isRun = false
}

func (p *Pomo) Interval(t time.Time) (*Interval, error) {
	if !p.isRun {
		return nil, ErrNotRunning
	}

	if t.Before(p.start) && !t.Equal(p.start) {
		return nil, ErrIncorrectTime
	}

	start := p.start

	// pomo life cycle
	for {
		for range countIteration - 1 {
			interval := NewWorkInterval(start)
			if !interval.IsEnd(t) {
				return interval, nil
			}

			interval = NewShortRelaxInterval(interval.end)
			if !interval.IsEnd(t) {
				return interval, nil
			}

			start = interval.end
		}

		interval := NewWorkInterval(start)
		if !interval.IsEnd(t) {
			return interval, nil
		}

		interval = NewLongRelaxInternal(interval.end)
		if !interval.IsEnd(t) {
			return interval, nil
		}

		start = interval.end
	}
}
