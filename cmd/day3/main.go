package main

import (
	"log"
	"strings"

	"github.com/mjm/advent-of-code-2019/day3"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

func main() {
	lines := strings.Split(input.ReadString(), "\n")
	m := day3.NewMap()

	p1, err := day3.PathFromString(lines[0])
	if err != nil {
		log.Fatal(err)
	}
	p2, err := day3.PathFromString(lines[1])
	if err != nil {
		log.Fatal(err)
	}

	p1.Apply(m, 0)
	p2.Apply(m, 1)

	x, y, d := m.NearestIntersection()
	log.Printf("Nearest intersection is at (%d, %d), distance is %d", x, y, d)

	x, y, steps := m.ShortestIntersection()
	log.Printf("Shortest intersection is at (%d, %d), with %d steps", x, y, steps)
}
