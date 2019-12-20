package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/mjm/advent-of-code-2019/day18"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	m := day18.MapFromString(input.ReadString())

	if *cpuprofile != "" {
		log.Println(*cpuprofile)
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	done := make(chan int)
	go func() {
		done <- m.ShortestWalk()
	}()

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt)
	select {
	case result := <-done:
		log.Printf("Shortest path is %d", result)
	case <-signalCh:
		log.Print("Exiting.")
	}
}
