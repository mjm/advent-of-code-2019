package day18

import (
	"strings"

	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// Map represents an underground cave map, which is restricted to a 2D
// Cartesian coordinate system and is full of keys and doors.
type Map struct {
	starts        map[point.Point2D]uint32
	points        map[point.Point2D]rune
	keys          map[point.Point2D]rune
	neighbors     map[point.Point2D][]point.Point2D
	paths         map[point.Point2D]map[point.Point2D]Path
	cachedResults map[cacheKey]int
}

type cacheKey struct {
	point point.Point2D
	keys  uint32
}

// MapFromString reads a map from an ASCII representation of it.
func MapFromString(s string) *Map {
	starts := make(map[point.Point2D]uint32)
	points := make(map[point.Point2D]rune)
	keys := make(map[point.Point2D]rune)
	for y, line := range strings.Split(s, "\n") {
		for x, c := range line {
			p := point.Point2D{X: x, Y: y}
			if c != '#' {
				points[p] = c
			}
			if c >= 'a' && c <= 'z' {
				keys[p] = c
			}
			if c == '@' {
				starts[p] = 0
			}
		}
	}

	neighbors := make(map[point.Point2D][]point.Point2D)
	for p := range points {
		for _, q := range p.CardinalNeighbors() {
			if _, ok := points[q]; ok {
				neighbors[p] = append(neighbors[p], q)
			}
		}
	}

	return &Map{
		starts:    starts,
		points:    points,
		keys:      keys,
		neighbors: neighbors,
	}
}

type pointPair struct {
	a point.Point2D
	b point.Point2D
}

func (m *Map) buildPaths() {
	m.paths = make(map[point.Point2D]map[point.Point2D]Path)
	for p := range m.starts {
		keys := m.buildPathsFrom(p)
		m.starts[p] = keys
	}
	for k := range m.keys {
		m.buildPathsFrom(k)
	}
}

func (m *Map) buildPathsFrom(start point.Point2D) uint32 {
	paths := make(map[point.Point2D]Path)
	m.paths[start] = paths

	q := NewPointQueue(128)
	parents := make(map[point.Point2D]point.Point2D)
	q.Enqueue(start)

	var keys uint32
	for !q.Empty() {
		p, _ := q.Dequeue()
		if c, ok := m.keys[p]; ok {
			i := uint32(c - 97)
			keys |= 1 << i
			if p != start {
				paths[p] = m.buildPath(start, p, parents)
			}
		}
		for _, n := range m.neighbors[p] {
			if _, ok := parents[n]; !ok {
				parents[n] = p
				q.Enqueue(n)
			}
		}
	}
	return keys
}

func (m *Map) buildPath(start, dest point.Point2D, parents map[point.Point2D]point.Point2D) Path {
	// traverse back through parent points to rebuild the path taken to get
	// from start to dest. only keep the total distance and record of any
	// doors along the path, which will need keys to open later.

	var path Path
	next := dest
	var d int
	for next != start {
		next = parents[next]
		d++
		c := m.points[next]
		if c >= 'A' && c <= 'Z' {
			i := uint32(c - 65)
			path.KeysNeeded |= 1 << i
		}
	}
	path.Distance = d
	return path
}

// ShortestWalk finds the shortest walk of the map that collects every key
// while only going through a door once it has the corresponding key.
func (m *Map) ShortestWalk() int {
	m.buildPaths()
	m.cachedResults = make(map[cacheKey]int)

	var total int
	for p, keys := range m.starts {
		total += m.shortestWalk(p, keys)
	}
	return total
}

func (m *Map) shortestWalk(start point.Point2D, keys uint32) int {
	if keys == 0 {
		// got all the keys
		return 0
	}

	// check for a previously cached result
	ck := cacheKey{start, keys}
	if result, ok := m.cachedResults[ck]; ok {
		return result
	}

	minWalk := -1

	for p, path := range m.paths[start] {
		if p == start {
			continue
		}
		c := m.keys[p]
		i := uint32(c - 97)
		if keys&(1<<i) == 0 {
			// we already have this key, don't need to go down the path
			continue
		}
		if !path.CanVisit(^keys) {
			// we don't have the right key for the doors needed to get to this
			// key, so ignore it for now.
			continue
		}

		if minWalk != -1 && path.Distance > minWalk {
			// short-circuit: don't recurse and do a bunch of work if the
			// distance just to get this key is longer than the shortest walk
			// we've found in a previous iteration.
			continue
		}

		newKeys := keys &^ (1 << i)
		remaining := m.shortestWalk(p, newKeys)
		if remaining == -1 {
			// it's a dead end, so skip this point.
			// this could cause a cascade of dead ends up the chain.
			continue
		}

		d := path.Distance + remaining
		if minWalk == -1 || d < minWalk {
			minWalk = d
		}
	}

	m.cachedResults[ck] = minWalk
	return minWalk
}
