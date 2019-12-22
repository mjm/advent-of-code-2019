package main

import (
	"log"

	"github.com/mjm/advent-of-code-2019/day22"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

func main() {
	ts, err := day22.ParseAllTechniques(input.ReadString())
	if err != nil {
		log.Fatal(err)
	}

	card := day22.ShuffleFind(ts, 2019, 10007)
	log.Printf("The position of card 2019 is %d", card)

	card = day22.ShuffleN(ts, 2020, 119315717514047, 101741582076661)
	log.Printf("The card at 2020 after an awful lot of shuffles is %d", card)
}
