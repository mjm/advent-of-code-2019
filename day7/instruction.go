package day7

type Instruction struct {
	Op     Op
	Params []Param
}

func (inst Instruction) Get(vm *VM, i int) int {
	p := inst.Params[i]
	switch p.Mode {
	case ModePosition:
		return vm.At(p.Value)
	case ModeImmediate:
		return p.Value
	default:
		panic("unexpected param mode")
	}
}

func (inst Instruction) Set(vm *VM, i int, value int) {
	p := inst.Params[i]
	vm.Set(p.Value, value)
}

type Op int

const (
	OpAdd         Op = 1
	OpMultiply    Op = 2
	OpInput       Op = 3
	OpOutput      Op = 4
	OpJumpIfTrue  Op = 5
	OpJumpIfFalse Op = 6
	OpLessThan    Op = 7
	OpEquals      Op = 8
	OpHalt        Op = 99
)

type ParamMode int

const (
	ModePosition  ParamMode = 0
	ModeImmediate ParamMode = 1
)

type Param struct {
	Value int
	Mode  ParamMode
}
