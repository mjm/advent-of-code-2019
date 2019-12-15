package day11

import "github.com/mjm/advent-of-code-2019/pkg/point"

// Direction is one of the cardinal directions that the robot can move.
type Direction int

// The four directions the robot can move.
const (
	Up Direction = iota
	Down
	Left
	Right
)

// TurnLeft alters the direction to face left from the current direction.
func (d *Direction) TurnLeft() {
	switch *d {
	case Up:
		*d = Left
	case Down:
		*d = Right
	case Left:
		*d = Down
	case Right:
		*d = Up
	}
}

// TurnRight alters the direction to face right from the current direction.
func (d *Direction) TurnRight() {
	switch *d {
	case Up:
		*d = Right
	case Down:
		*d = Left
	case Left:
		*d = Up
	case Right:
		*d = Down
	}
}

// Advance moves a point in this direction.
func (d Direction) Advance(p *point.Point2D) {
	switch d {
	case Up:
		p.Y--
	case Down:
		p.Y++
	case Left:
		p.X--
	case Right:
		p.X++
	}
}
