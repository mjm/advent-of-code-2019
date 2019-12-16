package main

import (
	"log"

	"github.com/mjm/advent-of-code-2019/day14"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

const oneTrillion = 1000000000000

func main() {
	table, err := day14.TableFromString(input.ReadString())
	if err != nil {
		log.Fatal(err)
	}

	ore := table.RequiredOre(day14.OneFuel)
	log.Printf("To produce one fuel, we need %d ore.", ore)

	// Part 2 might be a cop-out: we're just searching for a small-ish range and then
	// trying each till we find the largest value that is under 1 trillion ore

	lowerBound := oneTrillion / ore
	log.Printf("With 1 trillion ore, we can make a minimum of %d fuel.", lowerBound)

	var upperBound int
	for i := lowerBound; true; i += 50000 {
		ore = table.RequiredOre(day14.Quantity{Material: "FUEL", Amount: i})
		if ore < oneTrillion {
			lowerBound = i
			log.Printf("Lower bound shifted up to %d fuel", lowerBound)
		} else {
			upperBound = i
			log.Printf("Upper bound found: %d fuel", upperBound)
			break
		}
	}

	var i int
	for i = lowerBound; i < upperBound; i++ {
		ore = table.RequiredOre(day14.Quantity{Material: "FUEL", Amount: i})
		if ore > oneTrillion {
			break
		}
	}

	log.Printf("With 1 trillion ore, we can make %d fuel", i-1)
}
