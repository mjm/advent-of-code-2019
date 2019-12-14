package day10

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

// Map is a map of asteroids on a two-dimensional grid.
type Map struct {
	grid   [][]bool
	width  int
	height int
}

var (
	errEmptyMap = errors.New("the map string is empty")
)

// LoadFromString reads a map from a string representation of it.
func LoadFromString(s string) (*Map, error) {
	lines := strings.Split(s, "\n")
	height := len(lines)
	if height == 0 {
		return nil, errEmptyMap
	}
	width := len(lines[0])

	grid := make([][]bool, 0, height)
	for y, line := range lines {
		if len(line) != width {
			return nil, fmt.Errorf("line %d should have %d characters, but has %d instead", y, width, len(line))
		}

		row := make([]bool, 0, width)
		for x, c := range line {
			switch c {
			case '#':
				row = append(row, true)
			case '.':
				row = append(row, false)
			default:
				return nil, fmt.Errorf("unexpected character at row %d, column %d: %q", y, x, c)
			}
		}

		grid = append(grid, row)
	}

	return &Map{
		grid:   grid,
		width:  width,
		height: height,
	}, nil
}

// At checks whether there is an asteroid at the given coordinates.
func (m *Map) At(x, y int) bool {
	return m.grid[y][x]
}

// VisibleAsteroids returns the number of asteroids visible from a station on the asteroid
// at the given coordinates.
func (m *Map) VisibleAsteroids(x, y int) int {
	if !m.At(x, y) {
		panic(fmt.Errorf("cannot check visible asteroids from (%d, %d) because there is no asteroid there", x, y))
	}

	var count int
	for y1 := 0; y1 < m.height; y1++ {
		for x1 := 0; x1 < m.width; x1++ {
			if x == x1 && y == y1 {
				continue
			}

			if !m.At(x1, y1) {
				continue
			}

			if m.isVisibleFrom(x, y, x1, y1) {
				count++
			}
		}
	}

	return count
}

// BestAsteroid returns the coordinates of the asteroid with the best visibility to other
// asteroids, along with the number of asteroids visible from it.
func (m *Map) BestAsteroid() (int, int, int) {
	var xb, yb, asteroids int
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if !m.At(x, y) {
				continue
			}

			n := m.VisibleAsteroids(x, y)
			if n > asteroids {
				xb, yb, asteroids = x, y, n
			}
		}
	}

	return xb, yb, asteroids
}

func (m *Map) isVisibleFrom(x, y, x1, y1 int) bool {
	dx := int64(x1 - x)
	dy := int64(y1 - y)

	if dx == 0 {
		offset := -1
		if dy > 0 {
			offset = 1
		}
		return m.isVisibleFromDirection(x, y, x1, y1, 0, offset)
	}

	if dy == 0 {
		offset := -1
		if dx > 0 {
			offset = 1
		}
		return m.isVisibleFromDirection(x, y, x1, y1, offset, 0)
	}

	offset := big.NewRat(dx, dy)

	// check every interval between for a blocking asteroid
	var dx1, dy1 int64
	for {
		if dy < 0 {
			dx1 -= offset.Num().Int64()
			dy1 -= offset.Denom().Int64()
		} else {
			dx1 += offset.Num().Int64()
			dy1 += offset.Denom().Int64()
		}

		// if we've gotten to the original asteroid location, then we didn't
		// find a blocker and the asteroid is visible
		if dx1 == dx && dy1 == dy {
			return true
		}

		// if there's an asteroid here, then it's blocking the view of the one
		// we're checking.
		if m.At(int(dx1)+x, int(dy1)+y) {
			return false
		}
	}
}

func (m *Map) isVisibleFromDirection(x, y, x1, y1, dx, dy int) bool {
	for x2, y2 := x+dx, y+dy; x2 != x1 || y2 != y1; x2, y2 = x2+dx, y2+dy {
		if m.At(x2, y2) {
			return false
		}
	}
	return true
}
