package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type VM struct {
	Memory []int
	Input  chan int
	Output chan int
	pc     int
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
		Input:  make(chan int),
		Output: make(chan int),
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

		case OpHalt:
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
	case OpAdd, OpMultiply:
		return 3, nil
	case OpInput, OpOutput:
		return 1, nil
	case OpHalt:
		return 0, nil
	default:
		return 0, fmt.Errorf("unrecognized opcode %d", op)
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

func (vm *VM) Clone() *VM {
	memory := make([]int, len(vm.Memory))
	copy(memory, vm.Memory)
	return &VM{
		Memory: memory,
		pc:     vm.pc,
	}
}
