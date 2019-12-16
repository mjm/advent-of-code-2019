package main

import (
	"log"

	"github.com/mjm/advent-of-code-2019/day5"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

func main() {
	original, err := day5.LoadFromString(input.ReadString())
	if err != nil {
		log.Fatal(err)
	}

	vm := original.Clone()
	part1(vm)

	vm = original.Clone()
	part2(vm)
}

func part1(vm *day5.VM) {
	done := make(chan []int)
	go func() {
		vm.Input <- 1
		var codes []int
		for code := range vm.Output {
			codes = append(codes, code)
		}
		done <- codes
	}()

	if err := vm.Execute(); err != nil {
		log.Fatal(err)
	}

	codes := <-done
	log.Println(codes)
}

func part2(vm *day5.VM) {
	done := make(chan int)
	go func() {
		vm.Input <- 5
		code := <-vm.Output
		done <- code
	}()

	if err := vm.Execute(); err != nil {
		log.Fatal(err)
	}

	code := <-done
	log.Printf("The diagnostic code is %d", code)
}
