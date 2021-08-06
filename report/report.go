package report

import (
	"fmt"
	"sort"
	"strings"
)

func NewReport() *Report {
	return &Report{
		rows: make(map[string]*Row),
	}
}

type Report struct {
	rows map[string]*Row
}

func (r *Report) Sum() int {
	var sum int
	for _, row := range r.rows {
		sum += row.Price
	}
	return sum
}

func (r *Report) Append(row *Row) {
	r.rows[row.Name] = row
}

func (r *Report) Rows() map[string]*Row {
	return r.rows
}

func (r *Report) Count() int {
	return len(r.rows)
}

func (r *Report) PrintJson() string {
	return r.print(
		"%v: %v",
		"{sum: %v, items: {%v}}",
		", ",
	)
}

func (r *Report) PrintXml() string {
	return r.print(
		"<item><name>%v</name><price>%v</price></item>",
		"<report><sum>%v</sum><items>%v</items></report>",
		"",
	)
}

func (r *Report) print(item, report, sep string) string {
	var elements []string
	for name, row := range r.rows {
		elements = append(elements, fmt.Sprintf(item, name, row.Price))
	}

	sort.Strings(elements)

	items := strings.Join(elements, sep)

	return fmt.Sprintf(report, r.Sum(), items)
}

func NewRow(name string, price int) *Row {
	return &Row{
		Name:  name,
		Price: price,
	}
}

type Row struct {
	Name  string
	Price int
}
