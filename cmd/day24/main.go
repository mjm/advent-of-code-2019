package main

import (
	"flag"

	"github.com/mjm/advent-of-code-2019/day24"
	"github.com/mjm/advent-of-code-2019/pkg/input"
	log "github.com/sirupsen/logrus"
)

var verbose = flag.Bool("v", false, "Enable verbose logging")

func main() {
	in := input.ReadString()
	if *verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	g := day24.GridFromString(in)
	log.Printf("The biodiversity rating for this layout is %d", g.Simulate())

	rg := day24.RecursiveGridFromString(in)
	log.Printf("The number of bugs after 200 minutes is %d", rg.Simulate(200))
}
