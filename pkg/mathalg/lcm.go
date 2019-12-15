package mathalg

// LCM2 computes the least common multiple of two integers.
func LCM2(a, b int) int {
	return (a / GCD(a, b)) * b
}

// LCM computes the least common multiple of an arbitrary number of integers
func LCM(ns ...int) int {
	if len(ns) < 2 {
		panic("lcm is not defined for less than two numbers")
	}

	lcm := LCM2(ns[0], ns[1])
	for _, n := range ns[2:] {
		lcm = LCM2(lcm, n)
	}
	return lcm
}
