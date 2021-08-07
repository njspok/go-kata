package report

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReport(t *testing.T) {
	t.Run("empty report", func(t *testing.T) {
		rep := NewReport()
		require.Equal(t, 0, rep.Sum())
		require.Equal(t, 0, rep.Count())
	})
	t.Run("filled report", func(t *testing.T) {
		rep := NewReport()

		var row *Row

		row = NewRow("drive", 1111)
		rep.Append(row)
		require.Equal(t, 1, rep.Count())

		row = NewRow("whells", 999)
		rep.Append(row)
		require.Equal(t, 2, rep.Count())

		row = NewRow("body", 10000)
		rep.Append(row)
		require.Equal(t, 3, rep.Count())

		require.Equal(t, 12110, rep.Sum())
	})
	t.Run("json", func(t *testing.T) {
		rep := NewReport()
		require.Equal(t, "{sum: 0, items: {}}", rep.Print(jsonView))

		rep.Append(NewRow("drive", 1111))
		require.Equal(t, "{sum: 1111, items: {drive: 1111}}", rep.Print(jsonView))

		rep.Append(NewRow("whells", 999))
		require.Equal(t, "{sum: 2110, items: {drive: 1111, whells: 999}}", rep.Print(jsonView))

		rep.Append(NewRow("body", 10000))
		require.Equal(t, "{sum: 12110, items: {body: 10000, drive: 1111, whells: 999}}", rep.Print(jsonView))
	})
	t.Run("xml", func(t *testing.T) {
		rep := NewReport()
		require.Equal(t, strings.Join([]string{
			"<report>",
			"<sum>0</sum>",
			"<items></items>",
			"</report>",
		}, ""), rep.Print(xmlView))

		rep.Append(NewRow("driver", 1111))
		require.Equal(t, strings.Join([]string{
			"<report>",
			"<sum>1111</sum>",
			"<items>",
			"<item><name>driver</name><price>1111</price></item>",
			"</items>",
			"</report>",
		}, ""), rep.Print(xmlView))

		rep.Append(NewRow("whells", 999))
		require.Equal(t, strings.Join([]string{
			"<report>",
			"<sum>2110</sum>",
			"<items>",
			"<item><name>driver</name><price>1111</price></item>",
			"<item><name>whells</name><price>999</price></item>",
			"</items>",
			"</report>",
		}, ""), rep.Print(xmlView))

		rep.Append(NewRow("body", 10000))
		require.Equal(t, strings.Join([]string{
			"<report>",
			"<sum>12110</sum>",
			"<items>",
			"<item><name>body</name><price>10000</price></item>",
			"<item><name>driver</name><price>1111</price></item>",
			"<item><name>whells</name><price>999</price></item>",
			"</items>",
			"</report>",
		}, ""), rep.Print(xmlView))
	})
}
