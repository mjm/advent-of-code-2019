package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/mjm/advent-of-code-2019/day13"
	"github.com/mjm/advent-of-code-2019/pkg/intcode"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	paintVM, err := intcode.LoadFromString(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	part1(paintVM)
	part2(paintVM)
}

func part1(template *intcode.VM) {
	paintVM := template.Clone()

	canvas := day13.NewCanvas()
	game := day13.NewGame(canvas)

	go func() {
		if err := paintVM.Execute(); err != nil {
			log.Fatal(err)
		}
	}()
	game.Run(paintVM.Output)

	log.Printf("The number of block tiles is %d", canvas.CountColor(2))
}

func part2(template *intcode.VM) {

}
