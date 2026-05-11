package d03

import (
	"aoc2023/common"
	"fmt"
	"strings"
)

const TEST1 string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isGear(c rune) bool {
	return c == '*'
}

func Solve(input string, part int) (answer int) {
	var char rune
	schematic := [][]rune{}
	seen := map[string]bool{}
	deltas := [3]int{-1, 0, 1}
	for _, line := range strings.Split(input, "\n") {
		schematic = append(schematic, []rune(line))
	}
	height := len(schematic)
	width := len(schematic[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			char = schematic[y][x]
			if char != '.' && !isDigit(char) {
				adj := []int{}
				for _, dy := range deltas {
					for _, dx := range deltas {
						x := x + dx
						y := y + dy
						if dx == 0 && dx == dy ||
							y < 0 ||
							y > height ||
							x < 0 ||
							x >= width {
							continue
						}
						if isDigit(schematic[y][x]) {
							sx := x // start
							for ; sx >= 0 && isDigit(schematic[y][sx]); sx-- {
							}
							sx += 1
							ex := x
							for ; ex < width && isDigit(schematic[y][ex]); ex++ {
							}
							if seen[fmt.Sprintf("%d, %d", y, sx)] {
								continue
							}
							number := common.StrToInt(string(schematic[y][sx:ex]))
							if part == 1 {
								answer += number
							}
							adj = append(adj, number)
							seen[fmt.Sprintf("%d, %d", y, sx)] = true
						}
					}
				}
				if part == 2 && isGear(schematic[y][x]) && len(adj) == 2 {
					answer += adj[0] * adj[1]
				}
			}
		}
	}
	return
}
