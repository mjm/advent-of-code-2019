package main

import (
	"log"

	"github.com/mjm/advent-of-code-2019/day12"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

func main() {
	in := input.ReadString()
	sys, err := day12.NewSystemFromString(in)
	if err != nil {
		log.Fatal(err)
	}

	sys.Advance(1000)
	log.Printf("The total energy in the system after 1000 steps is %d", sys.Energy())

	sys, err = day12.NewSystemFromString(in)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("The first step where the state repeats is %d", sys.FirstRepetition())
}
