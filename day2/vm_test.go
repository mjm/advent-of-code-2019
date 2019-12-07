package day2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromString(t *testing.T) {
	cases := []struct {
		program string
		memory  []int
	}{
		{
			program: "1,9,10,3,2,3,11,0,99,30,40,50",
			memory:  []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			program: "1,0,0,0,99",
			memory:  []int{1, 0, 0, 0, 99},
		},
		{
			program: "2,3,0,3,99",
			memory:  []int{2, 3, 0, 3, 99},
		},
		{
			program: "2,4,4,5,99,0",
			memory:  []int{2, 4, 4, 5, 99, 0},
		},
		{
			program: "1,1,1,4,99,5,6,0,99",
			memory:  []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("load case %d", i), func(t *testing.T) {
			vm, err := LoadFromString(c.program)
			assert.NoError(t, err)
			assert.Equal(t, c.memory, vm.Memory)
		})
	}
}

func TestExecute(t *testing.T) {
	cases := []struct {
		program string
		memory  []int
	}{
		{
			program: "1,9,10,3,2,3,11,0,99,30,40,50",
			memory: []int{
				3500, 9, 10, 70,
				2, 3, 11, 0,
				99,
				30, 40, 50,
			},
		},
		{
			program: "1,0,0,0,99",
			memory:  []int{2, 0, 0, 0, 99},
		},
		{
			program: "2,3,0,3,99",
			memory:  []int{2, 3, 0, 6, 99},
		},
		{
			program: "2,4,4,5,99,0",
			memory:  []int{2, 4, 4, 5, 99, 9801},
		},
		{
			program: "1,1,1,4,99,5,6,0,99",
			memory:  []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("load case %d", i), func(t *testing.T) {
			vm, err := LoadFromString(c.program)
			assert.NoError(t, err)

			err = vm.Execute()
			assert.NoError(t, err)

			assert.Equal(t, c.memory, vm.Memory)
		})
	}
}
