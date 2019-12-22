package day22

import (
	"math/big"
)

// ShuffleFind shuffles position i using the given techniques, and returns the
// position in the deck where that card ended up after shuffling.
func ShuffleFind(ts []Technique, i, size int64) int64 {
	t := ComposeTechniques(ts, size)
	return Perform(t, i, size)
}

// ShuffleN shuffles using the given techniques for several times in a row, and
// returns the card at the ith position.
func ShuffleN(ts []Technique, i, size, times int64) int64 {
	t := ComposeTechniques(ts, size)
	t = RepeatTechnique(t, size, times)

	numer := (i - t.Constant()) % size
	if numer < 0 {
		numer += size
	}
	denom := t.Multiplier()
	if denom < 0 {
		denom += size
	}
	a1 := new(big.Int).ModInverse(big.NewInt(denom), big.NewInt(size))
	a1.Mul(a1, big.NewInt(numer))
	a1.Mod(a1, big.NewInt(size))

	return a1.Int64()
}
