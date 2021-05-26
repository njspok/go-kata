package dsl

import (
	"fmt"
	"reflect"
)

var (
	Root         = &SolarNode{}
	Current Node = Root
)

type Node interface{}

type SolarNode struct {
	Name        string
	Description string
	Planets     []*PlanetNode
}

func (sol *SolarNode) SetName(name string) {
	sol.Name = name
}

func (sol *SolarNode) SetDescription(desc string) {
	sol.Description = desc
}

func (sol *SolarNode) AddPlanet(planet *PlanetNode) {
	sol.Planets = append(sol.Planets, planet)
}

type PlanetNode struct {
	Name        string
	Description string
	Mass        uint
	Satellites  []*SatelliteNode
}

func (pl *PlanetNode) SetName(name string) {
	pl.Name = name
}

func (pl *PlanetNode) SetMass(mass uint) {
	pl.Mass = mass
}

func (pl *PlanetNode) SetDescription(desc string) {
	pl.Description = desc
}

func (pl *PlanetNode) AddSatellite(sat *SatelliteNode) {
	pl.Satellites = append(pl.Satellites, sat)
}

type SatelliteNode struct {
	Name        string
	Description string
	Mass        uint
}

func (sat *SatelliteNode) SetName(name string) {
	sat.Name = name
}

func (sat *SatelliteNode) SetMass(mass uint) {
	sat.Mass = mass
}

func (sat *SatelliteNode) SetDescription(desc string) {
	sat.Description = desc
}

func Clean() {
	Root = &SolarNode{}
	Current = Root
}

func SolarSystem(name string, f func()) {
	Clean()

	switch n := Current.(type) {
	case *SolarNode:
		n.Name = name
		process(n, f)
	default:
		nodeWithoutAttributePanic(n, "SolarSystem")
	}
}

func Description(d string) {
	n, ok := Current.(interface{ SetDescription(string) })

	if !ok {
		nodeWithoutAttributePanic(n, "Description")
	}

	n.SetDescription(d)
}

func Mass(m uint) {
	n, ok := Current.(interface{ SetMass(uint) })

	if !ok {
		nodeWithoutAttributePanic(n, "Mass")
	}

	n.SetMass(m)
}

func Name(name string) {
	n, ok := Current.(interface{ SetName(string) })

	if !ok {
		nodeWithoutAttributePanic(n, "Name")
	}

	n.SetName(name)
}

func Planet(name string, f func()) {
	n, ok := Current.(interface{ AddPlanet(*PlanetNode) })

	if !ok {
		nodeWithoutAttributePanic(n, "Planets")
	}

	p := &PlanetNode{Name: name}
	n.AddPlanet(p)
	process(p, f)
}

func Satellite(name string, f func()) {
	n, ok := Current.(interface{ AddSatellite(*SatelliteNode) })

	if !ok {
		nodeWithoutAttributePanic(n, "Satellites")
	}

	p := &SatelliteNode{Name: name}
	n.AddSatellite(p)
	process(p, f)
}

func nodeWithoutAttributePanic(node Node, attr string) {
	panic(fmt.Sprintf(
		"Node %s have not attribute %s",
		reflect.TypeOf(node).String(),
		attr,
	))
}

func process(node Node, f func()) {
	prev := Current
	Current = node
	f()
	Current = prev
}
