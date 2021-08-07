package report

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTotalView(t *testing.T) {
	require.Empty(t, TotalView(nil))

	rep := NewReport()
	require.Equal(t, "total: 0", TotalView(rep))

	rep.Append(NewRow("drive", 1111))
	require.Equal(t, "total: 1111", TotalView(rep))

	rep.Append(NewRow("wheels", 999))
	require.Equal(t, "total: 2110", TotalView(rep))

	rep.Append(NewRow("body", 10000))
	require.Equal(t, "total: 12110", TotalView(rep))
}

func TestJsonView(t *testing.T) {
	require.Empty(t, JsonView(nil))

	rep := NewReport()
	require.Equal(t, "{sum: 0, items: {}}", JsonView(rep))

	rep.Append(NewRow("drive", 1111))
	require.Equal(t, "{sum: 1111, items: {drive: 1111}}", JsonView(rep))

	rep.Append(NewRow("wheels", 999))
	require.Equal(t, "{sum: 2110, items: {drive: 1111, wheels: 999}}", JsonView(rep))

	rep.Append(NewRow("body", 10000))
	require.Equal(t, "{sum: 12110, items: {body: 10000, drive: 1111, wheels: 999}}", JsonView(rep))
}

func TestXmlView(t *testing.T) {
	require.Empty(t, XmlView(nil))

	rep := NewReport()
	require.Equal(t, strings.Join([]string{
		"<report>",
		"<sum>0</sum>",
		"<items></items>",
		"</report>",
	}, ""), XmlView(rep))

	rep.Append(NewRow("driver", 1111))
	require.Equal(t, strings.Join([]string{
		"<report>",
		"<sum>1111</sum>",
		"<items>",
		"<item><name>driver</name><price>1111</price></item>",
		"</items>",
		"</report>",
	}, ""), XmlView(rep))

	rep.Append(NewRow("wheels", 999))
	require.Equal(t, strings.Join([]string{
		"<report>",
		"<sum>2110</sum>",
		"<items>",
		"<item><name>driver</name><price>1111</price></item>",
		"<item><name>wheels</name><price>999</price></item>",
		"</items>",
		"</report>",
	}, ""), XmlView(rep))

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
	}, ""), XmlView(rep))

}
