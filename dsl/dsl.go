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

type PlanetNode struct {
	Name        string
	Description string
	Mass        uint
	Satellites  []*SatelliteNode
}

type SatelliteNode struct {
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
	switch n := Current.(type) {
	case *SolarNode:
		n.Description = d
	case *PlanetNode:
		n.Description = d
	case *SatelliteNode:
		n.Description = d
	default:
		NodeWithoutAttribute(n, "Description")
	}
}

func Mass(m uint) {
	switch n := Current.(type) {
	case *PlanetNode:
		n.Mass = m
	case *SatelliteNode:
		n.Mass = m
	default:
		NodeWithoutAttribute(n, "Mass")
	}
}

func Name(n string) {
	Root.Name = n
}

func Planet(name string, f func()) {
	p := &PlanetNode{Name: name}
	Root.Planets = append(Root.Planets, p)

	prev := Current
	Current = p

	f()

	Current = prev
}

func Satellite(name string, f func()) {
	switch n := Current.(type) {
	case *PlanetNode:
		s := &SatelliteNode{Name: name}
		n.Satellites = append(n.Satellites, s)

		prev := Current
		Current = s

		f()

		Current = prev
	default:
		NodeWithoutAttribute(n, "Satellite")
	}
}

func NodeWithoutAttribute(n interface{}, attr string) {
	panic(fmt.Sprintf(
		"Node %s have not attribute %s",
		reflect.TypeOf(n).String(),
		attr,
	))
}
