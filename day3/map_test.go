package day3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	m := NewMap()
	m.Set(1, 2, 0)
	m.Set(2, 3, 0)
	m.Set(2, 3, 1)
	m.Set(2, 3, 1)

	assert.Equal(t, 0, m.Count(1, 1))
	assert.Equal(t, 1, m.Count(1, 2))
	assert.Equal(t, 2, m.Count(2, 3))
}

func TestApplyPathSegment(t *testing.T) {
	m := NewMap()
	seg := PathSegment{Right, 3}

	var x, y int
	x, y = seg.Apply(m, x, y, 0)
	assert.Equal(t, 3, x)
	assert.Equal(t, 0, y)
	assert.Equal(t, 1, m.Count(1, 0))
	assert.Equal(t, 1, m.Count(3, 0))
	assert.Equal(t, 0, m.Count(0, 0))
	assert.Equal(t, 0, m.Count(4, 0))

	seg.Direction = Down
	seg.Length = 2
	x, y = seg.Apply(m, x, y, 0)
	assert.Equal(t, 3, x)
	assert.Equal(t, -2, y)
	assert.Equal(t, 1, m.Count(1, 0))
	assert.Equal(t, 1, m.Count(3, 0))
	assert.Equal(t, 1, m.Count(3, -1))
	assert.Equal(t, 1, m.Count(3, -2))
	assert.Equal(t, 0, m.Count(3, 1))
	assert.Equal(t, 0, m.Count(3, -3))
}

func TestApplyPath(t *testing.T) {
	m := NewMap()
	p, err := PathFromString("R3,D2")
	assert.NoError(t, err)

	p.Apply(m, 0)
	assert.Equal(t, 1, m.Count(1, 0))
	assert.Equal(t, 1, m.Count(3, 0))
	assert.Equal(t, 1, m.Count(3, -1))
	assert.Equal(t, 1, m.Count(3, -2))
	assert.Equal(t, 0, m.Count(0, 0))
	assert.Equal(t, 0, m.Count(4, 0))
	assert.Equal(t, 0, m.Count(3, 1))
	assert.Equal(t, 0, m.Count(3, -3))
}

func TestNearestIntersection(t *testing.T) {
	cases := []struct {
		p1 string
		p2 string
		x  int
		y  int
		d  int
	}{
		{
			"R8,U5,L5,D3",
			"U7,R6,D4,L4",
			3, 3, 6,
		},
		{
			"R75,D30,R83,U83,L12,D49,R71,U7,L72",
			"U62,R66,U55,R34,D71,R55,D58,R83",
			155, 4, 159,
		},
		{
			"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			124, 11, 135,
		},
	}

	for _, c := range cases {
		name := fmt.Sprintf("nearest intersection for %s and %s is %d", c.p1, c.p2, c.d)
		t.Run(name, func(t *testing.T) {
			p1, err := PathFromString(c.p1)
			assert.NoError(t, err)
			p2, err := PathFromString(c.p2)
			assert.NoError(t, err)

			m := NewMap()
			p1.Apply(m, 0)
			p2.Apply(m, 1)

			x, y, d := m.NearestIntersection()
			assert.Equal(t, c.x, x)
			assert.Equal(t, c.y, y)
			assert.Equal(t, c.d, d)
		})
	}
}
