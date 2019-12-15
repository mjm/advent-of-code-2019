package day12

import (
	"fmt"
	"strings"
)

type cycleFinder struct {
	system *MoonSystem

	x *cycle
	y *cycle
	z *cycle
}

type cycle struct {
	getter func(Vec3D) int
	seen   seenValuesMap
	length int
	offset int
}

type axisValuesKey string
type seenValuesMap map[axisValuesKey]int

func newCycleFinder(s *MoonSystem) *cycleFinder {
	return &cycleFinder{
		system: s,
		x:      newCycle(func(v Vec3D) int { return v.X }),
		y:      newCycle(func(v Vec3D) int { return v.Y }),
		z:      newCycle(func(v Vec3D) int { return v.Z }),
	}
}

func (cf *cycleFinder) process() {
	for i := 0; !(cf.x.isDone() && cf.y.isDone() && cf.z.isDone()); i++ {
		cf.x.process(i, cf.system.Moons)
		cf.y.process(i, cf.system.Moons)
		cf.z.process(i, cf.system.Moons)

		cf.system.advanceOnce()
	}
}

func (cf *cycleFinder) maxOffset() int {
	max := cf.x.offset
	if cf.y.offset > max {
		max = cf.y.offset
	}
	if cf.z.offset > max {
		max = cf.z.offset
	}
	return max
}

func newCycle(getter func(Vec3D) int) *cycle {
	return &cycle{
		getter: getter,
		seen:   make(seenValuesMap),
	}
}

func (c *cycle) isDone() bool {
	return c.seen == nil
}

func (c *cycle) process(i int, moons []*Moon) {
	if c.isDone() {
		return
	}

	vals := make([][2]int, 0, len(moons))
	for _, moon := range moons {
		vals = append(vals, [2]int{c.getter(moon.Pos), c.getter(moon.Vel)})
	}
	key := makeKey(vals)

	if prev, ok := c.seen[key]; ok {
		c.offset = prev
		c.length = i - prev
		c.seen = nil
	} else {
		c.seen[key] = i
	}
}

func makeKey(ns [][2]int) axisValuesKey {
	var s strings.Builder
	for _, vals := range ns {
		fmt.Fprintf(&s, "%d,%d,", vals[0], vals[1])
	}
	return axisValuesKey(s.String())
}
