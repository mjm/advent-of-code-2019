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

type cell struct {
	id    int8
	steps int
}

type grid map[point]map[int8]cell

type Map struct {
	grid grid
}

func NewMap() *Map {
	return &Map{
		grid: make(grid),
	}
}

func (m *Map) Count(x, y int) int {
	return m.countAtPoint(point{x, y})
}

func (m *Map) countAtPoint(p point) int {
	return len(m.grid[p])
}

func (m *Map) Set(x, y int, id int8, steps int) {
	p := point{x, y}
	cells := m.grid[p]
	if cells == nil {
		cells = make(map[int8]cell)
		m.grid[p] = cells
	}

	cell, ok := cells[id]
	if !ok {
		cell.id = id
		cell.steps = steps
		cells[id] = cell
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

func (m *Map) ShortestIntersection() (int, int, int) {
	var best point
	steps := math.MaxInt64
	for p, cells := range m.grid {
		if len(cells) < 2 {
			continue
		}

		var totalSteps int
		for _, cell := range cells {
			totalSteps += cell.steps
		}

		if best.X == 0 && best.Y == 0 {
			best = p
			steps = totalSteps
		} else {
			if totalSteps < steps {
				best = p
				steps = totalSteps
			}
		}
	}

	return best.X, best.Y, steps
}
