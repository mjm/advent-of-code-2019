package day15

import "github.com/mjm/advent-of-code-2019/pkg/point"

// Fill spreads the tile destination to adjacent spaces until the
// entire canvas is filled. It returns the number of steps required
// to do so.
func Fill(c *Canvas) int {
	for i := 0; true; i++ {
		var ps []point.Point2D
		var anyOpen bool
		for p, val := range c.paint {
			tile := Tile(val)
			if tile == TileDestination {
				ps = append(ps, p)
			} else if tile == TilePassable || tile == TileStart {
				anyOpen = true
			}
		}

		if !anyOpen {
			return i
		}

		for _, p := range ps {
			for _, dir := range []Direction{North, South, West, East} {
				if p := dir.offset(p); c.At(p) == int(TilePassable) || c.At(p) == int(TileStart) {
					c.Paint(p, int(TileDestination))
				}
			}
		}
	}

	panic("how did I get here?")
}
