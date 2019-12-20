package day19

import (
	"github.com/mjm/advent-of-code-2019/pkg/intcode"
	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// Drone represents a drone being used to map out the area affected by a
// tractor beam.
type Drone struct {
	canvas *Canvas
}

// NewDrone creates a new drone that records its findings onto a given canvas.
func NewDrone(c *Canvas) *Drone {
	return &Drone{
		canvas: c,
	}
}

// Scan checks a rectangular area for which points are affected by the tractor
// beam, and returns the number of affected points.
func (d *Drone) Scan(maxX, maxY int, program *intcode.VM) int {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			d.isInBeam(x, y, program)
		}
	}

	return d.canvas.CountColor(int('#'))
}

// FindSquare finds the closest corner of a size x size square that is entirely
// within the area affected by the tractor beam.
func (d *Drone) FindSquare(size int, program *intcode.VM) point.Point2D {
	var x, y int
	var startX int
	var seenOnRow bool
	for {
		inBeam := d.isInBeam(x, y, program)
		if inBeam {
			if !seenOnRow {
				seenOnRow = true
				startX = x
			}

			if x >= startX+99 && y >= 99 {
				if d.isInBeam(x-99, y-99, nil) && d.isInBeam(x, y-99, nil) {
					return point.New2D(x-99, y-99)
				}
			}

			x++
		} else if seenOnRow {
			y++
			x = startX
			seenOnRow = false
		} else {
			x++
		}
	}
}

func (d *Drone) isInBeam(x, y int, program *intcode.VM) bool {
	if result := d.canvas.At(point.New2D(x, y)); result != 0 {
		return result == '#'
	}

	if program == nil {
		return false
	}

	vm := program.Clone()
	vm.SetInputSeq([]int{x, y})
	go vm.MustExecute()

	status := <-vm.Output

	var result bool
	c := '.'
	if status == 1 {
		c = '#'
		result = true
	}
	d.canvas.Paint(point.Point2D{X: x, Y: y}, int(c))
	return result
}
