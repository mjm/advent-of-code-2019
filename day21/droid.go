package day21

import (
	"fmt"
	"strings"

	"github.com/mjm/advent-of-code-2019/pkg/intcode"
)

const Program = `NOT C J
AND D J
NOT A T
OR T J
WALK
`

type Droid struct {
	program string
}

func NewDroid(program string) *Droid {
	return &Droid{program: program}
}

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
