package d13

import (
	"strings"
)

const TEST1 string = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func Solve(input string, part int) (answer int) {
	hori, vert := 0, 0
	for _, block := range strings.Split(input, "\n\n") {
		grid := strings.Split(block, "\n")
		hori = FindMirror(grid, part)
		vert = FindMirror(Transpose(grid), part)
		answer += hori * 100
		answer += vert
	}
	return
}

func FindMirror(grid []string, part int) int {
	reflects := Reflects1
	if part == 2 {
		reflects = Reflects2
	}
	for i := 0; i < len(grid)-1; i++ {
		if reflects(grid, i) {
			return i + 1
		}
	}
	return 0
}

func Reflects1(grid []string, start int) bool {
	lo, up := start, start+1
	for lo >= 0 && up < len(grid) {
		if grid[lo] != grid[up] {
			return false
		}
		lo--
		up++
	}
	return true
}

func Fixable(s1, s2 string) bool {
	diffs := 0
	for i := 0; i < min(len(s1), len(s2)); i++ {
		if s1[i] != s2[i] {
			diffs++
		}
	}
	return diffs == 1
}

func Reflects2(grid []string, start int) bool {
	lo, up := start, start+1
	fixed := false
	for lo >= 0 && up < len(grid) {
		if grid[lo] != grid[up] {
			if fixed {
				return false
			} else if Fixable(grid[lo], grid[up]) {
				fixed = true
			} else {
				return false
			}
		}
		lo--
		up++
	}
	return fixed
}

func Transpose(grid []string) (ngrid []string) {
	ngrid = []string{}
	s := ""
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			s += string(grid[y][x])
		}
		ngrid = append(ngrid, s)
		s = ""
	}
	return
}
