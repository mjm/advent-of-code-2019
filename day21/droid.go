package day21

import (
	"fmt"
	"strings"

	"github.com/mjm/advent-of-code-2019/pkg/intcode"
)

// WalkProgram is a working program for the droid in walking mode.
const WalkProgram = `NOT C J
AND D J
NOT A T
OR T J
WALK
`

// RunProgram is a working program for the droid in running mode.
const RunProgram = `NOT B J
NOT C T
OR T J
NOT D T
NOT T T
AND T J
NOT H T
NOT T T
OR E T
AND T J
NOT A T
OR T J
RUN
`

// Droid controls a springdroid using a program of boolean instructions for
// when to jump.
type Droid struct {
	program string
}

// NewDroid creates a new droid with a program.
func NewDroid(program string) *Droid {
	return &Droid{program: program}
}

// Run runs the given Intcode VM to control the droid, using the droid's program.
// Returns either the amount of damage to the ship's hull that the droid found,
// or an error with the output from the VM which will show where the droid fell.
func (d *Droid) Run(vm *intcode.VM) (int, error) {
	input := make(chan int, len(d.program))
	for _, c := range d.program {
		input <- int(c)
	}

	vm.SetInputChan(input)
	go vm.MustExecute()

	var b strings.Builder
	for c := range vm.Output {
		if c > 255 {
			return c, nil
		}
		b.WriteByte(byte(c))
	}
	return 0, fmt.Errorf("the droid fell in a hole.\n%s", b.String())
}
