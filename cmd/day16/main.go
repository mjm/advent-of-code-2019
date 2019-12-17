package main

import (
	"log"

	"github.com/mjm/advent-of-code-2019/day16"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

func main() {
	digits := day16.Digits(input.ReadString())
	result := day16.FFT(digits, 100)
	log.Printf("Applying FFT 100 times gives the following first 8 digits: %s", day16.DigitsToString(result[:8]))
	log.Printf("Decoding the signal gives %s", day16.DigitsToString(day16.Decode(digits)))
}
