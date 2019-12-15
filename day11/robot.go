package day11

import (
	"fmt"
	"sync"

	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// Robot is an emergency hull painting robot.
type Robot struct {
	canvas    *Canvas
	location  point.Point2D
	direction Direction
	lock      sync.Mutex
}

// NewRobot creates a new robot that paints on the given canvas.
func NewRobot(c *Canvas) *Robot {
	return &Robot{
		canvas:    c,
		direction: Up,
	}
}

// Run runs the robot according to the instructions from output.
func (r *Robot) Run(output chan int) {
	for {
		color, ok := <-output
		if !ok {
			return
		}
		r.lock.Lock()
		turn, ok := <-output
		if !ok {
			return
		}

		r.canvas.Paint(r.location, color)

		if turn == 0 {
			r.direction.TurnLeft()
		} else if turn == 1 {
			r.direction.TurnRight()
		} else {
			panic(fmt.Errorf("unexpected direction command: %d", turn))
		}

		r.direction.Advance(&r.location)
		r.lock.Unlock()
	}
}

// CurrentColor returns the color of the canvas at the robot's current location.
func (r *Robot) CurrentColor() int {
	r.lock.Lock()
	color := r.canvas.At(r.location)
	r.lock.Unlock()
	return color
}
