package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mjm/advent-of-code-2019/day21"
	"github.com/mjm/advent-of-code-2019/pkg/input"
	"github.com/mjm/advent-of-code-2019/pkg/intcode"
)

func main() {
	vm, err := intcode.LoadFromString(input.ReadString())
	if err != nil {
		log.Fatal(err)
	}

	droid := day21.NewDroid(day21.Program)
	damage, err := droid.Run(vm.Clone())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return
	}

	log.Printf("The damage to the hull is %d", damage)

}
