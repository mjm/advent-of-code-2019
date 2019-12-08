package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/mjm/advent-of-code-2019/day6"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	tree, err := day6.TreeFromString(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("The total number of orbits in this map is %d", tree.TotalDepths("COM"))
	log.Printf("The minimum number of transfers from you to Santa is %d", tree.Distance("YOU", "SAN")-2)
}
