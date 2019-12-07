package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type VM struct {
	Memory []int
	pc     int
}

const (
	OpAdd      = 1
	OpMultiply = 2
	OpHalt     = 99
)

func LoadFromString(s string) (*VM, error) {
	memStrs := strings.Split(s, ",")
	memory := make([]int, 0, len(memStrs))

	for _, str := range memStrs {
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		memory = append(memory, i)
	}

	return &VM{
		Memory: memory,
	}, nil
}

func (vm *VM) Execute() error {
	for {
		opcode := vm.AtOffset(0)
		switch opcode {
		case OpAdd:
			loc1 := vm.AtOffset(1)
			loc2 := vm.AtOffset(2)
			dst := vm.AtOffset(3)

			sum := vm.At(loc1) + vm.At(loc2)
			vm.Set(dst, sum)

		case OpMultiply:
			loc1 := vm.AtOffset(1)
			loc2 := vm.AtOffset(2)
			dst := vm.AtOffset(3)

			sum := vm.At(loc1) * vm.At(loc2)
			vm.Set(dst, sum)

		case OpHalt:
			return nil

		default:
			return fmt.Errorf("unrecognized opcode %d", vm)
		}

		vm.Advance()
	}
}

func (vm *VM) At(i int) int {
	return vm.Memory[i]
}

func (vm *VM) AtOffset(i int) int {
	return vm.At(vm.pc + i)
}

func (vm *VM) Set(i int, value int) {
	vm.Memory[i] = value
}

func (vm *VM) Advance() {
	vm.pc += 4
}

func (vm *VM) Clone() *VM {
	memory := make([]int, len(vm.Memory))
	copy(memory, vm.Memory)
	return &VM{
		Memory: memory,
		pc:     vm.pc,
	}
}
