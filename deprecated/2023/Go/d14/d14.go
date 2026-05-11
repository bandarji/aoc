package d14

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

const TEST1 string = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

var cache = make(map[string]int)

func MakeGrid(input string) (grid [][]string) {
	grid = [][]string{}
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}
	return
}

func Weight(grid [][]string) (weight int) {
	factor := len(grid)
	for y, row := range grid {
		count := 0
		for x := range row {
			if grid[y][x] == "O" {
				count++
			}
		}
		weight += factor * count
		factor--
	}
	return
}

func Key(grid [][]string) (key string) {
	key = ""
	for _, row := range grid {
		key += strings.Join(row, "")
	}
	hash := md5.Sum([]byte(key))
	key = hex.EncodeToString(hash[:])
	return
}

const cycleCount int = 1_000_000_000

// func DisplayGrid(grid [][]string, cycle int) {
// 	fmt.Printf("\nAfter %d cycle:\n", cycle)
// 	for _, row := range grid {
// 		for _, ch := range row {
// 			fmt.Print(ch)
// 		}
// 		fmt.Println()
// 	}
// }

func DisplayGrid(grid [][]string, msg string) {
	fmt.Println(msg)
	for _, row := range grid {
		for _, ch := range row {
			fmt.Print(ch)
		}
		fmt.Println()
	}
}

func TiltNorth(grid [][]string) [][]string {
	moved := false
	for {
		moved = false
		for y := 1; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				if grid[y][x] == "O" && grid[y-1][x] == "." {
					moved = true
					grid[y][x] = "."
					grid[y-1][x] = "O"
				}
			}
		}
		if !moved {
			break
		}
	}
	return grid
}

func TiltSouth(grid [][]string) [][]string {
	moved := false
	for {
		moved = false
		for y := 0; y < len(grid)-1; y++ {
			for x := 0; x < len(grid[0]); x++ {
				if grid[y][x] == "O" && grid[y+1][x] == "." {
					moved = true
					grid[y][x] = "."
					grid[y+1][x] = "O"
				}
			}
		}
		if !moved {
			break
		}
	}
	return grid
}

func TiltWest(grid [][]string) [][]string {
	moved := false
	for {
		moved = false
		for x := 1; x < len(grid[0]); x++ {
			for y := 0; y < len(grid); y++ {
				if grid[y][x] == "O" && grid[y][x-1] == "." {
					moved = true
					grid[y][x] = "."
					grid[y][x-1] = "O"
				}
			}
		}
		if !moved {
			break
		}
	}
	return grid
}

func TiltEast(grid [][]string) [][]string {
	moved := false
	for {
		moved = false
		for x := 0; x < len(grid[0])-1; x++ {
			for y := 0; y < len(grid); y++ {
				if grid[y][x] == "O" && grid[y][x+1] == "." {
					moved = true
					grid[y][x] = "."
					grid[y][x+1] = "O"
				}
			}
		}
		if !moved {
			break
		}
	}
	return grid
}

func Cycle(grid [][]string) [][]string {
	grid = TiltNorth(grid)
	grid = TiltWest(grid)
	grid = TiltSouth(grid)
	grid = TiltEast(grid)
	return grid
}

func Solve(input string, part int) (answer int) {
	grid := MakeGrid(input)
	if part == 1 {
		grid = TiltNorth(grid)
	} else {
		for c := 1; c <= cycleCount; c++ {
			grid = Cycle(grid)
			hash := Key(grid)
			if cycle, exists := cache[hash]; exists {
				cycleLength := c - cycle
				remaining := cycleCount - c
				remaining -= (remaining / cycleLength) * cycleLength
				for i := 0; i < remaining; i++ {
					grid = Cycle(grid)
				}
				break
			} else {
				cache[hash] = c
			}
		}
	}
	answer = Weight(grid)
	return
}
