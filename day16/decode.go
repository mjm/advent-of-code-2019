package day16

import (
	"strconv"
)

// Decode decodes a signal using the FFT.
func Decode(signal []int) []int {
	input := repeat(signal, 10000)
	first7, err := strconv.Atoi(DigitsToString(input[:7]))
	if err != nil {
		panic(err)
	}
	output := OffsetFFT(input, 100, first7)
	return output[:8]
}

func repeat(ns []int, n int) []int {
	res := make([]int, len(ns)*n)
	for i := 0; i < len(res); i += len(ns) {
		copy(res[i:i+len(ns)], ns)
	}
	return res
}
