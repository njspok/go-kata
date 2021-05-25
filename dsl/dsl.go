package dsl

var Root = &SolarNode{}
var Current interface{} = Root

type SolarNode struct {
	Name        string
	Description string
	Planets     []*PlanetNode
}

type PlanetNode struct {
	Name        string
	Description string
	Mass        uint
}

func Clean() {
	Root = &SolarNode{}
	Current = Root
}

func SolarSystem(name string, f func()) {
	Clean()

	if _, ok := Current.(*SolarNode); !ok {
		panic("Invalid call order")
	}

	Root.Name = name
	f()
}

func Description(d string) {
	switch e := Current.(type) {
	case *SolarNode:
		e.Description = d
	case *PlanetNode:
		e.Description = d
	default:
		panic("Node have not attribute Description")
	}
}

func Mass(m uint) {
	switch e := Current.(type) {
	case *PlanetNode:
		e.Mass = m
	default:
		panic("Node have not attribute Mass")
	}
}

func Name(n string) {
	Root.Name = n
}

func Planet(name string, f func()) {
	p := &PlanetNode{Name: name}
	Root.Planets = append(Root.Planets, p)
	Current = p
	f()
}
