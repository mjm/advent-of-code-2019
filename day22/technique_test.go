package day22

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTechnique(t *testing.T) {
	tech, err := ParseTechnique("deal into new stack")
	assert.NoError(t, err)
	assert.Equal(t, &newStackTechnique{}, tech)

	tech, err = ParseTechnique("cut 6")
	assert.NoError(t, err)
	assert.Equal(t, &cutTechnique{n: 6}, tech)

	tech, err = ParseTechnique("cut -8")
	assert.NoError(t, err)
	assert.Equal(t, &cutTechnique{n: -8}, tech)

	tech, err = ParseTechnique("deal with increment 7")
	assert.NoError(t, err)
	assert.Equal(t, &dealTechnique{n: 7}, tech)

	tech, err = ParseTechnique("foo bar baz")
	assert.EqualError(t, err, `no technique pattern matched "foo bar baz"`)
	assert.Nil(t, tech)
}

func TestNewStackTechnique(t *testing.T) {
	tech := &newStackTechnique{}

	assert.Equal(t, int64(3), Perform(tech, 1, 5))
}

func TestCutTechniquePositive(t *testing.T) {
	tech := &cutTechnique{n: 3}

	assert.Equal(t, int64(6), Perform(tech, 1, 8))
	assert.Equal(t, int64(1), Perform(tech, 4, 8))
}

func TestCutTechniqueNegative(t *testing.T) {
	tech := &cutTechnique{n: -3}

	assert.Equal(t, int64(6), Perform(tech, 3, 8))
	assert.Equal(t, int64(1), Perform(tech, 6, 8))
}

func TestDealTechnique(t *testing.T) {
	tech := &dealTechnique{n: 3}

	assert.Equal(t, int64(0), Perform(tech, 0, 10))
	assert.Equal(t, int64(1), Perform(tech, 7, 10))
	assert.Equal(t, int64(3), Perform(tech, 1, 10))
	assert.Equal(t, int64(6), Perform(tech, 2, 10))
}

func TestComposeTechniques(t *testing.T) {
	ts, err := ParseAllTechniques(`cut 6
deal with increment 7
deal into new stack`)
	assert.NoError(t, err)

	tech := ComposeTechniques(ts, 10)
	assert.Equal(t, int64(1), Perform(tech, 0, 10))
}

func TestRepeatTechnique(t *testing.T) {
	ts, err := ParseAllTechniques(`cut 3
deal with increment 2
deal into new stack`)
	assert.NoError(t, err)

	tech := ComposeTechniques(ts, 7)
	log.Print(tech)
	tech = RepeatTechnique(tech, 7, 3)

	assert.Equal(t, int64(1), Perform(tech, 0, 7))
}

func TestShuffleN(t *testing.T) {
	ts, err := ParseAllTechniques(`cut 3
deal with increment 2
deal into new stack`)
	assert.NoError(t, err)

	assert.Equal(t, int64(0), ShuffleN(ts, 1, 7, 3))
	assert.Equal(t, int64(1), ShuffleN(ts, 0, 7, 3))
}
