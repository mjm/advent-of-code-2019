package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/mjm/advent-of-code-2019/day7"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	template, err := day7.LoadFromString(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	part1(template)
}

func part1(template *day7.VM) {
	a := day7.NewAmplifierArray(template, 5)

	var highestSequence []int
	var highestOutput int

	for nums := range day7.AllPermutations([]int{0, 1, 2, 3, 4}) {
		out := a.Run(nums)
		log.Printf("sequence %v produced signal %d", nums, out)
		if out > highestOutput {
			highestOutput = out
			highestSequence = nums
		}
	}

	log.Printf("The highest signal is %d, produced by phase settings %v", highestOutput, highestSequence)
}
