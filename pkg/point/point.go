package point

// Point2D is a two-dimensional point in Cartesian coordinates.
type Point2D struct {
	X int
	Y int
}

func New2D(x, y int) Point2D {
	return Point2D{X: x, Y: y}
}

func (p Point2D) Plus(x, y int) Point2D {
	p.X += x
	p.Y += y
	return p
}

func (p Point2D) CardinalNeighbors() []Point2D {
	return []Point2D{
		p.Plus(-1, 0),
		p.Plus(1, 0),
		p.Plus(0, -1),
		p.Plus(0, 1),
	}
}
