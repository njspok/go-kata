package report

func NewReport() *Report {
	return &Report{
		rows: make(Rows),
	}
}

type Report struct {
	rows Rows
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

func (r *Report) Rows() Rows {
	return r.rows
}

func (r *Report) Count() int {
	return len(r.rows)
}

func (r *Report) Print(v View) string {
	if v == nil {
		return ""
	}
	return v(r)
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

type View func(*Report) string

type Rows map[string]*Row
