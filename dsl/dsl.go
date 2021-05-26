package dsl

func SolarSystem(name string, f func()) {
	clean()

	switch n := Current.(type) {
	case *SolarNode:
		n.Name = name
		next(n, f)
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
	next(p, f)
}

func Satellite(name string, f func()) {
	n, ok := Current.(interface{ AddSatellite(*SatelliteNode) })

	if !ok {
		nodeWithoutAttributePanic(n, "Satellites")
	}

	p := &SatelliteNode{Name: name}
	n.AddSatellite(p)
	next(p, f)
}
