package d21

import (
	"strings"
)

const TEST1 string = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

type Spot struct {
	Y, X, S int
}

func ReadGarden(input string) (start Spot, garden [][]string) {
	garden = [][]string{}
	for y, row := range strings.Split(input, "\n") {
		garden = append(garden, strings.Split(row, ""))
		x := strings.Index(row, "S")
		if x != -1 {
			start = Spot{X: x, Y: y, S: 0}
		}
	}
	return
}

func GetSteps(start Spot, garden [][]string, steps int) (answer int) {
	q := []Spot{start}
	seen := map[Spot]bool{}
	for len(q) > 0 {
		spot := q[0]
		q = q[1:]
		if spot.S == 0 {
			answer += 1
			continue
		}
		for _, adjacent := range Adjacent(&garden, spot, spot.S) {
			if !seen[adjacent] {
				seen[adjacent] = true
				q = append(q, adjacent)
			}
		}
	}
	return
}

func Adjacent(garden *[][]string, spot Spot, step int) (spots []Spot) {
	spots = []Spot{}
	for _, dir := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		y, x := spot.Y+dir[0], spot.X+dir[1]
		if y < 0 || x < 0 || y >= len(*garden) || x >= len((*garden)[0]) {
			continue
		}
		if (*garden)[y][x] == "#" {
			continue
		}
		spots = append(spots, Spot{Y: y, X: x, S: step - 1})
	}
	return
}

func GetSteps2(start Spot, garden [][]string, steps int) (answer int) {
	return
}

// In exactly 6 steps, he can still reach 16 garden plots.
// In exactly 10 steps, he can reach any of 50 garden plots.
// In exactly 50 steps, he can reach 1594 garden plots.
// In exactly 100 steps, he can reach 6536 garden plots.
// In exactly 500 steps, he can reach 167004 garden plots.
// In exactly 1000 steps, he can reach 668697 garden plots.
// In exactly 5000 steps, he can reach 16733044 garden plots.

func Solve(input string, steps, part int) (answer int) {
	start, garden := ReadGarden(input)
	start.S = steps
	if part == 1 {
		answer = GetSteps(start, garden, steps)
	}
	if part == 2 {
		answer = GetSteps2(start, garden, steps)
	}
	return
}
