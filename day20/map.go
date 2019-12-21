package day20

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// Map represents a map of a Plutonian maze.
type Map struct {
	start        point.Point2D
	end          point.Point2D
	points       map[point.Point2D]interface{}
	neighbors    map[point.Point2D][]point.Point2D
	portalOffset map[pointPair]int
}

type pointPair struct {
	a point.Point2D
	b point.Point2D
}

// MapFromString reads a map from an ASCII representation of it.
func MapFromString(s string) *Map {
	var start, end point.Point2D
	points := make(map[point.Point2D]interface{})
	neighbors := make(map[point.Point2D][]point.Point2D)
	portals := make(map[string]point.Point2D)
	portalOffsets := make(map[pointPair]int)

	lines := strings.Split(s, "\n")
	for y, line := range lines {
		for x, c := range line {
			p := point.Point2D{X: x, Y: y}
			if c == '.' {
				points[p] = c

				var s string
				var isOuter bool
				if above := rune(lines[y-1][x]); unicode.IsLetter(above) {
					s = fmt.Sprintf("%c%c", lines[y-2][x], lines[y-1][x])
					isOuter = y-2 == 0
				} else if left := rune(line[x-1]); unicode.IsLetter(left) {
					s = fmt.Sprintf("%c%c", line[x-2], line[x-1])
					isOuter = x-2 == 0
				} else if right := rune(line[x+1]); unicode.IsLetter(right) {
					s = fmt.Sprintf("%c%c", line[x+1], line[x+2])
					isOuter = x+3 == len(line)
				} else if below := rune(lines[y+1][x]); unicode.IsLetter(below) {
					s = fmt.Sprintf("%c%c", lines[y+1][x], lines[y+2][x])
					isOuter = y+3 == len(lines)
				}

				if s != "" {
					if s == "AA" {
						start = p
					} else if s == "ZZ" {
						end = p
					} else if other, ok := portals[s]; ok {
						neighbors[p] = append(neighbors[p], other)
						neighbors[other] = append(neighbors[other], p)

						if isOuter {
							portalOffsets[pointPair{p, other}] = -1
							portalOffsets[pointPair{other, p}] = 1
						} else {
							portalOffsets[pointPair{p, other}] = 1
							portalOffsets[pointPair{other, p}] = -1
						}
					} else {
						portals[s] = p
					}
				}
			}
		}
	}

	for p := range points {
		for _, q := range p.CardinalNeighbors() {
			if _, ok := points[q]; ok {
				neighbors[p] = append(neighbors[p], q)
			}
		}
	}

	return &Map{
		start:        start,
		end:          end,
		points:       points,
		neighbors:    neighbors,
		portalOffset: portalOffsets,
	}
}

// ShortestPath finds the shortest path from the start to the end of the map.
// Returns the list of points that form the path, starting from the goal and
// ending with the start point, including both.
func (m *Map) ShortestPath() []point.Point2D {
	q := NewQueue(128)
	parents := make(map[point.Point2D]point.Point2D)
	q.Enqueue(m.start)

	for !q.Empty() {
		el, _ := q.Dequeue()
		p := el.(point.Point2D)
		if p == m.end {
			n := p
			path := []point.Point2D{n}
			for {
				var ok bool
				n, ok = parents[n]
				if !ok {
					return path
				}

				path = append(path, n)
			}
		}

		for _, n := range m.neighbors[p] {
			if n == m.start {
				continue
			}

			if _, ok := parents[n]; !ok {
				parents[n] = p
				q.Enqueue(n)
			}
		}
	}

	return nil
}

type levelPoint struct {
	point point.Point2D
	level int
}

func (m *Map) ShortestPathLevels() []point.Point2D {
	q := NewQueue(128)
	parents := make(map[levelPoint]levelPoint)
	q.Enqueue(levelPoint{m.start, 0})

	for !q.Empty() {
		el, _ := q.Dequeue()
		lp := el.(levelPoint)
		p, level := lp.point, lp.level
		if p == m.end && level == 0 {
			nlp := lp
			path := []point.Point2D{p}
			for {
				var ok bool
				nlp, ok = parents[nlp]
				if !ok {
					return path
				}

				path = append(path, nlp.point)
			}
		}

		for _, n := range m.neighbors[p] {
			if n == m.start {
				continue
			}

			nlp := levelPoint{point: n, level: level}
			if offset, ok := m.portalOffset[pointPair{p, n}]; ok {
				if level == 0 && offset < 0 {
					continue
				}

				nlp.level += offset
			}

			if _, ok := parents[nlp]; !ok {
				parents[nlp] = lp
				q.Enqueue(nlp)
			}
		}
	}

	return nil
}
