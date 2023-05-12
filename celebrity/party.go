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
	if p.isNoPersons() {
		return nil
	}

	if p.isSingleParty() {
		return nil
	}

	found := p.findPotentialCelebrity()

	if p.isCelebrity(found) {
		return found
	}

	return nil
}

func (p *Party) isCelebrity(found *Person) bool {
	for _, cur := range p.persons {
		if cur == found {
			continue
		}

		if cur.IsKnow(found) && found.IsDontKnow(cur) {
			continue
		}

		return false
	}
	return true
}

func (p *Party) findPotentialCelebrity() *Person {
	left := 0
	right := len(p.persons) - 1

	for left != right {
		if p.persons[left].IsKnow(p.persons[right]) {
			left++
		} else {
			right--
		}
	}

	return p.persons[left]
}

func (p *Party) isSingleParty() bool {
	return len(p.persons) == 1
}

func (p *Party) isNoPersons() bool {
	return len(p.persons) == 0
}
