package adventofcode

import (
	"strings"
)

func y15d18(input string, steps int, part int) (answer int) {
	grid := y15d18MakeGrid(input)
	for i := range steps {
		_ = i
		if part == 2 {
			grid = y15d18KeepCornersLit(grid)
		}
		grid = y15d18Step(grid)
		if part == 2 {
			grid = y15d18KeepCornersLit(grid)
		}
	}
	answer = y15d18CountLights(grid)
	return
}

func y15d18KeepCornersLit(grid [][]string) [][]string {
	grid[0][0] = "#"
	grid[0][len(grid[0])-1] = "#"
	grid[len(grid)-1][0] = "#"
	grid[len(grid)-1][len(grid[0])-1] = "#"
	return grid
}

func y15d18CountLights(grid [][]string) (count int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == "#" {
				count++
			}
		}
	}
	return
}

func y15d18Step(grid [][]string) (next [][]string) {
	next = [][]string{}
	for r, row := range grid {
		next = append(next, make([]string, len(grid[0])))
		for c, cell := range row {
			var neighbors int
			for rDiff := -1; rDiff <= 1; rDiff++ {
				for cDiff := -1; cDiff <= 1; cDiff++ {
					if !(rDiff == 0 && cDiff == 0) {
						nextRow := r + rDiff
						nextCol := c + cDiff
						if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) &&
							grid[nextRow][nextCol] == "#" {
							neighbors++
						}
					}
				}
			}
			switch {
			case cell == "#" && (neighbors == 2 || neighbors == 3):
				next[r][c] = "#"
			case cell == "." && neighbors == 3:
				next[r][c] = "#"
			default:
				next[r][c] = "."
			}
		}
	}
	return
}

func y15d18MakeGrid(input string) [][]string {
	grid := [][]string{}
	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(row, ""))
	}
	return grid
}
