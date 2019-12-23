package main

import (
	"log"

	"github.com/mjm/advent-of-code-2019/day23"
	"github.com/mjm/advent-of-code-2019/pkg/input"
	"github.com/mjm/advent-of-code-2019/pkg/intcode"
)

func main() {
	template, err := intcode.LoadFromString(input.ReadString())
	if err != nil {
		log.Fatal(err)
	}

	net := day23.NewNetwork()
	for i := 0; i < 50; i++ {
		net.Register(template.Clone())
	}

	x, y := net.Listen()
	log.Printf("The packet sent to 255 was %d, %d", x, y)
}
