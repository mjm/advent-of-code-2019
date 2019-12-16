package day14

import "fmt"

// Quantity represents a specific amount of a particular material
type Quantity struct {
	Material string
	Amount   int
}

// OneFuel is a quantity for a single unit of fuel.
var OneFuel = Quantity{Material: "FUEL", Amount: 1}

var _ fmt.Stringer = (Quantity{})

func (q Quantity) String() string {
	return fmt.Sprintf("%d %s", q.Amount, q.Material)
}
