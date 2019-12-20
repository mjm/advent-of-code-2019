package day18

// Path holds information about the shortest path that was found between
// significant points (a key or the start) on the map.
type Path struct {
	// Distance is the number of moves between the start and end of the path.
	Distance int
	// KeysNeeded is a bitset indicating which doors exist along the path, the
	// keys for which will need to have already been visited before this path
	// can be taken.
	KeysNeeded uint32
}

// CanVisit checks if the set of keys provided can open all the doors along this
// path.
func (p Path) CanVisit(keys uint32) bool {
	return p.KeysNeeded&keys == p.KeysNeeded
}
