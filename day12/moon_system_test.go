package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSystemFromString(t *testing.T) {
	s := `<x=5, y=13, z=-3>
<x=18, y=-7, z=13>
<x=16, y=3, z=4>
<x=0, y=8, z=8>`
	sys, err := NewSystemFromString(s)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(sys.Moons))
}

func TestAdvance(t *testing.T) {
	s := `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`
	sys, err := NewSystemFromString(s)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(sys.Moons))

	sys.Advance(1)
	assert.Equal(t, Vec3D{2, -1, 1}, sys.Moons[0].Pos)
	assert.Equal(t, Vec3D{3, -1, -1}, sys.Moons[0].Vel)
	assert.Equal(t, Vec3D{3, -7, -4}, sys.Moons[1].Pos)
	assert.Equal(t, Vec3D{1, 3, 3}, sys.Moons[1].Vel)
	assert.Equal(t, Vec3D{1, -7, 5}, sys.Moons[2].Pos)
	assert.Equal(t, Vec3D{-3, 1, -3}, sys.Moons[2].Vel)
	assert.Equal(t, Vec3D{2, 2, 0}, sys.Moons[3].Pos)
	assert.Equal(t, Vec3D{-1, -3, 1}, sys.Moons[3].Vel)

	sys.Advance(9)
	assert.Equal(t, Vec3D{2, 1, -3}, sys.Moons[0].Pos)
	assert.Equal(t, Vec3D{-3, -2, 1}, sys.Moons[0].Vel)
	assert.Equal(t, Vec3D{1, -8, 0}, sys.Moons[1].Pos)
	assert.Equal(t, Vec3D{-1, 1, 3}, sys.Moons[1].Vel)
	assert.Equal(t, Vec3D{3, -6, 1}, sys.Moons[2].Pos)
	assert.Equal(t, Vec3D{3, 2, -3}, sys.Moons[2].Vel)
	assert.Equal(t, Vec3D{2, 0, 4}, sys.Moons[3].Pos)
	assert.Equal(t, Vec3D{1, -1, -1}, sys.Moons[3].Vel)
}

func TestSystemEnergy(t *testing.T) {
	s := `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`
	sys, err := NewSystemFromString(s)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(sys.Moons))

	sys.Advance(10)
	assert.Equal(t, 179, sys.Energy())

	s = `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`
	sys, err = NewSystemFromString(s)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(sys.Moons))

	sys.Advance(100)
	assert.Equal(t, 1940, sys.Energy())
}

func TestFirstRepetition(t *testing.T) {
	s := `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`
	sys, err := NewSystemFromString(s)
	assert.NoError(t, err)

	assert.Equal(t, 2772, sys.FirstRepetition())

	s = `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`
	sys, err = NewSystemFromString(s)
	assert.NoError(t, err)

	assert.Equal(t, 4686774924, sys.FirstRepetition())
}
