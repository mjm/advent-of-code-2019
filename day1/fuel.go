package day1

func FuelRequired(m Mass) int {
	return (m.Mass() / 3) - 2
}

func TotalFuelRequired(m Mass) int {
	n := FuelRequired(m)
	if n < 0 {
		return 0
	}

	if n > 0 {
		n += TotalFuelRequired(&fuel{mass: n})
	}
	return n
}
