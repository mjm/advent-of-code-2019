package point

// Point2D is a two-dimensional point in Cartesian coordinates.
type Point2D struct {
	X int
	Y int
}

func (p Point2D) Plus(x, y int) Point2D {
	p.X += x
	p.Y += y
	return p
}
