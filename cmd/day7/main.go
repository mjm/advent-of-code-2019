package main

import (
	"log"

	"github.com/mjm/advent-of-code-2019/day7"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

func main() {
	template, err := day7.LoadFromString(input.ReadString())
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
