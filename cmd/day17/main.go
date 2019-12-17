package main

import (
	"log"
	"os"

	"github.com/mjm/advent-of-code-2019/day17"
	"github.com/mjm/advent-of-code-2019/pkg/input"
	"github.com/mjm/advent-of-code-2019/pkg/intcode"
)

func main() {
	vm, err := intcode.LoadFromString(input.ReadString())
	if err != nil {
		log.Fatal(err)
	}

	canvas := day17.NewCanvas()
	finder := day17.NewIntersectionFinder(canvas)

	go func() {
		if err := vm.Execute(); err != nil {
			log.Fatal(err)
		}
	}()
	finder.BuildMap(vm.Output)
	canvas.PrintTo(os.Stdout)
	log.Printf("The sum of the alignment parameters is %d", finder.AlignmentParameters())
}
