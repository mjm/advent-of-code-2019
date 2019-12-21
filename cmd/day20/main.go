package main

import (
	"log"

	"github.com/mjm/advent-of-code-2019/day20"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

func main() {
	m := day20.MapFromString(input.ReadString())
	path := m.ShortestPath()
	log.Printf("The fewest steps to get to the exit is %d", len(path)-1)
	path = m.ShortestPathLevels()
	log.Printf("Ok, but if you actually respect the levels, it's actually %d steps", len(path)-1)
}
