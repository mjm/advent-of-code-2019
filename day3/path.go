package day3

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errEmptySegment = errors.New("cannot parse segment from empty string")
)

type Path struct {
	Segments []PathSegment
}

func PathFromString(s string) (Path, error) {
	p := Path{}
	for _, str := range strings.Split(s, ",") {
		seg, err := SegmentFromString(str)
		if err != nil {
			return Path{}, err
		}

		p.Segments = append(p.Segments, seg)
	}
	return p, nil
}

func (p Path) Apply(m *Map, id int8) {
	var x, y int
	for _, seg := range p.Segments {
		x, y = seg.Apply(m, x, y, id)
	}
}

type Direction point

var (
	Up    = Direction{0, 1}
	Down  = Direction{0, -1}
	Left  = Direction{-1, 0}
	Right = Direction{1, 0}

	directions = map[rune]Direction{
		'U': Up,
		'D': Down,
		'L': Left,
		'R': Right,
	}
)

type PathSegment struct {
	Direction Direction
	Length    int
}

func SegmentFromString(s string) (PathSegment, error) {
	seg := PathSegment{}

	if len(s) == 0 {
		return seg, errEmptySegment
	}

	c := rune(s[0])
	dir, ok := directions[c]
	if !ok {
		return seg, fmt.Errorf("unknown direction character: %q", c)
	}

	seg.Direction = dir

	l, err := strconv.Atoi(s[1:])
	if err != nil {
		return seg, fmt.Errorf("cannot decode path segment length: %w", err)
	}

	seg.Length = l
	return seg, nil
}

func (s PathSegment) Apply(m *Map, x, y int, id int8) (int, int) {
	for i := 0; i < s.Length; i++ {
		x += s.Direction.X
		y += s.Direction.Y

		m.Set(x, y, id)
	}

	return x, y
}
