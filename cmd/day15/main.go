package main

import (
	"fmt"
	"log"

	"github.com/mjm/advent-of-code-2019/day15"
	"github.com/mjm/advent-of-code-2019/pkg/input"
	"github.com/mjm/advent-of-code-2019/pkg/intcode"
)

func main() {
	vm, err := intcode.LoadFromString(input.ReadString())
	if err != nil {
		log.Fatal(err)
	}

	input := make(chan int)
	vm.SetInputChan(input)

	canvas := day15.NewCanvas()
	mapper := day15.NewMapper(canvas)

	go func() {
		if err := vm.Execute(); err != nil {
			log.Fatal(err)
		}
	}()
	dest := mapper.Start(input, vm.Output)

	canvas.Draw(func(x, y, value int) {
		if x == 0 {
			fmt.Println()
		}
		tile := day15.Tile(value)

		var c rune
		switch tile {
		case day15.TileUnknown:
			c = ' '
		case day15.TilePassable:
			c = '.'
		case day15.TileWall:
			c = '#'
		case day15.TileDestination:
			c = 'O'
		case day15.TileStart:
			c = 'S'
		}

		fmt.Printf("%c", c)
	})

	pf := day15.NewPathFinder(canvas, dest)
	path := pf.ShortestPath()
	log.Printf("Shortest path to destination is %d steps:\n%v", len(path)-1, path)

	steps := day15.Fill(canvas)
	log.Printf("It took %d minutes to fill the map.", steps)
}
