package day9

import (
	"fmt"
	"strconv"
	"strings"
)

type VM struct {
	Memory       []int
	Input        chan int
	Output       chan int
	pc           int
	relativeBase int
}

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
		Input:  make(chan int, 1),
		Output: make(chan int, 1),
	}, nil
}

func (vm *VM) Execute() error {
	for {
		inst, err := vm.scanInstruction()
		if err != nil {
			return err
		}

		switch inst.Op {
		case OpAdd:
			n1 := inst.Get(vm, 0)
			n2 := inst.Get(vm, 1)
			inst.Set(vm, 2, n1+n2)

		case OpMultiply:
			n1 := inst.Get(vm, 0)
			n2 := inst.Get(vm, 1)
			inst.Set(vm, 2, n1*n2)

		case OpInput:
			n := <-vm.Input
			inst.Set(vm, 0, n)

		case OpOutput:
			n := inst.Get(vm, 0)
			vm.Output <- n

		case OpJumpIfTrue:
			n := inst.Get(vm, 0)
			if n != 0 {
				vm.pc = inst.Get(vm, 1)
			}

		case OpJumpIfFalse:
			n := inst.Get(vm, 0)
			if n == 0 {
				vm.pc = inst.Get(vm, 1)
			}

		case OpLessThan:
			a := inst.Get(vm, 0)
			b := inst.Get(vm, 1)
			if a < b {
				inst.Set(vm, 2, 1)
			} else {
				inst.Set(vm, 2, 0)
			}

		case OpEquals:
			a := inst.Get(vm, 0)
			b := inst.Get(vm, 1)
			if a == b {
				inst.Set(vm, 2, 1)
			} else {
				inst.Set(vm, 2, 0)
			}

		case OpSetRelativeBase:
			n := inst.Get(vm, 0)
			vm.relativeBase += n

		case OpHalt:
			close(vm.Output)
			return nil

		default:
			return fmt.Errorf("unrecognized opcode %d", inst.Op)
		}
	}
}

func (vm *VM) scanInstruction() (*Instruction, error) {
	var inst Instruction

	value := vm.AtOffset(0)
	inst.Op = Op(value % 100)
	modes := value / 100

	numParams, err := paramCount(inst.Op)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= numParams; i++ {
		param := Param{Value: vm.AtOffset(i)}
		param.Mode = ParamMode(modes % 10)
		modes = modes / 10

		inst.Params = append(inst.Params, param)
	}

	vm.pc += numParams + 1
	return &inst, nil
}

func paramCount(op Op) (int, error) {
	switch op {
	case OpAdd, OpMultiply, OpLessThan, OpEquals:
		return 3, nil
	case OpJumpIfFalse, OpJumpIfTrue:
		return 2, nil
	case OpInput, OpOutput, OpSetRelativeBase:
		return 1, nil
	case OpHalt:
		return 0, nil
	default:
		return 0, fmt.Errorf("unrecognized opcode %d", op)
	}
}

func (vm *VM) At(i int) int {
	vm.growMemoryIfNeeded(i)
	return vm.Memory[i]
}

func (vm *VM) AtOffset(i int) int {
	return vm.At(vm.pc + i)
}

func (vm *VM) Set(i int, value int) {
	vm.growMemoryIfNeeded(i)
	vm.Memory[i] = value
}

func (vm *VM) growMemoryIfNeeded(i int) {
	if len(vm.Memory) > i {
		return
	}

	if cap(vm.Memory) > i {
		// already have enough capacity, so just re-slice
		vm.Memory = vm.Memory[:i+1]
		return
	}

	newCapacity := cap(vm.Memory) * 2
	if newCapacity < i {
		newCapacity = i + 1
	}
	newMemory := make([]int, i+1, newCapacity)
	copy(newMemory, vm.Memory)
	vm.Memory = newMemory
}

func (vm *VM) Clone() *VM {
	memory := make([]int, len(vm.Memory))
	copy(memory, vm.Memory)
	return &VM{
		Memory:       memory,
		Input:        make(chan int, 1),
		Output:       make(chan int, 1),
		pc:           vm.pc,
		relativeBase: vm.relativeBase,
	}
}

func (vm *VM) PipeTo(other *VM) {
	for val := range vm.Output {
		other.Input <- val
	}
}
