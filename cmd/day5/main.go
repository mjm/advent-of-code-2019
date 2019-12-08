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

	original, err := day5.LoadFromString(string(contents))
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
