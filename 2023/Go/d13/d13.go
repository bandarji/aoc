package d13

import (
	"log"
	"slices"
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

func GetAbove(s []string, r int) []string {
	if r >= len(s) {
		return s
	}
	rows := s[0:r]
	slices.Reverse(rows)
	return rows
}

func GetBelow(s []string, r int) []string {
	if r < 1 {
		return s
	}
	return s[r:]
}

func SummarizeReflections(pattern []string, part int) int {
	for r := 1; r < len(pattern); r++ {
		above := GetAbove(pattern, r)
		below := GetBelow(pattern, r)
		aboveLen := len(above)
		belowLen := len(below)
		log.Printf("above=%+v, below=%+v, lengths=%d/%d", above, below, aboveLen, belowLen)
		if aboveLen < len(below) {
			below = below[0:aboveLen]
		}
		if belowLen < len(above) {
			above = above[0:belowLen]
		}
		if slices.Equal(above, below) {
			return r
		}
	}
	return 0
}

func Solve(input string, part int) (answer int) {
	for _, pattern := range strings.Split(input, "\n\n") {
		answer += SummarizeReflections(strings.Split(pattern, "\n"), part)
		break
	}
	return
}

// def find_mirror(grid):
//     for r in range(1, len(grid)):
//         above = grid[:r][::-1]
//         below = grid[r:]

//         above = above[:len(below)]
//         below = below[:len(above)]

//         if above == below:
//             return r

//     return 0

// total = 0

// for block in open(0).read().split("\n\n"):
//     grid = block.splitlines()

//     row = find_mirror(grid)
//     total += row * 100

//     col = find_mirror(list(zip(*grid)))
//     total += col

// print(total)
