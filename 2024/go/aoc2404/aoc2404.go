package aoc2404

import "strings"

type Grid struct {
	grid       []string
	rows, cols int
}

const AOC2404_TEST string = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func (g Grid) findXmasCross(r, c int) (count int) {
	if r+2 < g.rows && c+2 < g.cols && g.grid[r][c] == 'M' &&
		g.grid[r+1][c+1] == 'A' && g.grid[r+2][c+2] == 'S' &&
		g.grid[r+2][c] == 'M' && g.grid[r][c+2] == 'S' {
		count++
	}
	if r+2 < g.rows && c+2 < g.cols && g.grid[r][c] == 'M' &&
		g.grid[r+1][c+1] == 'A' && g.grid[r+2][c+2] == 'S' &&
		g.grid[r+2][c] == 'S' && g.grid[r][c+2] == 'M' {
		count++
	}
	if r+2 < g.rows && c+2 < g.cols && g.grid[r][c] == 'S' &&
		g.grid[r+1][c+1] == 'A' && g.grid[r+2][c+2] == 'M' &&
		g.grid[r+2][c] == 'M' && g.grid[r][c+2] == 'S' {
		count++
	}
	if r+2 < g.rows && c+2 < g.cols && g.grid[r][c] == 'S' &&
		g.grid[r+1][c+1] == 'A' && g.grid[r+2][c+2] == 'M' &&
		g.grid[r+2][c] == 'S' && g.grid[r][c+2] == 'M' {
		count++
	}
	return
}

func (g Grid) FindXmas(r, c int) (count int) {
	if c+3 < g.cols && g.grid[r][c] == 'X' && g.grid[r][c+1] == 'M' &&
		g.grid[r][c+2] == 'A' && g.grid[r][c+3] == 'S' {
		count++
	}
	if r+3 < g.rows && g.grid[r][c] == 'X' && g.grid[r+1][c] == 'M' &&
		g.grid[r+2][c] == 'A' && g.grid[r+3][c] == 'S' {
		count++
	}
	if c+3 < g.cols && g.grid[r][c] == 'S' && g.grid[r][c+1] == 'A' &&
		g.grid[r][c+2] == 'M' && g.grid[r][c+3] == 'X' {
		count++
	}
	if r+3 < g.rows && g.grid[r][c] == 'S' && g.grid[r+1][c] == 'A' &&
		g.grid[r+2][c] == 'M' && g.grid[r+3][c] == 'X' {
		count++
	}
	if r+3 < g.rows && c+3 < g.cols && g.grid[r][c] == 'X' &&
		g.grid[r+1][c+1] == 'M' && g.grid[r+2][c+2] == 'A' && g.grid[r+3][c+3] == 'S' {
		count++
	}
	if r-3 >= 0 && c+3 < g.cols && g.grid[r][c] == 'X' &&
		g.grid[r-1][c+1] == 'M' && g.grid[r-2][c+2] == 'A' && g.grid[r-3][c+3] == 'S' {
		count++
	}
	if r+3 < g.rows && c+3 < g.cols && g.grid[r][c] == 'S' &&
		g.grid[r+1][c+1] == 'A' && g.grid[r+2][c+2] == 'M' && g.grid[r+3][c+3] == 'X' {
		count++
	}
	if r-3 >= 0 && c+3 < g.cols && g.grid[r][c] == 'S' &&
		g.grid[r-1][c+1] == 'A' && g.grid[r-2][c+2] == 'M' && g.grid[r-3][c+3] == 'X' {
		count++
	}
	return
}

func Aoc240401(input string) (count int) {
	g := Grid{
		grid: strings.Split(input, "\n"),
	}
	g.rows, g.cols = len(g.grid), len(g.grid[0])
	for r := range g.rows {
		for c := range g.cols {
			count += g.FindXmas(r, c)
		}
	}
	return
}

func Aoc240402(input string) (count int) {
	g := Grid{
		grid: strings.Split(input, "\n"),
	}
	g.rows, g.cols = len(g.grid), len(g.grid[0])
	for r := range g.rows {
		for c := range g.cols {
			count += g.findXmasCross(r, c)
		}
	}
	return
}
