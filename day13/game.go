package day13

import (
	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// Game plays a game controlled by an Intcode program
type Game struct {
	canvas *Canvas
}

// NewGame creates a new game that draws on the given canvas.
func NewGame(c *Canvas) *Game {
	return &Game{
		canvas: c,
	}
}

// Run runs the game according to the instructions from output.
func (r *Game) Run(output chan int) {
	for {
		x, ok := <-output
		if !ok {
			return
		}
		y, ok := <-output
		if !ok {
			return
		}
		tile, ok := <-output
		if !ok {
			return
		}

		r.canvas.Paint(point.Point2D{X: x, Y: y}, tile)
	}
}
