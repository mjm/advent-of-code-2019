package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gdamore/tcell"
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

	go func() {
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		os.Exit(1)
	}()

	// part1(paintVM)
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
	r := bufio.NewReader(os.Stdin)
	r.ReadRune()
}

func part2(template *intcode.VM) {
	paintVM := template.Clone()
	paintVM.Set(0, 2) // cheating!

	canvas := day13.NewCanvas()
	game := day13.NewGame(canvas)

	autoPlayer := day13.NewAutoPlayer(game)
	paintVM.SetInputFunc(autoPlayer.HandleInput)

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	func() {
		screen.Init()
		defer screen.Fini()

		game.Connect(screen)

		go func() {
			if err := paintVM.Execute(); err != nil {
				log.Fatal(err)
			}
		}()
		game.Run(paintVM.Output)
	}()

	log.Printf("Your final score was %d", game.Score())
}
