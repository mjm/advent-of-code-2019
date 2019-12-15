package mathalg

// GCD2 finds the greatest common divisor of two integers
func GCD2(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// GCD finds the greatest common divisor of an arbitrary number of integers
func GCD(ns ...int) int {
	if len(ns) < 2 {
		panic("gcd is not defined for less than two numbers")
	}

	gcd := GCD2(ns[0], ns[1])
	for _, n := range ns[2:] {
		gcd = GCD2(gcd, n)
	}
	return gcd
}
