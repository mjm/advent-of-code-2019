package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mjm/advent-of-code-2019/day9"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	template, err := day9.LoadFromString(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	part1(template)
	part2(template)
}

func part1(template *day9.VM) {
	vm := template.Clone()
	vm.Input <- 1

	done := make(chan int)
	go func() {
		var values []int
		for out := range vm.Output {
			values = append(values, out)
		}

		if len(values) == 1 {
			done <- values[0]
		} else {
			panic(fmt.Errorf("got some failing opcodes: %v", values))
		}
	}()

	if err := vm.Execute(); err != nil {
		log.Fatal(err)
	}

	boostKeycode := <-done
	log.Printf("The BOOST keycode is %d", boostKeycode)
}

func part2(template *day9.VM) {
	vm := template.Clone()
	vm.Input <- 2

	done := make(chan int)
	go func() {
		n := <-vm.Output
		done <- n
	}()

	if err := vm.Execute(); err != nil {
		log.Fatal(err)
	}

	coordinates := <-done
	log.Printf("The coordinates of the distress signal are %d", coordinates)
}
