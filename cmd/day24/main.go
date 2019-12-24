package main

import (
	"log"

	"github.com/mjm/advent-of-code-2019/day24"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

func main() {
	g := day24.GridFromString(input.ReadString())

	log.Printf("The biodiversity rating for this layout is %d", g.Simulate())
}
