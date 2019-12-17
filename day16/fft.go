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
		pat := pattern(i, len(input))
		var n int
		for j, m := range input {
			n += pat[j] * m
		}
		output[i] = abs(n) % 10
	}
	return output
}

var patternBase = []int{0, 1, 0, -1}

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
