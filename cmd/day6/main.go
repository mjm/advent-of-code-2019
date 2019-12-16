package main

import (
	"log"

	"github.com/mjm/advent-of-code-2019/day6"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

func main() {
	tree, err := day6.TreeFromString(input.ReadString())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("The total number of orbits in this map is %d", tree.TotalDepths("COM"))
	log.Printf("The minimum number of transfers from you to Santa is %d", tree.Distance("YOU", "SAN")-2)
}
