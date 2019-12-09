package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageFromString(t *testing.T) {
	img, err := ImageFromString(`123456789012`, 3, 2)
	assert.NoError(t, err)

	assert.Equal(t, 3, img.Width)
	assert.Equal(t, 2, img.Height)
	assert.Equal(t, [][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 0, 1, 2}}, img.Layers)
}

func TestLayerCounts(t *testing.T) {
	img, err := ImageFromString(`112333152435`, 3, 2)
	assert.NoError(t, err)

	assert.Equal(t, []map[int]int{
		{1: 2, 2: 1, 3: 3},
		{1: 1, 5: 2, 2: 1, 3: 1, 4: 1},
	}, img.DigitCounts())
}
