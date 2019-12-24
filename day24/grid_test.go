package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridFromString(t *testing.T) {
	g := GridFromString(`....#
#..#.
#..##
..#..
#....`)
	assert.Equal(t, 5, g.width)
	assert.Equal(t, 5, g.height)

	assert.True(t, g.hasBug(2, 3))
	assert.False(t, g.hasBug(2, 4))
}

func TestSimulateOnce(t *testing.T) {
	g := GridFromString(`....#
#..#.
#..##
..#..
#....`)
	g.SimulateOnce()
	assert.Equal(t, uint64(0b11011011101110111101001), g.bugs)

	g.SimulateOnce()
	assert.Equal(t, uint64(0b1110101000100001000011111), g.bugs)
}

func TestSimulate(t *testing.T) {
	g := GridFromString(`....#
#..#.
#..##
..#..
#....`)

	assert.Equal(t, uint64(2129920), g.Simulate())
}
