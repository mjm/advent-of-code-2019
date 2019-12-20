package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/mjm/advent-of-code-2019/day18"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

func main() {
	m := day18.MapFromString(input.ReadString())
	log.Printf("Shortest path is %d", m.ShortestWalk())

	input2, err := ioutil.ReadFile(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	m2 := day18.MapFromString(string(input2))
	log.Printf("Shortest path for the 4 robots is %d", m2.ShortestWalk())
}
