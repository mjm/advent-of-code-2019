package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/mjm/advent-of-code-2019/day4"
)

func main() {
	flag.Parse()
	min, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	passes := day4.ValidPasswords(min, max)
	log.Printf("Number of valid passwords: %d", len(passes))

	passes = day4.ReallyValidPasswords(min, max)
	log.Printf("Number of really valid passwords: %d", len(passes))
}
