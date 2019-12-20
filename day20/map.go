package day20

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// Map represents a map of a Plutonian maze.
type Map struct {
	start     point.Point2D
	end       point.Point2D
	points    map[point.Point2D]interface{}
	neighbors map[point.Point2D][]point.Point2D
	fs        map[point.Point2D]int
	gs        map[point.Point2D]int
}

// MapFromString reads a map from an ASCII representation of it.
func MapFromString(s string) *Map {
	var start, end point.Point2D
	points := make(map[point.Point2D]interface{})
	neighbors := make(map[point.Point2D][]point.Point2D)
	portals := make(map[string]point.Point2D)

	lines := strings.Split(s, "\n")
	for y, line := range lines {
		for x, c := range line {
			p := point.Point2D{X: x, Y: y}
			if c == '.' {
				points[p] = c

				var s string
				if above := rune(lines[y-1][x]); unicode.IsLetter(above) {
					s = fmt.Sprintf("%c%c", lines[y-2][x], lines[y-1][x])
				} else if left := rune(line[x-1]); unicode.IsLetter(left) {
					s = fmt.Sprintf("%c%c", line[x-2], line[x-1])
				} else if right := rune(line[x+1]); unicode.IsLetter(right) {
					s = fmt.Sprintf("%c%c", line[x+1], line[x+2])
				} else if below := rune(lines[y+1][x]); unicode.IsLetter(below) {
					s = fmt.Sprintf("%c%c", lines[y+1][x], lines[y+2][x])
				}

				if s != "" {
					if s == "AA" {
						start = p
					} else if s == "ZZ" {
						end = p
					} else if other, ok := portals[s]; ok {
						neighbors[p] = append(neighbors[p], other)
						neighbors[other] = append(neighbors[other], p)
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
		start:     start,
		end:       end,
		points:    points,
		neighbors: neighbors,
	}
}

// ShortestPath finds the shortest path from the start to the end of the map.
// Returns the list of points that form the path, starting from the goal and
// ending with the start point, including both.
func (m *Map) ShortestPath() []point.Point2D {
	q := NewPointQueue(128)
	parents := make(map[point.Point2D]point.Point2D)
	q.Enqueue(m.start)

	for !q.Empty() {
		p, _ := q.Dequeue()
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
