package report

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReport(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		rep := NewReport()
		require.Equal(t, 0, rep.Sum())
		require.Equal(t, 0, rep.Count())
		require.Empty(t, rep.Rows())
	})
	t.Run("append rows", func(t *testing.T) {
		rep := NewReport()

		rep.Append(NewRow("drive", 1111))
		require.Equal(t, 1, rep.Count())
		require.Equal(t, 1111, rep.Sum())
		require.Equal(t, map[string]*Row{
			"drive": &Row{Name: "drive", Price: 1111},
		}, rep.Rows())

		rep.Append(NewRow("wheels", 999))
		require.Equal(t, 2, rep.Count())
		require.Equal(t, 2110, rep.Sum())
		require.Equal(t, map[string]*Row{
			"drive":  &Row{Name: "drive", Price: 1111},
			"wheels": &Row{Name: "wheels", Price: 999},
		}, rep.Rows())

		rep.Append(NewRow("body", 10000))
		require.Equal(t, 3, rep.Count())
		require.Equal(t, 12110, rep.Sum())
		require.Equal(t, map[string]*Row{
			"drive":  &Row{Name: "drive", Price: 1111},
			"body":   &Row{Name: "body", Price: 10000},
			"wheels": &Row{Name: "wheels", Price: 999},
		}, rep.Rows())
	})
	t.Run("append and replace row", func(t *testing.T) {
		rep := NewReport()

		rep.Append(NewRow("drive", 1111))
		rep.Append(NewRow("wheels", 999))
		rep.Append(NewRow("wheels", 2222))

		require.Equal(t, 3333, rep.Sum())
		require.Equal(t, 2, rep.Count())
		require.Equal(t, map[string]*Row{
			"drive":  &Row{Name: "drive", Price: 1111},
			"wheels": &Row{Name: "wheels", Price: 2222},
		}, rep.Rows())
	})
	t.Run("print empty view", func(t *testing.T) {
		rep := NewReport()
		require.Empty(t, rep.Print(nil))
	})
	t.Run("print some view", func(t *testing.T) {
		rep := NewReport()
		res := rep.Print(func(report *Report) string {
			require.Equal(t, report, rep)
			return "some print"
		})
		require.Equal(t, "some print", res)
	})
}
