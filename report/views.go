package report

import (
	"fmt"
	"sort"
	"strings"
)

func defaultView(r *Report, item, report, sep string) string {
	if r == nil {
		return ""
	}

	var elements []string
	for name, row := range r.Rows() {
		elements = append(elements, fmt.Sprintf(item, name, row.Price))
	}

	sort.Strings(elements)

	items := strings.Join(elements, sep)

	return fmt.Sprintf(report, r.Sum(), items)
}

func XmlView(r *Report) string {
	return defaultView(
		r,
		"<item><name>%v</name><price>%v</price></item>",
		"<report><sum>%v</sum><items>%v</items></report>",
		"",
	)
}

func JsonView(r *Report) string {
	return defaultView(
		r,
		"%v: %v",
		"{sum: %v, items: {%v}}",
		", ",
	)
}

func TotalView(r *Report) string {
	if r == nil {
		return ""
	}
	return fmt.Sprintf("total: %v", r.Sum())
}
