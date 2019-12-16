package day14

import (
	"fmt"
	"math"
	"strings"
)

// Table is a list of conversion rules between types of materials.
type Table struct {
	rules       []*Rule
	byOutput    map[string]*Rule
	sortedRules []*Rule
}

// TableFromString creates a conversion table from lines of rules.
func TableFromString(s string) (*Table, error) {
	lines := strings.Split(s, "\n")

	rules := make([]*Rule, 0, len(lines))
	byOutput := make(map[string]*Rule)
	for _, line := range lines {
		r, err := RuleFromString(line)
		if err != nil {
			return nil, err
		}

		rules = append(rules, &r)
		byOutput[r.Output.Material] = &r
	}

	t := &Table{
		rules:    rules,
		byOutput: byOutput,
	}
	t.sortRules()
	return t, nil
}

func (t *Table) sortRules() {
	marks := make(map[string]int)
	sorted := make([]*Rule, 0, len(t.rules))

	var visit func(string)
	visit = func(material string) {
		mark := marks[material]
		if mark == 2 {
			return
		}
		if mark == 1 {
			panic("rule table is not a DAG")
		}

		marks[material] = 1
		r := t.byOutput[material]
		if r == nil {
			// This must be ORE
			marks[material] = 2
			return
		}

		for _, input := range r.Inputs {
			visit(input.Material)
		}
		marks[material] = 2
		sorted = append(sorted, r)
	}

	for {
		var visited bool
		for _, r := range t.rules {
			mark := marks[r.Output.Material]
			if mark == 0 {
				visit(r.Output.Material)
				visited = true
				break
			}
		}
		if !visited {
			break
		}
	}

	// reverse the list
	for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
		sorted[i], sorted[j] = sorted[j], sorted[i]
	}

	t.sortedRules = sorted
}

// RequiredOre returns the minimum amount of ore needed to produce the desired
// quantity of material with the rules in this table.
func (t *Table) RequiredOre(desired Quantity) int {
	if desired.Material == "ORE" {
		return desired.Amount
	}

	reqs := make(map[string]int)
	reqs[desired.Material] = desired.Amount

	// apply rules in topo sort order, figuring out how much of each input we need
	// for the output of that rule.
	for _, r := range t.sortedRules {
		amount := reqs[r.Output.Material]
		multiplier := int(math.Ceil(float64(amount) / float64(r.Output.Amount)))
		for _, q := range r.Inputs {
			reqs[q.Material] += multiplier * q.Amount
		}

		// clear out the requirement now that it's been satisfied
		reqs[r.Output.Material] = 0
	}

	return reqs["ORE"]
}

var _ fmt.Stringer = (*Table)(nil)

func (t *Table) String() string {
	var b strings.Builder
	for i, r := range t.rules {
		if i > 0 {
			b.WriteRune('\n')
		}
		b.WriteString(r.String())
	}
	return b.String()
}
