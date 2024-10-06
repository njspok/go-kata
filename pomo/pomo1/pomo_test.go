package pomo1

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Run("construct", func(t *testing.T) {
		sut := New()
		interval, err := sut.Interval(time.Now())
		require.ErrorIs(t, err, ErrNotRunning)
		require.Nil(t, interval)
	})
	t.Run("repeating start", func(t *testing.T) {
		sut := New()

		start := time.Now()

		err := sut.Start(start)
		require.NoError(t, err)

		err = sut.Start(start)
		require.ErrorIs(t, err, ErrAlreadyRunning)
	})
	t.Run("work", func(t *testing.T) {
		sut := New()

		start := makeTime(t, "10:00:00")

		err := sut.Start(start)
		require.NoError(t, err)

		cases := []struct {
			now   string
			start string
			end   string
			tp    IntervalType
		}{
			{
				now:   "10:00:00",
				tp:    Work,
				start: "10:00:00",
				end:   "10:25:00",
			},
			{
				now:   "10:25:00",
				tp:    ShortRelax,
				start: "10:25:00",
				end:   "10:30:00",
			},
			{
				now:   "10:30:00",
				tp:    Work,
				start: "10:30:00",
				end:   "10:55:00",
			},
			{
				now:   "10:55:00",
				tp:    ShortRelax,
				start: "10:55:00",
				end:   "11:00:00",
			},
			{
				now:   "11:00:00",
				tp:    Work,
				start: "11:00:00",
				end:   "11:25:00",
			},
			{
				now:   "11:25:00",
				tp:    ShortRelax,
				start: "11:25:00",
				end:   "11:30:00",
			},
			{
				now:   "11:30:00",
				tp:    Work,
				start: "11:30:00",
				end:   "11:55:00",
			},
			{
				now:   "11:55:00",
				tp:    LongRelax,
				start: "11:55:00",
				end:   "12:10:00",
			},
			// new cycle
			{
				now:   "12:10:00",
				tp:    Work,
				start: "12:10:00",
				end:   "12:35:00",
			},
		}

		for n, c := range cases {
			t.Run(strconv.Itoa(n), func(t *testing.T) {
				interval, err := sut.Interval(makeTime(t, c.now))
				require.NoError(t, err)
				require.EqualValues(t, c.tp, interval.Type())
				require.EqualValues(t, makeTime(t, c.start), interval.Start())
				require.EqualValues(t, makeTime(t, c.end), interval.End())
			})
		}
	})
	t.Run("incorrect time", func(t *testing.T) {
		sut := New()

		start := time.Now()

		err := sut.Start(start)
		require.NoError(t, err)

		interval, err := sut.Interval(start.Add(-time.Minute))
		require.ErrorIs(t, err, ErrIncorrectTime)
		require.Nil(t, interval)
	})
	t.Run("reset", func(t *testing.T) {
		sut := New()

		err := sut.Start(makeTime(t, "10:00:00"))
		require.NoError(t, err)

		interval, err := sut.Interval(makeTime(t, "10:25:00"))
		require.NoError(t, err)
		require.EqualValues(t, ShortRelax, interval.Type())

		sut.Reset()
		err = sut.Start(makeTime(t, "10:25:00"))
		require.NoError(t, err)

		interval, err = sut.Interval(makeTime(t, "10:25:00"))
		require.NoError(t, err)
		require.EqualValues(t, Work, interval.Type())
	})
}

func makeTime(t *testing.T, str string) time.Time {
	t.Helper()
	tm, err := time.Parse(time.TimeOnly, str)
	require.NoError(t, err)
	return tm
}
