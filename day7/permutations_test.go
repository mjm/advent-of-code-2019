package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPermutations(t *testing.T) {
	numChan := AllPermutations([]int{1, 2, 3})
	var nums [][]int
	for n := range numChan {
		nums = append(nums, n)
	}

	assert.Equal(t, [][]int{
		{1, 2, 3},
		{2, 1, 3},
		{3, 1, 2},
		{1, 3, 2},
		{2, 3, 1},
		{3, 2, 1},
	}, nums)
}
