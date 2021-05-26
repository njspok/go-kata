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

	switch n := Current.(type) {
	case *SolarNode:
		n.Name = name
		process(n, f)
	default:
		NodeWithoutAttribute(n, "SolarSystem")
	}
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

func Name(name string) {
	switch n := Current.(type) {
	case *SolarNode:
		n.Name = name
	case *PlanetNode:
		n.Name = name
	case *SatelliteNode:
		n.Name = name
	default:
		NodeWithoutAttribute(n, "Name")
	}
}

func Planet(name string, f func()) {
	switch n := Current.(type) {
	case *SolarNode:
		p := &PlanetNode{Name: name}
		n.Planets = append(Root.Planets, p)
		process(p, f)
	default:
		NodeWithoutAttribute(n, "Planets")
	}
}

func Satellite(name string, f func()) {
	switch n := Current.(type) {
	case *PlanetNode:
		s := &SatelliteNode{Name: name}
		n.Satellites = append(n.Satellites, s)
		process(s, f)
	default:
		NodeWithoutAttribute(n, "Satellites")
	}
}

func NodeWithoutAttribute(node Node, attr string) {
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
