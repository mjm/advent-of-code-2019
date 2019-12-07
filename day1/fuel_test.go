package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuelRequired(t *testing.T) {
	cases := []struct {
		mass         int
		fuelRequired int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, c := range cases {
		name := fmt.Sprintf("mass %d requires %d fuel", c.mass, c.fuelRequired)
		t.Run(name, func(t *testing.T) {
			m := NewModule(c.mass)
			assert.Equal(t, c.fuelRequired, FuelRequired(m))
		})
	}
}

func TestTotalFuelRequired(t *testing.T) {
	cases := []struct {
		mass         int
		fuelRequired int
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, c := range cases {
		name := fmt.Sprintf("mass %d requires %d fuel", c.mass, c.fuelRequired)
		t.Run(name, func(t *testing.T) {
			m := NewModule(c.mass)
			assert.Equal(t, c.fuelRequired, TotalFuelRequired(m))
		})
	}
}
