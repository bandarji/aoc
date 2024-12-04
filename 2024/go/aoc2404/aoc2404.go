package aoc2404

import "strings"

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

func Aoc240401(input string) (count int) {
	grid := strings.Split(input, "\n")
	rows := len(grid)
	cols := len(grid[0])
	for r := range rows {
		for c := range cols {
			if c+3 < cols && grid[r][c] == 'X' && grid[r][c+1] == 'M' && grid[r][c+2] == 'A' && grid[r][c+3] == 'S' {
				count++
			}
			if r+3 < rows && grid[r][c] == 'X' && grid[r+1][c] == 'M' && grid[r+2][c] == 'A' && grid[r+3][c] == 'S' {
				count++
			}
			if c+3 < cols && grid[r][c] == 'S' && grid[r][c+1] == 'A' && grid[r][c+2] == 'M' && grid[r][c+3] == 'X' {
				count++
			}
			if r+3 < rows && grid[r][c] == 'S' && grid[r+1][c] == 'A' && grid[r+2][c] == 'M' && grid[r+3][c] == 'X' {
				count++
			}
			if r+3 < rows && c+3 < cols && grid[r][c] == 'X' && grid[r+1][c+1] == 'M' && grid[r+2][c+2] == 'A' && grid[r+3][c+3] == 'S' {
				count++
			}
			if r-3 >= 0 && c+3 < cols && grid[r][c] == 'X' && grid[r-1][c+1] == 'M' && grid[r-2][c+2] == 'A' && grid[r-3][c+3] == 'S' {
				count++
			}
			if r+3 < rows && c+3 < cols && grid[r][c] == 'S' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'M' && grid[r+3][c+3] == 'X' {
				count++
			}
			if r-3 >= 0 && c+3 < cols && grid[r][c] == 'S' && grid[r-1][c+1] == 'A' && grid[r-2][c+2] == 'M' && grid[r-3][c+3] == 'X' {
				count++
			}
		}
	}
	return
}

func Aoc240402(input string) (count int) {
	grid := strings.Split(input, "\n")
	rows := len(grid)
	cols := len(grid[0])
	for r := range rows {
		for c := range cols {
			if r+2 < rows && c+2 < cols && grid[r][c] == 'M' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'S' && grid[r+2][c] == 'M' && grid[r][c+2] == 'S' {
				count++
			}
			if r+2 < rows && c+2 < cols && grid[r][c] == 'M' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'S' && grid[r+2][c] == 'S' && grid[r][c+2] == 'M' {
				count++
			}
			if r+2 < rows && c+2 < cols && grid[r][c] == 'S' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'M' && grid[r+2][c] == 'M' && grid[r][c+2] == 'S' {
				count++
			}
			if r+2 < rows && c+2 < cols && grid[r][c] == 'S' && grid[r+1][c+1] == 'A' && grid[r+2][c+2] == 'M' && grid[r+2][c] == 'S' && grid[r][c+2] == 'M' {
				count++
			}
		}
	}
	return
}
