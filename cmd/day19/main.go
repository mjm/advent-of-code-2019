package main

import (
	"log"
	"os"

	"github.com/mjm/advent-of-code-2019/day19"
	"github.com/mjm/advent-of-code-2019/pkg/input"
	"github.com/mjm/advent-of-code-2019/pkg/intcode"
)

func main() {
	vm, err := intcode.LoadFromString(input.ReadString())
	if err != nil {
		log.Fatal(err)
	}

	canvas := day19.NewCanvas()
	drone := day19.NewDrone(canvas)

	area := drone.Scan(50, 50, vm)
	log.Printf("The area affected by the tractor beam is %d squares.", area)

	canvas.PrintTo(os.Stdout)

	log.Printf("The point where we might be able to fit Santa's ship is %v", drone.FindSquare(100, vm))
	f, err := os.Create("/tmp/canvas.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	canvas.PrintTo(f)
	// canvas.PrintTo(os.Stdout)
}
