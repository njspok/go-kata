package bowling

func NewScorer() *Scorer {
	return &Scorer{
		total: 0,
	}
}

type Scorer struct {
	total uint
}

func (s *Scorer) Roll(count uint) {
	s.total += count
}

func (s *Scorer) Earn(frame Framer) {
	frame.SetScore(s.total)
}

func (s *Scorer) Total() uint {
	return s.total
}
