package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/mjm/advent-of-code-2019/day5"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	vm, err := day5.LoadFromString(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	var codes []int

	go func() {
		vm.Input <- 1
		for {
			code := <-vm.Output
			codes = append(codes, code)
		}
	}()

	if err := vm.Execute(); err != nil {
		log.Fatal(err)
	}

	log.Println(codes)
}
