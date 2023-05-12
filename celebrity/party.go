package celebrity

func NewParty() *Party {
	return &Party{}
}

type Party struct {
	persons []*Person
}

func (p *Party) Add(person *Person) {
	p.persons = append(p.persons, person)
}

func (p *Party) Celebrity() *Person {
	if len(p.persons) == 0 {
		return nil
	}

	if len(p.persons) == 1 {
		return nil
	}

	// find potential celebrity

	left := 0
	right := len(p.persons) - 1

	for left != right {
		if p.persons[left].IsKnow(p.persons[right]) {
			left++
		} else {
			right--
		}
	}

	// check potential celebrity

	found := p.persons[left]

	for _, cur := range p.persons {
		if cur == found {
			continue
		}

		if cur.IsKnow(found) && found.IsDontKnow(cur) {
			continue
		}

		return nil
	}

	return found
}
