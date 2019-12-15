package mathalg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCM(t *testing.T) {
	assert.Equal(t, 6, LCM(2, 3))
	assert.Equal(t, 15, LCM(5, 15))
	assert.Equal(t, 150, LCM(30, 75))
	assert.Equal(t, 90, LCM(30, 90))
	assert.Equal(t, 120, LCM(30, 40))
	assert.Equal(t, 30, LCM(2, 3, 5))
	assert.Equal(t, 105, LCM(5, 15, 7))
}
