package adventofcode

import (
	"strings"
)

var y17d22Directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func y17d22(input string, cycles, part int) (answer int) {
	grid, pos := y17d22Grid(input)
	if part == 1 {
		answer = y17d22Part1(grid, pos, cycles)
	} else {
		answer = y17d22Part2(grid, pos, cycles)
	}
	return
}

func y17d22Part1(grid map[[2]int]int, pos [2]int, cycles int) (count int) {
	dir := 0
	for range cycles {
		switch grid[pos] {
		case 0:
			count++
			dir = (dir + 3) % 4
			grid[pos] = 1
		case 1:
			dir = (dir + 1) % 4
			grid[pos] = 0
		}
		pos[0] += y17d22Directions[dir][0]
		pos[1] += y17d22Directions[dir][1]
	}
	return
}

func y17d22Part2(grid map[[2]int]int, pos [2]int, cycles int) (count int) {
	dir := 0
	for range cycles {
		switch grid[pos] {
		case 0:
			dir = (dir + 3) % 4
			grid[pos] = 2
		case 1:
			dir = (dir + 1) % 4
			grid[pos] = 3
		case 2:
			grid[pos] = 1
			count++
		case 3:
			dir = (dir + 2) % 4
			grid[pos] = 0
		}
		pos[0] += y17d22Directions[dir][0]
		pos[1] += y17d22Directions[dir][1]
	}
	return
}

func y17d22Grid(input string) (grid map[[2]int]int, start [2]int) {
	grid = map[[2]int]int{}
	lines := strings.Split(input, "\n")
	for row, line := range lines {
		for col, char := range line {
			switch char {
			case '.':
				grid[[2]int{row, col}] = 0
			case '#':
				grid[[2]int{row, col}] = 1
			}
		}
	}
	mid := len(lines) / 2
	start = [2]int{mid, mid}
	return
}
