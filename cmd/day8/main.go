package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/mjm/advent-of-code-2019/day8"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

func main() {
	img, err := day8.ImageFromString(input.ReadString(), 25, 6)
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

	black := color.New(color.BgBlack)
	white := color.New(color.BgWhite)

	composite := img.Composite()

	for y := 0; y < img.Height; y++ {
		black.Print(" ")
		for x := 0; x < img.Width; x++ {
			val := composite[y*img.Width+x]
			if val == 0 {
				black.Print(" ")
			} else if val == 1 {
				white.Print(" ")
			} else {
				panic("unexpected color")
			}
		}
		black.Print(" ")
		fmt.Println()
	}
}
