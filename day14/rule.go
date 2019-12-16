package day14

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// Rule defines how an output can be constructed from some input materials.
type Rule struct {
	Inputs []Quantity
	Output Quantity
}

// RuleFromString creates a new rule from its string representation.
func RuleFromString(s string) (Rule, error) {
	var r Rule

	var scan scanner.Scanner
	scan.Init(strings.NewReader(s))
	scan.Mode = scanner.ScanInts | scanner.ScanIdents

	var isOutput bool
	var q Quantity
	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		switch tok {
		case '=':
			tok = scan.Scan()
			if tok == '>' {
				isOutput = true
			} else {
				return r, fmt.Errorf("invalid rule string")
			}

		case scanner.Int:
			amount, err := strconv.Atoi(scan.TokenText())
			if err != nil {
				return r, fmt.Errorf("could not parse amount: %w", err)
			}
			q.Amount = amount

		case scanner.Ident:
			q.Material = scan.TokenText()
			if isOutput {
				r.Output = q
			} else {
				r.Inputs = append(r.Inputs, q)
			}
			q = Quantity{}
		}
	}

	return r, nil
}

var _ fmt.Stringer = (Rule{})

func (r Rule) String() string {
	var b strings.Builder
	for i, q := range r.Inputs {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(q.String())
	}
	b.WriteString(" => ")
	b.WriteString(r.Output.String())
	return b.String()
}
