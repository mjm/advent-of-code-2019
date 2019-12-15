package day10

import (
	"math"
)

// Asteroid is a single asteroid on the map. An Asteroid knows its location in Cartesian coordinates,
// and it also knows the location of all other asteroids in polar coordinates, relative to itself.
type Asteroid struct {
	X      int
	Y      int
	angles map[int][]*AsteroidEdge
}

func newAsteroid(x, y int) *Asteroid {
	return &Asteroid{
		X:      x,
		Y:      y,
		angles: make(map[int][]*AsteroidEdge),
	}
}

// ConnectAll connects this asteroid to all asteroids in the provided list.
func (a *Asteroid) ConnectAll(as []*Asteroid) {
	for _, b := range as {
		a.Connect(b)
	}
}

// Connect ensures that the two asteroids know their relative positions to each other.
func (a *Asteroid) Connect(b *Asteroid) {
	dx, dy := float64(b.X-a.X), float64(b.Y-a.Y)
	angle := math.Atan2(dy, dx)
	radius := math.Sqrt(dx*dx + dy*dy)

	a.addEdge(b, int(angle*10000), radius)

	otherAngle := angle + math.Pi
	if otherAngle > math.Pi*2 {
		otherAngle -= math.Pi * 2
	}
	b.addEdge(a, int(otherAngle*10000), radius)
}

func (a *Asteroid) addEdge(b *Asteroid, angle int, radius float64) {
	edge := &AsteroidEdge{
		Asteroid: b,
		Radius:   radius,
	}

	if edges, ok := a.angles[angle]; ok {
		for i, e := range edges {
			if radius < e.Radius {
				// insert into the edges slice
				edges = append(edges, nil)
				copy(edges[i+1:], edges[i:])
				edges[i] = edge
				a.angles[angle] = edges
				return
			}
		}
		a.angles[angle] = append(a.angles[angle], edge)
		return
	}

	a.angles[angle] = []*AsteroidEdge{edge}
}

// VisibleAsteroids returns the list of asteroids that are visible from this one.
// This will exclude asteroids that are occluded by a closer one.
func (a *Asteroid) VisibleAsteroids() []*Asteroid {
	var asteroids []*Asteroid
	for _, edges := range a.angles {
		asteroids = append(asteroids, edges[0].Asteroid)
	}
	return asteroids
}

// AsteroidEdge stores an asteroid and its distance from another astroid.
type AsteroidEdge struct {
	Asteroid *Asteroid
	Radius   float64
}
