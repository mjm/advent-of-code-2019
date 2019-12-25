package day24

import (
	"fmt"
	"math/bits"
	"strings"

	log "github.com/sirupsen/logrus"
)

// RecursiveGrid is a rectangular grid of bugs with multiple recursive levels.
type RecursiveGrid struct {
	bugs     map[int]uint64
	width    int
	height   int
	minLevel int
	maxLevel int
}

// EmptyRecursiveGrid creates an empty bug grid of a given size.
func EmptyRecursiveGrid(width, height int) *RecursiveGrid {
	return &RecursiveGrid{
		bugs:   make(map[int]uint64),
		width:  width,
		height: height,
	}
}

// RecursiveGridFromString reads a string representation of a single level
// of a grid of bugs. It creates a recursive grid with this level as level 0.
func RecursiveGridFromString(s string) *RecursiveGrid {
	lines := strings.Split(s, "\n")
	h := len(lines)
	w := len(lines[0])

	g := EmptyRecursiveGrid(w, h)

	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				g.Infest(0, x, y)
			}
		}
	}

	return g
}

// Clone creates a new grid with the same contents as this one.
func (g *RecursiveGrid) Clone() *RecursiveGrid {
	ng := *g

	ng.bugs = make(map[int]uint64)
	for level, bugs := range g.bugs {
		ng.bugs[level] = bugs
	}

	return &ng
}

func (g *RecursiveGrid) String() string {
	var b strings.Builder

	for l := g.minLevel; l <= g.maxLevel; l++ {
		fmt.Fprintf(&b, "Level %d\n", l)
		for y := 0; y < g.height; y++ {
			for x := 0; x < g.width; x++ {
				if g.hasBug(l, x, y) {
					fmt.Fprintf(&b, "#")
				} else {
					fmt.Fprintf(&b, ".")
				}
			}
			b.WriteRune('\n')
		}
		b.WriteRune('\n')
	}

	return b.String()
}

// SimulateOnce runs a simulation of the grid of bugs for one time step.
func (g *RecursiveGrid) SimulateOnce() {
	ng := g.Clone()

	var infested, killed int
	for level := g.minLevel - 1; level <= g.maxLevel+1; level++ {
		for y := 0; y < ng.height; y++ {
			for x := 0; x < ng.width; x++ {
				if x == 2 && y == 2 {
					continue
				}

				if g.shouldBecomeInfested(level, x, y) {
					ng.Infest(level, x, y)
					infested++
				}
				if g.shouldDie(level, x, y) {
					ng.Kill(level, x, y)
					killed++
				}
			}
		}
	}

	log.WithFields(log.Fields{
		"infested":      infested,
		"killed":        killed,
		"old.min_level": g.minLevel,
		"old.max_level": g.maxLevel,
		"new.min_level": ng.minLevel,
		"new.max_level": ng.maxLevel,
	}).Debug("ran single simulation")
	log.Debugf("simulation result:\n%s", ng)

	*g = *ng
}

// Simulate runs a simulation of the grid for n minutes. Returns the number of
// bugs on the map.
func (g *RecursiveGrid) Simulate(n int) int {
	for i := 0; i < n; i++ {
		g.SimulateOnce()
	}

	var bugs int
	for level := g.minLevel; level <= g.maxLevel; level++ {
		bugs += bits.OnesCount64(g.bugs[level])
	}
	return bugs
}

// Kill marks the tile at (x, y) as not having a bug.
func (g *RecursiveGrid) Kill(level, x, y int) {
	g.bugs[level] &^= g.mask(x, y)
}

// Infest marks the tile at (x, y) has having a bug.
func (g *RecursiveGrid) Infest(level, x, y int) {
	if level < g.minLevel {
		g.minLevel = level
	}
	if level > g.maxLevel {
		g.maxLevel = level
	}
	g.bugs[level] |= g.mask(x, y)
}

func (g *RecursiveGrid) offset(x, y int) int {
	return y*g.width + x
}

func (g *RecursiveGrid) mask(x, y int) uint64 {
	return 1 << g.offset(x, y)
}

func (g *RecursiveGrid) hasBug(level, x, y int) bool {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return false
	}

	return g.bugs[level]&g.mask(x, y) != 0
}

func (g *RecursiveGrid) neighborCount(level, x, y int) int {
	var n int
	if g.hasBug(level, x-1, y) {
		n++
	}
	if g.hasBug(level, x+1, y) {
		n++
	}
	if g.hasBug(level, x, y-1) {
		n++
	}
	if g.hasBug(level, x, y+1) {
		n++
	}

	// if at an edge of this level, check neighbors at the level containing this one
	if x == 0 && g.hasBug(level-1, 1, 2) {
		n++
	}
	if y == 0 && g.hasBug(level-1, 2, 1) {
		n++
	}
	if x == g.width-1 && g.hasBug(level-1, 3, 2) {
		n++
	}
	if y == g.height-1 && g.hasBug(level-1, 2, 3) {
		n++
	}

	// if at the middle of the level, add the neighbors from the contained level
	if x == 1 && y == 2 {
		for y1 := 0; y1 < g.height; y1++ {
			if g.hasBug(level+1, 0, y1) {
				n++
			}
		}
	}
	if x == 2 && y == 1 {
		for x1 := 0; x1 < g.width; x1++ {
			if g.hasBug(level+1, x1, 0) {
				n++
			}
		}
	}
	if x == 3 && y == 2 {
		for y1 := 0; y1 < g.height; y1++ {
			if g.hasBug(level+1, g.width-1, y1) {
				n++
			}
		}
	}
	if x == 2 && y == 3 {
		for x1 := 0; x1 < g.width; x1++ {
			if g.hasBug(level+1, x1, g.height-1) {
				n++
			}
		}
	}

	return n
}

func (g *RecursiveGrid) shouldBecomeInfested(level, x, y int) bool {
	if g.hasBug(level, x, y) {
		return false
	}

	n := g.neighborCount(level, x, y)
	return n == 1 || n == 2
}

func (g *RecursiveGrid) shouldDie(level, x, y int) bool {
	if !g.hasBug(level, x, y) {
		return false
	}

	return g.neighborCount(level, x, y) != 1
}
