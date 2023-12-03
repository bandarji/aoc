package d03

import (
	"aoc2023/common"
	"strings"
)

// type LocatedSymbol struct {
// 	Y int
// 	X int
// }

// func (s LocatedSymbol) IsAdjacent(n LocatedNumber) bool {
// 	for x := n.X; x <= n.X+n.Length; x++ {
// 		if (s.X == x && s.Y == y) ||
// 			(s.X == x && s.Y == y-1) ||
// 			(s.X == x && s.Y == y+1) {
// 			return true
// 		}
// 	}
// 	return false
// }

// type LocatedNumber struct {
// 	Y      int
// 	X      int
// 	Length int
// 	Value  int
// 	Digits []string
// }

// func (n *LocatedNumber) Start(y, x int) {
// 	n.Y, n.X = y, x
// }

// func (n *LocatedNumber) AddDigit(d rune) {
// 	n.Digits = append(n.Digits, string(d))
// }

// func (n *LocatedNumber) CalcValue() {
// 	n.Value = common.StrToInt(strings.Join(n.Digits, ""))
// }

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
	schematic := [][]rune{}
	for _, line := range strings.Split(input, "\n") {
		schematic = append(schematic, []rune(line))
	}
	seen := map[[2]int]bool{}
	height := len(schematic)
	width := len(schematic[0])
	deltas := [3]int{-1, 0, 1}
	var char rune

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			char = schematic[y][x]
			if char != '.' && !isDigit(char) {
				adj := []int{}
				gear := isGear(schematic[y][x])
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
							for ; ex < height && isDigit(schematic[y][ex]); ex++ {
							}
							if seen[[2]int{y, sx}] {
								continue
							}
							number := common.StrToInt(string(schematic[y][sx:ex]))
							if part == 1 {
								answer += number
							}
							adj = append(adj, number)
							seen[[2]int{y, sx}] = true
						}
					}
				}
				if part == 2 && gear && len(adj) == 2 {
					answer += adj[0] * adj[1]
				}
			}
		}
	}
	return
}
