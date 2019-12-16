package main

import (
	"log"
	"os"

	"github.com/mjm/advent-of-code-2019/day11"
	"github.com/mjm/advent-of-code-2019/pkg/input"
	"github.com/mjm/advent-of-code-2019/pkg/intcode"
	"github.com/mjm/advent-of-code-2019/pkg/point"
)

func main() {
	paintVM, err := intcode.LoadFromString(input.ReadString())
	if err != nil {
		log.Fatal(err)
	}

	part1(paintVM)
	part2(paintVM)
}

func part1(template *intcode.VM) {
	paintVM := template.Clone()

	canvas := day11.NewCanvas()
	robot := day11.NewRobot(canvas)

	paintVM.SetInputFunc(robot.CurrentColor)
	go func() {
		if err := paintVM.Execute(); err != nil {
			log.Fatal(err)
		}
	}()
	robot.Run(paintVM.Output)

	log.Printf("The number of painted squares is %d", canvas.Count())
}

func part2(template *intcode.VM) {
	paintVM := template.Clone()

	canvas := day11.NewCanvas()
	canvas.Paint(point.Point2D{}, 1)
	robot := day11.NewRobot(canvas)

	paintVM.SetInputFunc(robot.CurrentColor)
	go func() {
		if err := paintVM.Execute(); err != nil {
			log.Fatal(err)
		}
	}()
	robot.Run(paintVM.Output)

	canvas.PrintTo(os.Stdout)
}
