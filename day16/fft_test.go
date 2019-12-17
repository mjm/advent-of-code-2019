package day16

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPattern(t *testing.T) {
	cases := []struct {
		i   int
		n   int
		res []int
	}{
		{0, 6, []int{1, 0, -1, 0, 1, 0}},
		{1, 6, []int{0, 1, 1, 0, 0, -1}},
		{1, 10, []int{0, 1, 1, 0, 0, -1, -1, 0, 0, 1}},
		{2, 8, []int{0, 0, 1, 1, 1, 0, 0, 0}},
		{3, 10, []int{0, 0, 0, 1, 1, 1, 1, 0, 0, 0}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("pattern where i=%d n=%d", c.i, c.n), func(t *testing.T) {
			assert.Equal(t, c.res, pattern(c.i, c.n))
			for i, n := range c.res {
				assert.Equal(t, n, patternValue(c.i, i), "expected patternValue(%d, %d) to be %d", c.i, i, n)
			}
		})
	}
}

func TestDigits(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, Digits("123"))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, Digits("12345678"))
}

func TestFFT(t *testing.T) {
	seq := Digits("12345678")
	seq = FFT(seq, 1)
	assert.Equal(t, Digits("48226158"), seq)
	seq = FFT(seq, 3)
	assert.Equal(t, Digits("01029498"), seq)

	seq = FFT(Digits("80871224585914546619083218645595"), 100)
	assert.Equal(t, Digits("24176176"), seq[:8])
	seq = FFT(Digits("19617804207202209144916044189917"), 100)
	assert.Equal(t, Digits("73745418"), seq[:8])
	seq = FFT(Digits("69317163492948606335995924319873"), 100)
	assert.Equal(t, Digits("52432133"), seq[:8])
}
