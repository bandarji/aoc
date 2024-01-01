package d24

import (
	"slices"
	"strconv"
	"strings"
)

const TEST1 string = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

const (
	px = iota
	py
	pz
	vx
	vy
	vz
)

func ReadStones(input string) (hail [][6]float64) {
	hail = [][6]float64{}
	for _, entry := range strings.Split(input, "\n") {
		s := strings.Replace(entry, "@", ",", -1)
		s = strings.Replace(s, ",", " ", -1)
		subHail := [6]float64{}
		for i, v := range strings.Fields(s) {
			subHail[i], _ = strconv.ParseFloat(v, 64)
		}
		hail = append(hail, subHail)
	}
	return
}

func IntersectCount(stones [][6]float64, lo, hi float64) (count int) {
	for i := 0; i < len(stones)-1; i++ {
		for j := i + 1; j < len(stones); j++ {
			s1, s2 := stones[i], stones[j]
			if p, yn := Intersects(s1, s2); yn {
				if p[px] >= lo && p[px] <= hi && p[py] >= lo && p[py] <= hi {
					dx, dy := p[px]-s1[px], p[py]-s1[py]
					if (dx > 0) == (s1[vx] > 0) && (dy > 0) == (s1[vy] > 0) {
						dx, dy = p[px]-s2[px], p[py]-s2[py]
						if (dx > 0) == (s2[vx] > 0) && (dy > 0) == (s2[vy] > 0) {
							count++
						}
					}
				}
			}
		}
	}
	return
}

func Intersects(s1, s2 [6]float64) ([2]float64, bool) {
	s11 := [2]float64{s1[vx], s1[vy]}
	s21 := [2]float64{s2[vx], s2[vy]}
	d := [2]float64{s2[px] - s1[px], s2[py] - s1[py]}
	d0 := (s11[0] * s21[1]) - (s11[1] * s21[0])
	if d0 == 0.0 {
		return s11, false
	}
	cp := ((d[0] * s21[1]) - (d[1] * s21[0])) / d0
	return [2]float64{s1[px] + s1[vx]*cp, s1[py] + s1[vy]*cp}, true
}

func MatchVelocity(dims int, p1, p2, v1 float64) (matches []int) {
	matches = []int{}
	pp := int(p1 - p2)
	pv := int(v1)
	for v := -dims; v < dims; v++ {
		if v != pv && pp%(v-pv) == 0 {
			matches = append(matches, v)
		}
	}
	return
}

func Intersect(m1, m2 []int) (intersections []int) {
	intersections = []int{}
	for _, value := range m1 {
		if slices.Contains(m2, value) {
			intersections = append(intersections, value)
		}
	}
	return
}

func ThrowOneStone(stones [][6]float64) (answer int) {
	return
}

func Solve(input string, part int, lo, hi float64) (answer int) {
	stones := ReadStones(input)
	if part == 1 {
		answer = IntersectCount(stones, lo, hi)
	} else {
		answer = ThrowOneStone(stones)
	}
	return
}
