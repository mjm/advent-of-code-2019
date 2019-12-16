package day15

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// Canvas is an infinitely paintable 2D surface.
type Canvas struct {
	paint     map[point.Point2D]int
	minCorner point.Point2D
	maxCorner point.Point2D
}

// NewCanvas creates an empty canvas.
func NewCanvas() *Canvas {
	return &Canvas{
		paint: make(map[point.Point2D]int),
	}
}

// Paint paints a point with a color, extending the bounds of the canvas if needed.
func (c *Canvas) Paint(p point.Point2D, color int) {
	c.paint[p] = color
	c.adjustSizeIfNeeded(p)
}

// At gets the color painted at a point.
func (c *Canvas) At(p point.Point2D) int {
	return c.paint[p]
}

// Count returns the number of points that have been painted at all.
func (c *Canvas) Count() int {
	return len(c.paint)
}

// CountColor returns the number of points painted with the given color.
func (c *Canvas) CountColor(color int) int {
	var n int
	for _, c := range c.paint {
		if c == color {
			n++
		}
	}
	return n
}

// Draw calls the provided drawFn for every square on the canvas, regardless
// of whether it has been painted to.
func (c *Canvas) Draw(drawFn func(int, int, int)) {
	width := c.Width()
	height := c.Height()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := c.paint[point.Point2D{
				X: x + c.minCorner.X,
				Y: y + c.minCorner.Y,
			}]
			drawFn(x, y, color)
		}
	}
}

// Width returns the width of the canvas
func (c *Canvas) Width() int {
	return c.maxCorner.X - c.minCorner.X + 1
}

// Height returns the height of the canvas
func (c *Canvas) Height() int {
	return c.maxCorner.Y - c.minCorner.Y + 1
}

func (c *Canvas) adjustSizeIfNeeded(p point.Point2D) {
	if p.X < c.minCorner.X {
		c.minCorner.X = p.X
	}
	if p.Y < c.minCorner.Y {
		c.minCorner.Y = p.Y
	}
	if p.X > c.maxCorner.X {
		c.maxCorner.X = p.X
	}
	if p.Y > c.maxCorner.Y {
		c.maxCorner.Y = p.Y
	}
}

// PrintTo prints the image on the canvas to the given writer.
func (c *Canvas) PrintTo(w io.Writer) {
	black := color.New(color.BgBlack)
	white := color.New(color.BgWhite)

	for y := c.minCorner.Y; y <= c.maxCorner.Y; y++ {
		for x := c.minCorner.X; x <= c.maxCorner.X; x++ {
			val := c.paint[point.Point2D{X: x, Y: y}]
			if val == 0 {
				black.Fprint(w, " ")
			} else if val == 1 {
				white.Fprint(w, " ")
			} else {
				panic(fmt.Errorf("unexpected color %d", val))
			}
		}
		fmt.Fprintln(w)
	}
}
