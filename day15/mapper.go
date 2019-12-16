package day15

import (
	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// Mapper is responsible for constructing a map of a location
// by controlling a robot with an Intcode program.
type Mapper struct {
	canvas    *Canvas
	location  point.Point2D
	direction Direction
}

// NewMapper creates a new mapper using a given canvas to draw the map.
func NewMapper(c *Canvas) *Mapper {
	return &Mapper{canvas: c, direction: North}
}

// Start begins mapping using the given input and output channels. It returns
// the location of the destination.
func (m *Mapper) Start(input chan<- int, output <-chan int) point.Point2D {
	var dest point.Point2D

	m.canvas.Paint(m.location, int(TileStart))

	for !m.isDone() {
		dir := m.chooseDirection()
		input <- int(dir)

		status, ok := <-output
		if !ok {
			return dest
		}

		switch Status(status) {
		case HitWall:
			wall := dir.offset(m.location)
			m.canvas.Paint(wall, int(TileWall))
		case Moved:
			m.direction = dir
			m.location = dir.offset(m.location)
			if m.canvas.At(m.location) != int(TileStart) {
				m.canvas.Paint(m.location, int(TilePassable))
			}
		case FoundDestination:
			m.direction = dir
			m.location = dir.offset(m.location)
			dest = m.location
			if m.canvas.At(m.location) != int(TileStart) {
				m.canvas.Paint(m.location, int(TileDestination))
			}
		}
	}

	return dest
}

func (m *Mapper) isDone() bool {
	if m.canvas.At(m.location) != int(TileStart) {
		return false
	}

	dirs := []Direction{North, South, West, East}
	for _, dir := range dirs {
		p := dir.offset(m.location)
		tile := Tile(m.canvas.At(p))
		if tile == TileUnknown {
			return false
		}
	}

	// all directions from starting point have been explored
	return true
}

func (m *Mapper) chooseDirection() Direction {
	// first, check the space to the right of the current location.
	// if it's open, turn that way.
	dir := m.direction.right()
	if p := dir.offset(m.location); !m.isWall(p) {
		return dir
	}

	// if the space to the right is a wall, then try to just keep going
	// forward.
	if p := m.direction.offset(m.location); !m.isWall(p) {
		return m.direction
	}

	// if right and forward are walls, follow them and try to go left
	dir = m.direction.left()
	if p := dir.offset(m.location); !m.isWall(p) {
		return dir
	}

	// go backwards then, which is left from left.
	return dir.left()
}

func (m *Mapper) isWall(p point.Point2D) bool {
	return m.canvas.At(p) == int(TileWall)
}

// Tile is a type of tile drawn on a map.
type Tile int

const (
	// TileUnknown is a tile that hasn't been discovered
	TileUnknown Tile = iota
	// TilePassable is a tile that the droid can enter
	TilePassable
	// TileWall is a tile that is an impassable wall
	TileWall
	// TileDestination is the tile with the goal
	TileDestination
	// TileStart is the location where the droid started
	TileStart
)

// Direction is a cardinal direction that the robot can be instructed to move in.
type Direction int

// Constants for directions that can be given to the robot.
const (
	None Direction = iota
	North
	South
	West
	East
)

func (d Direction) offsetBy(p point.Point2D, n int) point.Point2D {
	q := p
	switch d {
	case North:
		q.Y -= n
	case South:
		q.Y += n
	case West:
		q.X -= n
	case East:
		q.X += n
	}
	return q
}

func (d Direction) offset(p point.Point2D) point.Point2D {
	return d.offsetBy(p, 1)
}

func (d Direction) left() Direction {
	switch d {
	case North:
		return West
	case South:
		return East
	case West:
		return South
	case East:
		return North
	default:
		panic("invalid direction")
	}
}

func (d Direction) right() Direction {
	switch d {
	case North:
		return East
	case South:
		return West
	case West:
		return North
	case East:
		return South
	default:
		panic("invalid direction")
	}
}

// Status is a status code that can be returned from the Intcode program.
type Status int

const (
	// HitWall signals that the droid hit a wall and was not able to move.
	HitWall Status = iota
	// Moved signals that the droid moved in the requested direction but
	// did not find the destination.
	Moved
	// FoundDestination signals that the droid moved in the requested
	// direction and found the destination.
	FoundDestination
)
