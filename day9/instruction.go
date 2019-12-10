package day9

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
	case ModeRelative:
		return vm.At(p.Value + vm.relativeBase)
	default:
		panic("unexpected param mode")
	}
}

func (inst Instruction) Set(vm *VM, i int, value int) {
	p := inst.Params[i]
	switch p.Mode {
	case ModePosition:
		vm.Set(p.Value, value)
	case ModeRelative:
		vm.Set(p.Value+vm.relativeBase, value)
	default:
		panic("unexpected param mode")
	}
}

type Op int

const (
	OpAdd             Op = 1
	OpMultiply        Op = 2
	OpInput           Op = 3
	OpOutput          Op = 4
	OpJumpIfTrue      Op = 5
	OpJumpIfFalse     Op = 6
	OpLessThan        Op = 7
	OpEquals          Op = 8
	OpSetRelativeBase Op = 9
	OpHalt            Op = 99
)

type ParamMode int

const (
	ModePosition  ParamMode = 0
	ModeImmediate ParamMode = 1
	ModeRelative  ParamMode = 2
)

type Param struct {
	Value int
	Mode  ParamMode
}
