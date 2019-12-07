package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"text/scanner"

	"github.com/mjm/advent-of-code-2019/day1"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	var s scanner.Scanner
	s.Init(file)

	var fuel, total int
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		mass, err := strconv.Atoi(s.TokenText())
		if err != nil {
			log.Fatal(err)
		}

		m := day1.NewModule(mass)
		fuel += day1.FuelRequired(m)
		total += day1.TotalFuelRequired(m)
	}

	log.Printf("Base required fuel: %d\n", fuel)
	log.Printf("Total required fuel: %d\n", total)
}
