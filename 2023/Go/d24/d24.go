package d24

import (
	"aoc2023/common"
	"strings"
)

const TEST1 string = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

const (
	lo float64 = 200_000_000_000_000
	hi float64 = 400_000_000_000_000
)

type D3 struct {
	X, Y, Z float64
}

type D2 struct {
	X, Y float64
}

type Stone struct {
	Position, Velocity D3
}

func ReadStones(input string) (stones []Stone) {
	stones = []Stone{}
	for _, line := range strings.Split(input, "\n") {
		entries := strings.Split(strings.Replace(line, "@", ",", -1), ",")
		stone := Stone{
			D3{
				float64(common.StrToInt(strings.Trim(entries[0], " "))),
				float64(common.StrToInt(strings.Trim(entries[1], " "))),
				float64(common.StrToInt(strings.Trim(entries[2], " ")))},
			D3{
				float64(common.StrToInt(strings.Trim(entries[3], " "))),
				float64(common.StrToInt(strings.Trim(entries[4], " "))),
				float64(common.StrToInt(strings.Trim(entries[5], " ")))},
		}
		stones = append(stones, stone)
	}
	return
}

func Intersects(a, b Stone) (D2, bool) {
	a2 := D2{a.Velocity.X, a.Velocity.Y}
	b2 := D2{b.Velocity.X, b.Velocity.Y}
	d2 := D2{b.Position.X - a.Position.X, b.Position.Y - a.Position.Y}
	ab := (a2.X * b2.Y) - (a2.Y * b2.X)
	db := (d2.X * b2.Y) - (d2.Y * b2.X)
	if ab == 0.0 {
		return D2{-1, -1}, false
	}
	div := db / ab
	return D2{a.Position.X + a.Velocity.X*div, a.Position.Y + a.Velocity.Y*div}, true
}

func IntersectCount(stones []Stone) (answer int) {
	for i := 0; i < len(stones)-1; i++ {
		for j := i + 1; j < len(stones); j++ {
			s1, s2 := stones[i], stones[j]
			if point, intersects := Intersects(s1, s2); intersects {
				if point.X >= lo && point.X <= hi && point.Y >= lo && point.Y <= hi {
					dx, dy := point.X-s1.Position.X, point.Y-s1.Position.Y
					if (dx > 0) == (s1.Velocity.X > 0) && (dy > 0) == (s1.Velocity.Y > 0) {
						dx, dy := point.X-s2.Position.X, point.Y-s2.Position.Y
						if (dx > 0) == (s2.Velocity.X > 0) && (dy > 0) == (s2.Velocity.Y > 0) {
							answer++
						}
					}
				}
			}
		}
	}
	return
}

func Solve(input string, part int) (answer int) {
	stones := ReadStones(input)
	answer = IntersectCount(stones)
	return
}
