package day15

import (
	"math"

	"github.com/mjm/advent-of-code-2019/pkg/point"
)

// PathFinder finds paths between a start and goal point on a map.
type PathFinder struct {
	canvas *Canvas
	start  point.Point2D
	goal   point.Point2D
	fs     map[point.Point2D]int
	gs     map[point.Point2D]int
}

// NewPathFinder creates a new path finder with a given map and goal.
func NewPathFinder(canvas *Canvas, goal point.Point2D) *PathFinder {
	var start point.Point2D

	gs := make(map[point.Point2D]int)
	gs[start] = 0
	fs := make(map[point.Point2D]int)

	pf := &PathFinder{
		canvas: canvas,
		start:  start,
		goal:   goal,
		fs:     fs,
		gs:     gs,
	}
	fs[start] = pf.h(start)
	return pf
}

// ShortestPath finds the shortest path from (0, 0) to the path finder's
// goal. Returns the list of points that form the path, starting from the
// goal and ending with the start point, including both.
func (pf *PathFinder) ShortestPath() []point.Point2D {
	open := make(map[point.Point2D]interface{})
	open[pf.start] = struct{}{}
	prev := make(map[point.Point2D]point.Point2D)

	for len(open) > 0 {
		var current *point.Point2D
		for p := range open {
			if current == nil || pf.f(p) < pf.f(*current) {
				current = &p
			}
		}

		if *current == pf.goal {
			n := *current
			path := []point.Point2D{n}
			for {
				var ok bool
				n, ok = prev[n]
				if !ok {
					return path
				}

				path = append(path, n)
			}
		}

		delete(open, *current)
		for _, n := range pf.neighbors(*current) {
			if g := pf.g(*current) + 1; g < pf.g(n) {
				prev[n] = *current
				pf.gs[n] = g
				pf.fs[n] = g + pf.h(n)
				open[n] = struct{}{}
			}
		}
	}

	return nil
}

// // A* finds a path from start to goal.
// // h is the heuristic function. h(n) estimates the cost to reach goal from node n.
// function A_Star(start, goal, h)
//     // The set of discovered nodes that may need to be (re-)expanded.
//     // Initially, only the start node is known.
//     openSet := {start}

//     // For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start to n currently known.
//     cameFrom := an empty map

//     // For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
//     gScore := map with default value of Infinity
//     gScore[start] := 0

//     // For node n, fScore[n] := gScore[n] + h(n).
//     fScore := map with default value of Infinity
//     fScore[start] := h(start)

//     while openSet is not empty
//         current := the node in openSet having the lowest fScore[] value
//         if current = goal
//             return reconstruct_path(cameFrom, current)

//         openSet.Remove(current)
//         for each neighbor of current
//             // d(current,neighbor) is the weight of the edge from current to neighbor
//             // tentative_gScore is the distance from start to the neighbor through current
//             tentative_gScore := gScore[current] + d(current, neighbor)
//             if tentative_gScore < gScore[neighbor]
//                 // This path to neighbor is better than any previous one. Record it!
//                 cameFrom[neighbor] := current
//                 gScore[neighbor] := tentative_gScore
//                 fScore[neighbor] := gScore[neighbor] + h(neighbor)
//                 if neighbor not in openSet
//                     openSet.add(neighbor)

//     // Open set is empty but goal was never reached
//     return failure

func (pf *PathFinder) neighbors(p point.Point2D) []point.Point2D {
	var ps []point.Point2D
	for _, dir := range []Direction{North, South, West, East} {
		if p := dir.offset(p); pf.canvas.At(p) != int(TileWall) {
			ps = append(ps, p)
		}
	}
	return ps
}

func (pf *PathFinder) f(p point.Point2D) int {
	if f, ok := pf.fs[p]; ok {
		return f
	}
	return math.MaxInt64
}

func (pf *PathFinder) g(p point.Point2D) int {
	if g, ok := pf.gs[p]; ok {
		return g
	}
	return math.MaxInt64
}

func (pf *PathFinder) h(p point.Point2D) int {
	return abs(p.X-pf.goal.X) + abs(p.Y-pf.goal.Y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
