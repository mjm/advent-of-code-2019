package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/mjm/advent-of-code-2019/day10"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	m, err := day10.LoadFromString(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	x, y, n := m.BestAsteroid()
	log.Printf("The best asteroid is at (%d, %d) and can see %d asteroids", x, y, n)

	x1, y1 := m.Vaporized(x, y, 199)
	log.Printf("The 200th vaporized asteroid is at (%d, %d)", x1, y1)
}
