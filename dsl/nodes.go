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

func clean() {
	Root = &SolarNode{}
	Current = Root
}
