package day10

import (
	"errors"
	"fmt"
	"strings"
)

type point struct {
	x int
	y int
}

// Map is a map of asteroids on a two-dimensional grid.
type Map struct {
	asteroids []*Asteroid
	grid      map[point]*Asteroid
}

var (
	errEmptyMap = errors.New("the map string is empty")
)

// LoadFromString reads a map from a string representation of it.
func LoadFromString(s string) (*Map, error) {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return nil, errEmptyMap
	}

	var asteroids []*Asteroid
	grid := make(map[point]*Asteroid)
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				a := newAsteroid(x, y)
				a.ConnectAll(asteroids)
				asteroids = append(asteroids, a)
				grid[point{x, y}] = a
			}
		}
	}

	return &Map{
		asteroids: asteroids,
		grid:      grid,
	}, nil
}

// At checks whether there is an asteroid at the given coordinates.
func (m *Map) At(x, y int) bool {
	_, ok := m.grid[point{x, y}]
	return ok
}

// VisibleAsteroids returns the number of asteroids visible from a station on the asteroid
// at the given coordinates.
func (m *Map) VisibleAsteroids(x, y int) int {
	a := m.grid[point{x, y}]
	if a == nil {
		panic(fmt.Errorf("cannot check visible asteroids from (%d, %d) because there is no asteroid there", x, y))
	}

	return len(a.VisibleAsteroids())
}

// BestAsteroid returns the coordinates of the asteroid with the best visibility to other
// asteroids, along with the number of asteroids visible from it.
func (m *Map) BestAsteroid() (int, int, int) {
	var best *Asteroid
	var count int
	for _, a := range m.asteroids {
		n := a.VisibleAsteroids()
		if len(n) > count {
			best, count = a, len(n)
		}
	}
	return best.X, best.Y, count
}

// Vaporized returns the coordinates of the Nth asteroid to be vaporized by the laser at
// the given coordinates.
func (m *Map) Vaporized(x, y, n int) (int, int) {
	a := m.grid[point{x, y}]

	var seen int
	for i := 0; true; i++ {
		as := a.AsteroidsAtDistance(i)
		if seen+len(as) <= n {
			seen += len(as)
		} else {
			b := as[n-seen]
			return b.X, b.Y
		}
	}

	panic("how did I get here?")
}
