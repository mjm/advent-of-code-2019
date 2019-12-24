package day24

import "strings"

// Grid is a rectangular grid of bugs.
type Grid struct {
	bugs   uint64
	width  int
	height int
}

// EmptyGrid creates an empty bug grid of a given size.
func EmptyGrid(width, height int) *Grid {
	return &Grid{
		width:  width,
		height: height,
	}
}

// GridFromString reads a string representation of a grid of bugs.
func GridFromString(s string) *Grid {
	lines := strings.Split(s, "\n")
	h := len(lines)
	w := len(lines[0])

	g := EmptyGrid(w, h)

	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				g.Infest(x, y)
			}
		}
	}

	return g
}

// SimulateOnce runs a simulation of the grid of bugs for one time step.
func (g *Grid) SimulateOnce() {
	ng := *g

	for y := 0; y < ng.height; y++ {
		for x := 0; x < ng.width; x++ {
			if g.shouldBecomeInfested(x, y) {
				ng.Infest(x, y)
			}
			if g.shouldDie(x, y) {
				ng.Kill(x, y)
			}
		}
	}

	*g = ng
}

// Simulate runs a simulation of the grid until a layout of bugs repeats.
// Returns the biodiversity rating of that layout.
func (g *Grid) Simulate() uint64 {
	seen := make(map[uint64]interface{})

	for i := 0; true; i++ {
		if _, ok := seen[g.bugs]; ok {
			return g.bugs
		}

		seen[g.bugs] = struct{}{}
		g.SimulateOnce()
	}

	panic("how did we get here?")
}

// Kill marks the tile at (x, y) as not having a bug.
func (g *Grid) Kill(x, y int) {
	g.bugs &^= g.mask(x, y)
}

// Infest marks the tile at (x, y) has having a bug.
func (g *Grid) Infest(x, y int) {
	g.bugs |= g.mask(x, y)
}

func (g *Grid) offset(x, y int) int {
	return y*g.width + x
}

func (g *Grid) mask(x, y int) uint64 {
	return 1 << g.offset(x, y)
}

func (g *Grid) hasBug(x, y int) bool {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return false
	}

	return g.bugs&g.mask(x, y) != 0
}

func (g *Grid) neighborCount(x, y int) int {
	var n int
	if g.hasBug(x-1, y) {
		n++
	}
	if g.hasBug(x+1, y) {
		n++
	}
	if g.hasBug(x, y-1) {
		n++
	}
	if g.hasBug(x, y+1) {
		n++
	}
	return n
}

func (g *Grid) shouldBecomeInfested(x, y int) bool {
	if g.hasBug(x, y) {
		return false
	}

	n := g.neighborCount(x, y)
	return n == 1 || n == 2
}

func (g *Grid) shouldDie(x, y int) bool {
	if !g.hasBug(x, y) {
		return false
	}

	return g.neighborCount(x, y) != 1
}
