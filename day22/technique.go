package day22

import (
	"fmt"
	"math/big"
	"strings"
)

// Technique is an interface for methods of shuffling a deck of cards.
type Technique interface {
	Multiplier() int64
	Constant() int64
}

// ParseTechnique reads a technique from a string description.
func ParseTechnique(s string) (Technique, error) {
	if s == "deal into new stack" {
		return &newStackTechnique{}, nil
	}

	var n int
	if _, err := fmt.Sscanf(s, "cut %d", &n); err == nil {
		return &cutTechnique{n}, nil
	}

	if _, err := fmt.Sscanf(s, "deal with increment %d", &n); err == nil {
		return &dealTechnique{n}, nil
	}

	return nil, fmt.Errorf("no technique pattern matched %q", s)
}

// ParseAllTechniques reads a technique from each line of the given string.
func ParseAllTechniques(s string) ([]Technique, error) {
	var ts []Technique
	for i, line := range strings.Split(s, "\n") {
		t, err := ParseTechnique(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing technique %d: %w", i, err)
		}

		ts = append(ts, t)
	}
	return ts, nil
}

// ComposeTechniques creates a new shuffling technique that has the same result
// as performing the given sequence of techniques.
func ComposeTechniques(ts []Technique, size int64) Technique {
	a := ts[0].Multiplier()
	b := ts[0].Constant()

	for _, t := range ts[1:] {
		a, b = (a*t.Multiplier())%size, (b*t.Multiplier()+t.Constant())%size
		if a < 0 {
			a += size
		}
		if b < 0 {
			b += size
		}
	}

	return &composedTechnique{a, b}
}

// RepeatTechnique creates a new shuffling technique that is the result of
// repeating the given technique a certain number of times.
func RepeatTechnique(t Technique, size, n int64) Technique {
	m := big.NewInt(n)
	s := big.NewInt(size)
	a := big.NewInt(t.Multiplier())
	a.Exp(a, m, s)

	b := big.NewInt(t.Constant())
	one := big.NewInt(1)
	numer := new(big.Int).Sub(a, one)
	denom := new(big.Int).Sub(big.NewInt(t.Multiplier()), one)
	denom.ModInverse(denom, s)
	b.Mul(b, new(big.Int).Mul(numer, denom))
	b.Mod(b, s)

	return &composedTechnique{a.Int64(), b.Int64()}
}

// Perform applies a shuffling technique to a particular position in a deck.
func Perform(t Technique, n, size int64) int64 {
	var result big.Int
	result.Mul(big.NewInt(t.Multiplier()), big.NewInt(n))
	result.Add(&result, big.NewInt(t.Constant()))
	result.Mod(&result, big.NewInt(size))

	i := result.Int64()
	if i < 0 {
		return size + i
	}
	return i
}

type composedTechnique struct {
	a int64
	b int64
}

func (t *composedTechnique) Multiplier() int64 {
	return t.a
}

func (t *composedTechnique) Constant() int64 {
	return t.b
}

type newStackTechnique struct {
}

func (*newStackTechnique) Multiplier() int64 {
	return -1
}

func (*newStackTechnique) Constant() int64 {
	return -1
}

type cutTechnique struct {
	n int
}

func (*cutTechnique) Multiplier() int64 {
	return 1
}

func (t *cutTechnique) Constant() int64 {
	return -int64(t.n)
}

type dealTechnique struct {
	n int
}

func (t *dealTechnique) Multiplier() int64 {
	return int64(t.n)
}

func (t *dealTechnique) Constant() int64 {
	return 0
}
