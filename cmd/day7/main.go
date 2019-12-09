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
	part2(template)
}

func part1(template *day7.VM) {
	a := day7.NewAmplifierArray(template, 5)
	highestOutput, highestSequence := a.HighestSignal([]int{0, 1, 2, 3, 4})
	log.Printf("The highest signal is %d, produced by phase settings %v", highestOutput, highestSequence)
}

func part2(template *day7.VM) {
	a := day7.NewAmplifierArray(template, 5)
	highestOutput, highestSequence := a.HighestFeedbackSignal([]int{5, 6, 7, 8, 9})
	log.Printf("The highest feedback signal is %d, produced by phase settings %v", highestOutput, highestSequence)
}
