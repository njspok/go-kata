package pomo3

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		work := time.Millisecond * 5
		short := time.Millisecond * 2
		long := time.Millisecond * 3
		interval := time.Millisecond
		sut, msg := New(work, short, long, interval)

		var totalDur time.Duration
		sut.sleep = func(d time.Duration) {
			totalDur += d
		}

		exp := []Event{
			{Duration: work, IntervalType: Work, Elapsed: 0},
			{Duration: work, IntervalType: Work, Elapsed: 1 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 2 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 3 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 4 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 5 * time.Millisecond},
			{Duration: short, IntervalType: ShortRelax, Elapsed: 0},
			{Duration: short, IntervalType: ShortRelax, Elapsed: 1 * time.Millisecond},
			{Duration: short, IntervalType: ShortRelax, Elapsed: 2 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 0},
			{Duration: work, IntervalType: Work, Elapsed: 1 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 2 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 3 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 4 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 5 * time.Millisecond},
			{Duration: short, IntervalType: ShortRelax, Elapsed: 0},
			{Duration: short, IntervalType: ShortRelax, Elapsed: 1 * time.Millisecond},
			{Duration: short, IntervalType: ShortRelax, Elapsed: 2 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 0},
			{Duration: work, IntervalType: Work, Elapsed: 1 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 2 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 3 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 4 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 5 * time.Millisecond},
			{Duration: short, IntervalType: ShortRelax, Elapsed: 0},
			{Duration: short, IntervalType: ShortRelax, Elapsed: 1 * time.Millisecond},
			{Duration: short, IntervalType: ShortRelax, Elapsed: 2 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 0},
			{Duration: work, IntervalType: Work, Elapsed: 1 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 2 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 3 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 4 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 5 * time.Millisecond},
			{Duration: long, IntervalType: LongRelax, Elapsed: 0},
			{Duration: long, IntervalType: LongRelax, Elapsed: 1 * time.Millisecond},
			{Duration: long, IntervalType: LongRelax, Elapsed: 2 * time.Millisecond},
			{Duration: long, IntervalType: LongRelax, Elapsed: 3 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 0},
			{Duration: work, IntervalType: Work, Elapsed: 1 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 2 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 3 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 4 * time.Millisecond},
			{Duration: work, IntervalType: Work, Elapsed: 5 * time.Millisecond},
			{Duration: short, IntervalType: ShortRelax, Elapsed: 0},
		}

		events := make([]Event, 0, len(exp))
		for range len(exp) {
			events = append(events, <-msg)
		}

		for n := range exp {
			t.Run(strconv.Itoa(n), func(t *testing.T) {
				require.Equal(t, exp[n], events[n])
			})
		}

		require.EqualValues(t, 44000000, totalDur)
	})
	t.Run("paused", func(t *testing.T) {
		work := time.Millisecond * 5
		short := time.Millisecond * 2
		long := time.Millisecond * 3
		interval := time.Millisecond
		sut, msg := New(work, short, long, interval)

		var count int
		sut.sleep = func(d time.Duration) {
			count++
		}

		// todo check pause events

		events := make([]Event, 0, 10)
		for v := range 10 {
			events = append(events, <-msg)
			switch v {
			case 3:
				sut.Pause()
			case 6:
				sut.Resume()
			}
		}

		require.EqualValues(t, 9, count)
	})
}
