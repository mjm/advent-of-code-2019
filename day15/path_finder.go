package day15

import (
	"math"

	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// PathFinder finds paths between a start and goal point on a map.
type PathFinder struct {
	canvas *Canvas
	start  point.Point2D
	goal   point.Point2D
	fs     map[point.Point2D]int
	gs     map[point.Point2D]int
}

// NewPathFinder creates a new path finder with a given map and goal.
func NewPathFinder(canvas *Canvas, goal point.Point2D) *PathFinder {
	var start point.Point2D

	gs := make(map[point.Point2D]int)
	gs[start] = 0
	fs := make(map[point.Point2D]int)

	pf := &PathFinder{
		canvas: canvas,
		start:  start,
		goal:   goal,
		fs:     fs,
		gs:     gs,
	}
	fs[start] = pf.h(start)
	return pf
}

// ShortestPath finds the shortest path from (0, 0) to the path finder's
// goal. Returns the list of points that form the path, starting from the
// goal and ending with the start point, including both.
func (pf *PathFinder) ShortestPath() []point.Point2D {
	open := make(map[point.Point2D]interface{})
	open[pf.start] = struct{}{}
	prev := make(map[point.Point2D]point.Point2D)

	for len(open) > 0 {
		var current *point.Point2D
		for p := range open {
			if current == nil || pf.f(p) < pf.f(*current) {
				current = &p
			}
		}

		if *current == pf.goal {
			n := *current
			path := []point.Point2D{n}
			for {
				var ok bool
				n, ok = prev[n]
				if !ok {
					return path
				}

				path = append(path, n)
			}
		}

		delete(open, *current)
		for _, n := range pf.neighbors(*current) {
			if g := pf.g(*current) + 1; g < pf.g(n) {
				prev[n] = *current
				pf.gs[n] = g
				pf.fs[n] = g + pf.h(n)
				open[n] = struct{}{}
			}
		}
	}

	return nil
}

func (pf *PathFinder) neighbors(p point.Point2D) []point.Point2D {
	var ps []point.Point2D
	for _, dir := range []Direction{North, South, West, East} {
		if p := dir.offset(p); pf.canvas.At(p) != int(TileWall) {
			ps = append(ps, p)
		}
	}
	return ps
}

func (pf *PathFinder) f(p point.Point2D) int {
	if f, ok := pf.fs[p]; ok {
		return f
	}
	return math.MaxInt64
}

func (pf *PathFinder) g(p point.Point2D) int {
	if g, ok := pf.gs[p]; ok {
		return g
	}
	return math.MaxInt64
}

func (pf *PathFinder) h(p point.Point2D) int {
	return abs(p.X-pf.goal.X) + abs(p.Y-pf.goal.Y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
