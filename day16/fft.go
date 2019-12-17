package day16

import (
	"fmt"
	"strings"
)

// Digits takes a string of all digit characters and returns a slice
// of all of the individual digits as integers.
func Digits(s string) []int {
	ns := make([]int, 0, len(s))
	for _, c := range s {
		ns = append(ns, int(c-'0'))
	}
	return ns
}

// DigitsToString converts a list of digits to a string concatenating
// them all together.
func DigitsToString(ns []int) string {
	var b strings.Builder
	for _, n := range ns {
		fmt.Fprintf(&b, "%d", n)
	}
	return b.String()
}

// FFT runs a Flawed Frequency Transmission algorithm over the input
// n times.
func FFT(input []int, n int) []int {
	ns := input
	for i := 0; i < n; i++ {
		ns = fft(ns)
	}
	return ns
}

func fft(input []int) []int {
	output := make([]int, len(input))
	for i := range output {
		var n int
		for j := i; j < len(input); j++ {
			n += patternValue(i, j) * input[j]
		}
		output[i] = abs(n) % 10
	}
	return output
}

// OffsetFFT calculates the FFT offset into some point in the second half of
// the input. This is a significantly optimized version, but it can only work
// on the second half. It returns only the section of the result after the
// given offset (the rest of the sequence is ignored entirely).
func OffsetFFT(input []int, n int, offset int) []int {
	if offset < len(input)/2 {
		panic("cannot use OffsetFFT if offset is in the first half of the input")
	}

	ns := input[offset:]
	for i := 0; i < n; i++ {
		ns = offsetFFT(ns, offset)
	}
	return ns
}

func offsetFFT(input []int, offset int) []int {
	output := make([]int, len(input))
	output[len(output)-1] = input[len(input)-1]
	for i := len(input) - 2; i >= 0; i-- {
		output[i] = (output[i+1] + input[i]) % 10
	}
	return output
}

var patternBase = []int{0, 1, 0, -1}

func patternValue(i, j int) int {
	i++
	j++

	k := (j / i) % 4
	return patternBase[k]
}

func pattern(i, n int) []int {
	res := make([]int, 0, n+1)
	for j, k := 0, 0; j < n+1; j, k = j+i+1, (k+1)%len(patternBase) {
		for l := 0; l < (i+1) && j+l < n+1; l++ {
			res = append(res, patternBase[k])
		}
	}
	return res[1:]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
