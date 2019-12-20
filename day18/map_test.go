package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const ex1 = `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`

const ex2 = `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`

const ex3 = `#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`

const ex4 = `########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`

func TestMapFromString(t *testing.T) {
	m1 := MapFromString(ex1)
	assert.Equal(t, 45, len(m1.points))
	assert.Equal(t, 6, len(m1.keys))

	m2 := MapFromString(ex2)
	assert.Equal(t, 45, len(m2.points))
	assert.Equal(t, 7, len(m2.keys))

	m3 := MapFromString(ex3)
	assert.Equal(t, 63, len(m3.points))
	assert.Equal(t, 16, len(m3.keys))

	m4 := MapFromString(ex4)
	assert.Equal(t, 31, len(m4.points))
	assert.Equal(t, 9, len(m4.keys))
}

func TestShortestWalk(t *testing.T) {
	m1 := MapFromString(ex1)
	assert.Equal(t, 86, m1.ShortestWalk())

	m2 := MapFromString(ex2)
	assert.Equal(t, 132, m2.ShortestWalk())

	m3 := MapFromString(ex3)
	assert.Equal(t, 136, m3.ShortestWalk())

	m4 := MapFromString(ex4)
	assert.Equal(t, 81, m4.ShortestWalk())
}
