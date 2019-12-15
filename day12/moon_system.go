package day12

import (
	"strings"
)

// MoonSystem is a set of moons that all affect each other's motion.
type MoonSystem struct {
	Moons []*Moon
}

// NewSystemFromString reads a list of moons from lines in a string, and creates
// a new system with those moons.
func NewSystemFromString(s string) (*MoonSystem, error) {
	lines := strings.Split(s, "\n")
	moons := make([]*Moon, 0, len(lines))

	for _, line := range lines {
		moon, err := NewMoonFromString(line)
		if err != nil {
			return nil, err
		}
		moons = append(moons, moon)
	}

	return &MoonSystem{
		Moons: moons,
	}, nil
}

// Advance advances by the given number of time steps, applying gravity and simulating
// velocity for each step.
func (s *MoonSystem) Advance(n int) {
	for i := 0; i < n; i++ {
		s.advanceOnce()
	}
}

// Energy returns the total energy of all of the moons in the system.
func (s *MoonSystem) Energy() int {
	var total int
	for _, m := range s.Moons {
		total += m.Energy()
	}
	return total
}

func (s *MoonSystem) advanceOnce() {
	for _, mp := range s.pairedMoons() {
		mp.applyGravity()
	}

	for _, moon := range s.Moons {
		moon.Pos.Add(moon.Vel)
	}
}

func (s *MoonSystem) pairedMoons() []moonPair {
	var pairs []moonPair
	for i, a := range s.Moons {
		for j := i + 1; j < len(s.Moons); j++ {
			b := s.Moons[j]
			pairs = append(pairs, moonPair{a, b})
		}
	}
	return pairs
}

type moonPair struct {
	a *Moon
	b *Moon
}

func (mp moonPair) applyGravity() {
	applyGravityToAxis(mp.a.Pos.X, mp.b.Pos.X, &mp.a.Vel.X, &mp.b.Vel.X)
	applyGravityToAxis(mp.a.Pos.Y, mp.b.Pos.Y, &mp.a.Vel.Y, &mp.b.Vel.Y)
	applyGravityToAxis(mp.a.Pos.Z, mp.b.Pos.Z, &mp.a.Vel.Z, &mp.b.Vel.Z)
}

func applyGravityToAxis(pa, pb int, va, vb *int) {
	if pa < pb {
		*va++
		*vb--
	} else if pa > pb {
		*vb++
		*va--
	}
}
