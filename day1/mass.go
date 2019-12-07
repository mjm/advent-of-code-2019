package day1

type Mass interface {
	Mass() int
}

type Module struct {
	mass int
}

func NewModule(mass int) *Module {
	return &Module{mass: mass}
}

func (m *Module) Mass() int {
	return m.mass
}

type fuel struct {
	mass int
}

func (f *fuel) Mass() int {
	return f.mass
}
