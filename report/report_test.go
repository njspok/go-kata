package report

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReport(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		rep := NewReport()
		require.Equal(t, 0, rep.Sum())
		require.Equal(t, 0, rep.Count())
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
	t.Run("json view", func(t *testing.T) {
		rep := NewReport()
		require.Equal(t, "{sum: 0, items: {}}", rep.Print(JsonView))

		rep.Append(NewRow("drive", 1111))
		require.Equal(t, "{sum: 1111, items: {drive: 1111}}", rep.Print(JsonView))

		rep.Append(NewRow("wheels", 999))
		require.Equal(t, "{sum: 2110, items: {drive: 1111, wheels: 999}}", rep.Print(JsonView))

		rep.Append(NewRow("body", 10000))
		require.Equal(t, "{sum: 12110, items: {body: 10000, drive: 1111, wheels: 999}}", rep.Print(JsonView))
	})
	t.Run("xml view", func(t *testing.T) {
		rep := NewReport()
		require.Equal(t, strings.Join([]string{
			"<report>",
			"<sum>0</sum>",
			"<items></items>",
			"</report>",
		}, ""), rep.Print(XmlView))

		rep.Append(NewRow("driver", 1111))
		require.Equal(t, strings.Join([]string{
			"<report>",
			"<sum>1111</sum>",
			"<items>",
			"<item><name>driver</name><price>1111</price></item>",
			"</items>",
			"</report>",
		}, ""), rep.Print(XmlView))

		rep.Append(NewRow("wheels", 999))
		require.Equal(t, strings.Join([]string{
			"<report>",
			"<sum>2110</sum>",
			"<items>",
			"<item><name>driver</name><price>1111</price></item>",
			"<item><name>wheels</name><price>999</price></item>",
			"</items>",
			"</report>",
		}, ""), rep.Print(XmlView))

		rep.Append(NewRow("body", 10000))
		require.Equal(t, strings.Join([]string{
			"<report>",
			"<sum>12110</sum>",
			"<items>",
			"<item><name>body</name><price>10000</price></item>",
			"<item><name>driver</name><price>1111</price></item>",
			"<item><name>wheels</name><price>999</price></item>",
			"</items>",
			"</report>",
		}, ""), rep.Print(XmlView))
	})
	t.Run("total sum view", func(t *testing.T) {
		rep := NewReport()
		require.Equal(t, "total: 0", rep.Print(TotalView))

		rep.Append(NewRow("drive", 1111))
		require.Equal(t, "total: 1111", rep.Print(TotalView))

		rep.Append(NewRow("wheels", 999))
		require.Equal(t, "total: 2110", rep.Print(TotalView))

		rep.Append(NewRow("body", 10000))
		require.Equal(t, "total: 12110", rep.Print(TotalView))
	})
	t.Run("empty view", func(t *testing.T) {
		rep := NewReport()
		require.Empty(t, rep.Print(nil))
	})
}
