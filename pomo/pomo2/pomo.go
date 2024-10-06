package pomo3

import (
	"sync/atomic"
	"time"
)

var PauseEvent = Event{IntervalType: Paused}

type Event struct {
	Duration     time.Duration
	Elapsed      time.Duration
	IntervalType IntervalType
}

type IntervalType string

func (i IntervalType) String() string {
	return string(i)
}

const (
	Work       IntervalType = "work"
	ShortRelax              = "short"
	LongRelax               = "long"
	Paused                  = "paused"
)

const (
	WorkDuration       = time.Minute * 25
	ShortRelaxDuration = time.Minute * 5
	LongRelaxDuration  = time.Minute * 15
)

func IntervalTypeByNumber(number int) IntervalType {
	switch {
	case number%8 == 0:
		return LongRelax
	case number%2 == 0:
		return ShortRelax
	default:
		return Work
	}
}

func New(work, short, long, tick time.Duration) (*Pomo, <-chan Event) {
	p := &Pomo{
		workInterval:       work,
		shortRelaxInterval: short,
		longRelaxInterval:  long,
		notification:       make(chan Event),
		tick:               tick,
		sleep:              time.Sleep,
	}

	go p.run()

	return p, p.notification
}

type Pomo struct {
	workInterval       time.Duration
	shortRelaxInterval time.Duration
	longRelaxInterval  time.Duration

	tick time.Duration

	isPaused atomic.Bool

	notification chan Event

	sleep func(duration time.Duration)
}

func (p *Pomo) run() {
	infinityCycle(p.runInterval)
}

func (p *Pomo) Pause() {
	p.isPaused.Store(true)
}

func (p *Pomo) Stop() {
	panic("unimplemented")
}

func (p *Pomo) Resume() {
	p.isPaused.Store(false)
}

func (p *Pomo) mustMakeEvent(number int) Event {
	var event Event
	switch IntervalTypeByNumber(number) {
	case LongRelax:
		event = Event{IntervalType: LongRelax, Duration: p.longRelaxInterval, Elapsed: 0}
	case ShortRelax:
		event = Event{IntervalType: ShortRelax, Duration: p.shortRelaxInterval, Elapsed: 0}
	case Work:
		event = Event{IntervalType: Work, Duration: p.workInterval, Elapsed: 0}
	default:
		panic("unknown interval type")
	}
	return event
}

func (p *Pomo) runInterval(intervalNumber int) {
	event := p.mustMakeEvent(intervalNumber)

	ticks := int64(event.Duration) / int64(p.tick)
	counter := int64(0)

	for counter <= ticks {
		if p.isPaused.Load() {
			p.notification <- PauseEvent
		} else {
			p.notification <- event
			counter++
			event.Elapsed += p.tick
		}

		p.sleep(p.tick)
	}
}

func infinityCycle(f func(n int)) {
	n := 0
	for {
		n++
		f(n)
	}
}
