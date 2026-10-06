package adventofcode

import (
	"math"
	"os"
	"slices"
	"strings"
)

var y17d21Initial = []string{
	".#.",
	"..#",
	"###",
}

func y17d21(input string, iterations int) (answer int) {
	grid := y17d21NewGrid(y17d21Initial)
	grid.y17d21SetRules(strings.Split(input, "\n"))
	for range iterations {
		grid.y17d21Run()
	}
	answer = grid.y17d21CountOn()
	return
}

type y17d21Pattern struct {
	output   []string
	variants []string
}

func y17d21NewPattern(initial, output string) *y17d21Pattern {
	p := &y17d21Pattern{
		output:   strings.Split(output, "/"),
		variants: []string{},
	}
	inputRows := strings.Split(initial, "/")
	p.y17d21MakeVariants(inputRows)
	return p
}

func (p *y17d21Pattern) y17d21Match(square []string) bool {
	return slices.Contains(p.variants, strings.Join(square, ""))
}

func (p *y17d21Pattern) y17d21MakeVariants(pattern []string) {
	for range 4 {
		pattern = y17d21Rotate(pattern)
		p.variants = append(p.variants, strings.Join(pattern, ""))
	}
	slices.Reverse(pattern)
	for range 4 {
		pattern = y17d21Rotate(pattern)
		p.variants = append(p.variants, strings.Join(pattern, ""))
	}
}

// y17d21Rotate rotates a square pattern 90 degrees clockwise.
func y17d21Rotate(pattern []string) []string {
	n := len(pattern)
	newRows := make([]string, n)
	for i := range n {
		b := strings.Builder{}
		for j := n - 1; j >= 0; j-- {
			b.WriteByte(pattern[j][i])
		}
		newRows[i] = b.String()
	}
	return newRows
}

func y17d21Split(sq []string) [][]string {
	rowCount := len(sq)
	if rowCount <= 3 {
		return [][]string{sq}
	}
	newSqLen := 3
	if rowCount%2 == 0 {
		newSqLen = 2
	}
	newSqs := [][]string{}
	for y := 0; y < rowCount; y += newSqLen {
		for x := 0; x < rowCount; x += newSqLen {
			newSq := make([]string, 0, newSqLen)

			for i := 0; i < newSqLen; i++ {
				newSq = append(
					newSq,
					sq[y+i][x:x+newSqLen],
				)
			}
			newSqs = append(newSqs, newSq)
		}
	}
	return newSqs
}

func y17d21Join(sqs [][]string) []string {
	sqCount := len(sqs)
	if sqCount == 1 {
		return sqs[0]
	}
	sqLen := len(sqs[0])
	sqPerRow := int(math.Sqrt(float64(sqCount)))
	rowCount := sqLen * sqPerRow
	grid := make([]string, 0, rowCount)
	y := 0
	sqNo := 0
	for y < rowCount {
		rows := make([]string, sqLen)
		for j := range sqLen {
			for x := range sqPerRow {
				rows[j] += sqs[sqNo+x][(y+j)%sqLen]
			}
		}
		sqNo += sqPerRow
		y += sqLen
		grid = append(grid, rows...)
	}
	return grid
}

type y17d21Grid struct {
	rules []*y17d21Pattern
	grid  []string
}

func y17d21NewGrid(initial []string) *y17d21Grid {
	return &y17d21Grid{
		rules: []*y17d21Pattern{},
		grid:  initial,
	}
}

func (g *y17d21Grid) y17d21SetRules(rules []string) {
	for _, rule := range rules {
		rule = strings.TrimSpace(rule)
		if rule == "" {
			continue
		}
		p0, p1, _ := strings.Cut(rule, " => ")
		g.rules = append(g.rules, y17d21NewPattern(p0, p1))
	}
}

func (g *y17d21Grid) y17d21Run() {
	newSquares := y17d21Split(g.grid)
	toJoin := make([][]string, 0, len(newSquares))
	for _, newSquare := range newSquares {
		matched := false
		for _, rule := range g.rules {
			if rule.y17d21Match(newSquare) {
				toJoin = append(toJoin, rule.output)
				matched = true
				break
			}
		}
		if !matched {
			os.Exit(1)
		}
	}
	g.grid = y17d21Join(toJoin)
}

func (g *y17d21Grid) y17d21CountOn() int {
	return strings.Count(strings.Join(g.grid, ""), "#")
}
