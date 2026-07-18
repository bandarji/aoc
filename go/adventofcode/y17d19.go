package adventofcode

import (
	"fmt"
	"strings"
)

const y17d19TestInput2 string = `     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 

`

func y17d19(input string, part int) (answer string) {
	grid := y17d19ParseGrid(input)
	path, steps := y17d19FindPath(grid)
	if part == 1 {
		answer = path
	} else {
		answer = fmt.Sprintf("%d", steps)
	}
	return
}

func y17d19ParseGrid(input string) (grid []string) {
	// Padding to avoid boundary checks
	lines := strings.Split(input, "\n")
	grid = append(grid, "")
	for _, line := range lines {
		grid = append(grid, fmt.Sprintf(" %s ", line))
	}
	s := strings.Repeat(" ", len(lines[0]))
	grid = append(grid, s)
	grid[0] = s
	return grid
}

func y17d19FindPath(grid []string) (path string, steps int) {
	facing := [2]int{0, 1}
	y, x := 1, strings.Index(grid[1], "|")
	vertical := true
	steps++
	for {
		nx, ny := x+facing[0], y+facing[1]
		next := grid[ny][nx]
		if next == ' ' {
			break
		}
		if next >= 'A' && next <= 'Z' {
			path = fmt.Sprintf("%s%c", path, next)
		}
		if next == '+' {
			if vertical {
				vertical = false
				if grid[ny][nx-1] != ' ' {
					facing = [2]int{-1, 0}
				} else {
					facing = [2]int{1, 0}
				}
			} else {
				vertical = true
				if grid[ny-1][nx] != ' ' {
					facing = [2]int{0, -1}
				} else {
					facing = [2]int{0, 1}
				}
			}
		}
		x, y = nx, ny
		steps++
	}
	return
}
