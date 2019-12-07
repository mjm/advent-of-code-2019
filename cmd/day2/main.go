package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/mjm/advent-of-code-2019/day2"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	vm, err := day2.LoadFromString(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	vm.Set(1, 12)
	vm.Set(2, 2)

	if err := vm.Execute(); err != nil {
		log.Fatal(err)
	}

	log.Printf("The value at position 0 is %d", vm.At(0))
}
