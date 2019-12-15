package day12

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Moon is a single moon in the simulation.
type Moon struct {
	// Pos is the position of the moon.
	Pos Vec3D
	// Vel is the velocity of the moon.
	Vel Vec3D
}

var errInvalidString = errors.New("moon string was not valid")

var moonRegex = regexp.MustCompile("<x=(-?\\d+), y=(-?\\d+), z=(-?\\d+)>")

// NewMoonFromString reads the position coordinates from a string and creates
// a new moon at that position. The initial velocity of the moon is zero.
func NewMoonFromString(s string) (*Moon, error) {
	matches := moonRegex.FindStringSubmatch(s)
	if len(matches) == 0 {
		return nil, errInvalidString
	}

	x, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, fmt.Errorf("could not parse x value: %w", err)
	}
	y, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, fmt.Errorf("could not parse y value: %w", err)
	}
	z, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, fmt.Errorf("could not parse z value: %w", err)
	}

	return &Moon{
		Pos: Vec3D{X: x, Y: y, Z: z},
	}, nil
}

// PotentialEnergy calculates the potential energy of the moon.
func (m *Moon) PotentialEnergy() int {
	return abs(m.Pos.X) + abs(m.Pos.Y) + abs(m.Pos.Z)
}

// KineticEnergy calculates the kinetic energy of the moon.
func (m *Moon) KineticEnergy() int {
	return abs(m.Vel.X) + abs(m.Vel.Y) + abs(m.Vel.Z)
}

// Energy is the product of the moon's potential and kinetic energy.
func (m *Moon) Energy() int {
	return m.PotentialEnergy() * m.KineticEnergy()
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
