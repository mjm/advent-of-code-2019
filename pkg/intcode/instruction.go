package intcode

// Instruction is a single instruction read from the memory of a VM.
type Instruction struct {
	// Op is the opcode for the action the instruction should perform.
	Op Op
	// Params is the list of parameters provided to the instruction.
	Params []Param
}

// Get reads the value for a parameter of the instruction. It uses the parameter's
// mode to read from the correct location.
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

// Set updates a value in memory for a parameter of the instruction. It uses the
// parameter's mode to determine the correct location to write to.
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

// Op is an opcode for an Intcode instruction.
type Op int

// The supported opcodes for this Intcode interpreter.
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

// ParamMode represents how to use the value in an Intcode instruction parameter.
type ParamMode int

const (
	// ModePosition indicates the parameter's value refers to an absolute address
	// memory.
	ModePosition ParamMode = 0
	// ModeImmediate indicates the parameter's value should be used directly, and
	// does not correspond to another address in memory.
	ModeImmediate ParamMode = 1
	// ModeRelative indicates the parameter's value refers to an offset from the
	// VM's current relative base.
	ModeRelative ParamMode = 2
)

// Param is a single parameter from an Intcode instruction.
type Param struct {
	// Value is the integer value of the parameter.
	Value int
	// Mode is the mode that should be used to interpret the value.
	Mode ParamMode
}
