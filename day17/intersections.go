package day17

import (
	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// IntersectionFinder builds a map of scaffolds and finds their intersections.
type IntersectionFinder struct {
	canvas *Canvas
}

// NewIntersectionFinder creates a new intersection finder using a canvas as
// a store for the map.
func NewIntersectionFinder(c *Canvas) *IntersectionFinder {
	return &IntersectionFinder{canvas: c}
}

// BuildMap reads ASCII characters from the output and paints the map onto the
// canvas.
func (finder *IntersectionFinder) BuildMap(output <-chan int) {
	p := point.Point2D{}
	for val := range output {
		if val == 10 {
			p.Y++
			p.X = 0
		} else {
			finder.canvas.Paint(p, val)
			p.X++
		}
	}
}

// AlignmentParameters returns the sum of the alignment parameters of all of
// the intersections of scaffolds.
func (finder *IntersectionFinder) AlignmentParameters() int {
	var n int
	for _, p := range finder.allIntersections() {
		n += p.X * p.Y
	}
	return n
}

const scaffold int = int('#')

func (finder *IntersectionFinder) allIntersections() []point.Point2D {
	c := finder.canvas
	var ps []point.Point2D
	for p := range c.paint {
		if finder.isScaffold(p) &&
			finder.isScaffold(p.Plus(-1, 0)) &&
			finder.isScaffold(p.Plus(1, 0)) &&
			finder.isScaffold(p.Plus(0, -1)) &&
			finder.isScaffold(p.Plus(0, 1)) {
			ps = append(ps, p)
		}
	}
	return ps
}

func (finder *IntersectionFinder) isScaffold(p point.Point2D) bool {
	val := finder.canvas.At(p)
	return val == scaffold || val == int('^')
}
