package main

import (
	"flag"
	"github.com/mjm/advent-of-code-2019/day8"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	img, err := day8.ImageFromString(string(contents), 25, 6)
	if err != nil {
		log.Fatal(err)
	}

	counts := img.DigitCounts()
	var minZeros map[int]int
	for _, layerCounts := range counts {
		if minZeros == nil || layerCounts[0] < minZeros[0] {
			minZeros = layerCounts
		}
	}

	log.Printf("The layer with the fewest zeros has n1 * n2 = %d", minZeros[1]*minZeros[2])
}
