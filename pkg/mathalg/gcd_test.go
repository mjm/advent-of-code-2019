package mathalg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	assert.Equal(t, 21, GCD(252, 105))
	assert.Equal(t, 3, GCD(3, 6))
	assert.Equal(t, 3, GCD(6, 9))
	assert.Equal(t, 1, GCD(2, 3, 5))
	assert.Equal(t, 5, GCD(5, 10, 15))
}
