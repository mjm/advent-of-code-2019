package intcode

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
		{
			program: "1002,4,3,4,33",
			memory:  []int{1002, 4, 3, 4, 99},
		},
		{
			program: "1101,100,-1,4,0",
			memory:  []int{1101, 100, -1, 4, 99},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("execute case %d", i), func(t *testing.T) {
			vm, err := LoadFromString(c.program)
			assert.NoError(t, err)

			err = vm.Execute()
			assert.NoError(t, err)

			assert.Equal(t, c.memory, vm.Memory)
		})
	}
}

func TestInputOutput(t *testing.T) {
	cases := []struct {
		program string
		input   []int
		output  []int
	}{
		{
			program: "3,9,8,9,10,9,4,9,99,-1,8",
			input:   []int{8},
			output:  []int{1},
		},
		{
			program: "3,9,8,9,10,9,4,9,99,-1,8",
			input:   []int{10},
			output:  []int{0},
		},
		{
			program: "3,9,8,9,10,9,4,9,99,-1,8",
			input:   []int{7},
			output:  []int{0},
		},
		{
			program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			input:   []int{0},
			output:  []int{0},
		},
		{
			program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1",
			input:   []int{4},
			output:  []int{1},
		},
		{
			program: "1102,34915192,34915192,7,4,7,99,0",
			input:   nil,
			output:  []int{1219070632396864},
		},
		{
			program: "104,1125899906842624,99",
			input:   nil,
			output:  []int{1125899906842624},
		},
		{
			program: "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99",
			input:   nil,
			output:  []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("input/output case %d", i), func(t *testing.T) {
			vm, err := LoadFromString(c.program)
			assert.NoError(t, err)

			input := make(chan int)
			done := make(chan []int)
			go func() {
				for _, val := range c.input {
					input <- val
				}
				var outputs []int
				for n := range vm.Output {
					outputs = append(outputs, n)
				}
				done <- outputs
			}()

			vm.SetInputChan(input)
			err = vm.Execute()
			assert.NoError(t, err)
			outputs := <-done
			assert.Equal(t, c.output, outputs)
		})
	}
}

func TestClone(t *testing.T) {
	vm, err := LoadFromString("1,0,0,0,99")
	assert.NoError(t, err)

	vm2 := vm.Clone()
	vm.Set(1, 12)
	vm2.Set(2, 2)

	assert.Equal(t, vm.Memory, []int{1, 12, 0, 0, 99})
	assert.Equal(t, vm2.Memory, []int{1, 0, 2, 0, 99})
}
