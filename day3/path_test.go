package day3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegmentFromString(t *testing.T) {
	cases := []struct {
		s         string
		direction Direction
		length    int
		err       string
	}{
		{"R1004", Right, 1004, ""},
		{"D53", Down, 53, ""},
		{"L10", Left, 10, ""},
		{"U126", Up, 126, ""},
		{"R130", Right, 130, ""},
		{"U533", Up, 533, ""},

		{
			s:   "",
			err: "cannot parse segment from empty string",
		},
		{
			s:   "E324",
			err: "unknown direction character: 'E'",
		},
		{
			s:   "RASDF",
			err: "cannot decode path segment length: strconv.Atoi: parsing \"ASDF\": invalid syntax",
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("parses %s", c.s), func(t *testing.T) {
			seg, err := SegmentFromString(c.s)
			if c.err == "" {
				assert.NoError(t, err)
				assert.Equal(t, c.direction, seg.Direction)
				assert.Equal(t, c.length, seg.Length)
			} else {
				assert.EqualError(t, err, c.err)
			}
		})
	}
}

func TestPathFromString(t *testing.T) {
	p, err := PathFromString("R75,D30,R83,U83,L12")
	assert.NoError(t, err)
	assert.Equal(t, []PathSegment{
		{Right, 75},
		{Down, 30},
		{Right, 83},
		{Up, 83},
		{Left, 12},
	}, p.Segments)
}
