package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/mjm/advent-of-code-2019/day12"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	sys, err := day12.NewSystemFromString(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	sys.Advance(1000)
	log.Printf("The total energy in the system after 1000 steps is %d", sys.Energy())

	sys, err = day12.NewSystemFromString(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("The first step where the state repeats is %d", sys.FirstRepetition())
}
