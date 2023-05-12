package celebrity

// https://www.youtube.com/watch?v=xGvQN_g-JCI

func NewPerson() *Person {
	return &Person{
		make(map[*Person]struct{}),
	}
}

type Person struct {
	knows map[*Person]struct{}
}

func (m *Person) Show(man *Person) {
	m.knows[man] = struct{}{}
}

func (m *Person) IsDontKnow(man *Person) bool {
	return !m.IsKnow(man)
}

func (m *Person) IsKnow(man *Person) bool {
	_, found := m.knows[man]
	return found
}
