package day12

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMoonFromString(t *testing.T) {
	cases := []struct {
		s   string
		pos Vec3D
	}{
		{`<x=5, y=13, z=-3>`, Vec3D{5, 13, -3}},
		{`<x=18, y=-7, z=13>`, Vec3D{18, -7, 13}},
		{`<x=16, y=3, z=4>`, Vec3D{16, 3, 4}},
		{`<x=0, y=8, z=8>`, Vec3D{0, 8, 8}},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("new moon from string case %d", i), func(t *testing.T) {
			moon, err := NewMoonFromString(c.s)
			assert.NoError(t, err)
			assert.Equal(t, c.pos, moon.Pos)
			assert.Equal(t, Vec3D{0, 0, 0}, moon.Vel)
		})
	}
}

func TestNewMoonFromStringFailures(t *testing.T) {
	cases := []struct {
		s   string
		err string
	}{
		{``, "moon string was not valid"},
		{`x=1, y=2, z=3`, "moon string was not valid"},
		{`<y=1, x=2, z=3>`, "moon string was not valid"},
		{`<x=abc, y=2, z=3>`, "moon string was not valid"},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("new moon from string failure case %d", i), func(t *testing.T) {
			_, err := NewMoonFromString(c.s)
			assert.EqualError(t, err, c.err)
		})
	}
}

func TestMoonEnergy(t *testing.T) {
	cases := []struct {
		m     *Moon
		pot   int
		kin   int
		total int
	}{
		{
			m:     &Moon{Pos: Vec3D{2, 1, -3}, Vel: Vec3D{-3, -2, 1}},
			pot:   6,
			kin:   6,
			total: 36,
		},
		{
			m:     &Moon{Pos: Vec3D{1, -8, 0}, Vel: Vec3D{-1, 1, 3}},
			pot:   9,
			kin:   5,
			total: 45,
		},
		{
			m:     &Moon{Pos: Vec3D{3, -6, 1}, Vel: Vec3D{3, 2, -3}},
			pot:   10,
			kin:   8,
			total: 80,
		},
		{
			m:     &Moon{Pos: Vec3D{2, 0, 4}, Vel: Vec3D{1, -1, -1}},
			pot:   6,
			kin:   3,
			total: 18,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("moon energy case %d", i), func(t *testing.T) {
			assert.Equal(t, c.pot, c.m.PotentialEnergy())
			assert.Equal(t, c.kin, c.m.KineticEnergy())
			assert.Equal(t, c.total, c.m.Energy())
		})
	}
}
