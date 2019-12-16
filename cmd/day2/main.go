package main

import (
	"log"

	"github.com/mjm/advent-of-code-2019/day2"
	"github.com/mjm/advent-of-code-2019/pkg/input"
)

const maxVal = 100

func main() {
	vm, err := day2.LoadFromString(input.ReadString())
	if err != nil {
		log.Fatal(err)
	}

	cloned := vm.Clone()

	cloned.Set(1, 12)
	cloned.Set(2, 2)

	if err := cloned.Execute(); err != nil {
		log.Fatal(err)
	}

	log.Printf("The value at position 0 is %d", cloned.At(0))

	for n := 0; n < maxVal; n++ {
		for v := 0; v < maxVal; v++ {
			cloned = vm.Clone()
			cloned.Set(1, n)
			cloned.Set(2, v)
			if err := cloned.Execute(); err != nil {
				log.Fatal(err)
			}

			if cloned.At(0) == 19690720 {
				log.Printf("n = %d\n", n)
				log.Printf("v = %d\n", v)
				log.Printf("100 * n + v = %d\n", 100*n+v)
				return
			}
		}
	}

	log.Fatalln("Could not find correct noun and verb, try a higher max value.")
}
