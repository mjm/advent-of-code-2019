package day12

// Vec3D is a three-dimensional vector of integers.
type Vec3D struct {
	X int
	Y int
	Z int
}

// Add add each component of b to the corresponding component of a, and stores the result
// in a.
func (a *Vec3D) Add(b Vec3D) {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
}
