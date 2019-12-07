package day1

type Module struct {
	Mass int
}

func (m *Module) FuelRequired() int {
	return (m.Mass / 3) - 2
}
