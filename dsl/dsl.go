package dsl

func SolarSystem(name string, f func()) {
	clean()

	switch n := Current.(type) {
	case *SolarNode:
		n.Name = name
		next(n, f)
	default:
		nodeWithoutAttributePanic(Current, "SolarSystem")
	}
}

func Description(d string) {
	n, ok := Current.(interface{ SetDescription(string) })

	if !ok {
		nodeWithoutAttributePanic(Current, "Description")
	}

	n.SetDescription(d)
}

func Mass(m uint) {
	n, ok := Current.(interface{ SetMass(uint) })

	if !ok {
		nodeWithoutAttributePanic(Current, "Mass")
	}

	n.SetMass(m)
}

func Name(name string) {
	n, ok := Current.(interface{ SetName(string) })

	if !ok {
		nodeWithoutAttributePanic(Current, "Name")
	}

	n.SetName(name)
}

func Planet(name string, f func()) {
	n, ok := Current.(interface{ AddPlanet(*PlanetNode) })

	if !ok {
		nodeWithoutAttributePanic(Current, "Planets")
	}

	p := &PlanetNode{Name: name}
	n.AddPlanet(p)
	next(p, f)
}

func Satellite(name string, f func()) {
	n, ok := Current.(interface{ AddSatellite(*SatelliteNode) })

	if !ok {
		nodeWithoutAttributePanic(Current, "Satellites")
	}

	p := &SatelliteNode{Name: name}
	n.AddSatellite(p)
	next(p, f)
}
