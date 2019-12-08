package day3

import (
	"math"
)

type point struct {
	X int
	Y int
}

func (p point) distance() int {
	return abs(p.X) + abs(p.Y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Map struct {
	grid      map[point]int
	highestId int8
}

func NewMap() *Map {
	return &Map{
		grid: make(map[point]int),
	}
}

func (m *Map) Count(x, y int) int {
	return m.countAtPoint(point{x, y})
}

func (m *Map) countAtPoint(p point) int {
	val := m.grid[p]
	var count int
	for i := m.highestId; i >= 0; i-- {
		check := 1 << i
		if val&check != 0 {
			count++
		}
	}
	return count
}

func (m *Map) Set(x, y int, id int8) {
	p := point{x, y}
	val := m.grid[p]
	val |= 1 << id
	m.grid[p] = val

	if id > m.highestId {
		m.highestId = id
	}
}

func (m *Map) NearestIntersection() (int, int, int) {
	var best point
	d := math.MaxInt64
	for p := range m.grid {
		n := m.countAtPoint(p)
		if n < 2 {
			continue
		}

		if best.X == 0 && best.Y == 0 {
			best = p
			d = p.distance()
		} else {
			nd := p.distance()
			if nd < d {
				best = p
				d = nd
			}
		}
	}

	return best.X, best.Y, d
}
