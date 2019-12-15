package day10

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type loadMapCheck struct {
	x        int
	y        int
	asteroid bool
}

func TestLoadMap(t *testing.T) {
	cases := []struct {
		s      string
		checks []loadMapCheck
	}{
		{
			s: `.#..#
.....
#####
....#
...##`,
			checks: []loadMapCheck{
				{1, 0, true},
				{4, 0, true},
				{2, 2, true},
				{0, 0, false},
				{2, 4, false},
			},
		},
		{
			s: `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`,
			checks: []loadMapCheck{
				{0, 0, true},
				{1, 0, false},
				{1, 1, true},
				{0, 3, true},
				{8, 1, true},
				{0, 2, false},
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("load map case %d", i), func(t *testing.T) {
			m, err := LoadFromString(c.s)
			assert.NoError(t, err)

			for _, check := range c.checks {
				if check.asteroid {
					assert.True(t, m.At(check.x, check.y))
				} else {
					assert.False(t, m.At(check.x, check.y))
				}
			}
		})
	}
}

type visibleCheck struct {
	x     int
	y     int
	count int
}

func TestVisibleAsteroids(t *testing.T) {
	cases := []struct {
		s      string
		checks []visibleCheck
	}{
		{
			s: `.#..#
.....
#####
....#
...##`,
			checks: []visibleCheck{
				{1, 0, 7},
				{4, 0, 7},
				{2, 2, 7},
				{3, 4, 8},
			},
		},
		{
			s: `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`,
			checks: []visibleCheck{
				{1, 2, 35},
			},
		},
		{
			s: `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`,
			checks: []visibleCheck{
				{5, 8, 33},
			},
		},
		{
			s: `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`,
			checks: []visibleCheck{
				{6, 3, 41},
			},
		},
		{
			s: `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
			checks: []visibleCheck{
				{11, 13, 210},
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("visible asteroids case %d", i), func(t *testing.T) {
			m, err := LoadFromString(c.s)
			assert.NoError(t, err)

			for _, check := range c.checks {
				assert.Equal(t, check.count, m.VisibleAsteroids(check.x, check.y))
			}
		})
	}
}

func TestBestAsteroid(t *testing.T) {
	cases := []struct {
		s string
		x int
		y int
		n int
	}{
		{
			s: `.#..#
.....
#####
....#
...##`,
			x: 3,
			y: 4,
			n: 8,
		},
		{
			s: `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`,
			x: 1,
			y: 2,
			n: 35,
		},
		{
			s: `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`,
			x: 5,
			y: 8,
			n: 33,
		},
		{
			s: `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`,
			x: 6,
			y: 3,
			n: 41,
		},
		{
			s: `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
			x: 11,
			y: 13,
			n: 210,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("best asteroid case %d", i), func(t *testing.T) {
			m, err := LoadFromString(c.s)
			assert.NoError(t, err)

			x, y, n := m.BestAsteroid()
			assert.Equal(t, c.x, x)
			assert.Equal(t, c.y, y)
			assert.Equal(t, c.n, n)
		})
	}
}

func TestVaporized(t *testing.T) {
	s := `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
	m, err := LoadFromString(s)
	assert.NoError(t, err)

	cases := []struct {
		n int
		x int
		y int
	}{
		{0, 11, 12},
		{1, 12, 1},
		{2, 12, 2},
		{9, 12, 8},
		{19, 16, 0},
		{49, 16, 9},
		{99, 10, 16},
		{198, 9, 6},
		{199, 8, 2},
		{200, 10, 9},
		{298, 11, 1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("asteroid #%d to be vaporized is at %d, %d", c.n, c.x, c.y), func(t *testing.T) {
			x, y := m.Vaporized(11, 13, c.n)
			assert.Equal(t, c.x, x)
			assert.Equal(t, c.y, y)
		})
	}
}
