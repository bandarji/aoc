package d10

import (
	"log"
	"sort"
	"strings"
)

const TEST1 string = `.....
.S-7.
.|.|.
.L-J.
.....`

const TEST2 string = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

const TEST3 string = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

const TEST4 string = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

func ReadDiagram(input string) (grid [][]string, sr, sc int) {
	grid = [][]string{}
	for y, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "S") {
			for x := 0; x < len(line); x++ {
				if line[x] == 'S' {
					sr, sc = y, x
					break
				}
			}
		}
		grid = append(grid, strings.Split(line, ""))
	}
	return
}

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.

func FarthestSpot(grid [][]string, sr, sc int) (seen map[[2]int]bool, farthest int) {
	seen = map[[2]int]bool{}
	seen[[2]int{sr, sc}] = true
	q := [][2]int{{sr, sc}}
	height := len(grid)
	width := len(grid[0])
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		row, col := current[0], current[1]
		// up
		if row > 0 && strings.Contains("S|JL", grid[row][col]) && strings.Contains("|7F", grid[row-1][col]) && !seen[[2]int{row - 1, col}] {
			// log.Printf("col = %d, row = %d", row-1, col)
			seen[[2]int{row - 1, col}] = true
			q = append(q, [2]int{row - 1, col})
		}
		// down
		if row < width-1 && strings.Contains("S|7F", grid[row][col]) && strings.Contains("|JL", grid[row+1][col]) && !seen[[2]int{row + 1, col}] {
			// log.Printf("col = %d, row = %d", row+1, col)
			seen[[2]int{row + 1, col}] = true
			q = append(q, [2]int{row + 1, col})
		}
		// left
		if col > 0 && strings.Contains("S-J7", grid[row][col]) && strings.Contains("-LF", grid[row][col-1]) && !seen[[2]int{row, col - 1}] {
			// log.Printf("col = %d, row = %d", row, col-1)
			seen[[2]int{row, col - 1}] = true
			q = append(q, [2]int{row, col - 1})
		}
		// right
		if col < height-1 && strings.Contains("S-LF", grid[row][col]) && strings.Contains("-J7", grid[row][col+1]) && !seen[[2]int{row, col + 1}] {
			// log.Printf("col = %d, row = %d", row+1, col)
			seen[[2]int{row, col + 1}] = true
			q = append(q, [2]int{row, col + 1})
		}
	}
	farthest = len(seen) / 2
	return
}

func TilesEnclosed(grid [][]string, loop map[[2]int]bool, sr, sc int) (answer int) {
	q := [][2]int{{sr, sc}}
	height := len(grid)
	width := len(grid[0])
	seen := map[[2]int]bool{}
	possibles := strings.Split("|-JL7F", "")
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		row, col := current[0], current[1]
		// up
		if row > 0 && strings.Contains("S|JL", grid[row][col]) && strings.Contains("|7F", grid[row-1][col]) && !seen[[2]int{row - 1, col}] {
			// log.Printf("col = %d, row = %d", row-1, col)
			seen[[2]int{row - 1, col}] = true
			q = append(q, [2]int{row - 1, col})
			if grid[row][col] == "S" {
				possibles = IntersectStringSlices(possibles, strings.Split("|JL", ""))
			}
		}
		// down
		if row < width-1 && strings.Contains("S|7F", grid[row][col]) && strings.Contains("|JL", grid[row+1][col]) && !seen[[2]int{row + 1, col}] {
			// log.Printf("col = %d, row = %d", row+1, col)
			seen[[2]int{row + 1, col}] = true
			q = append(q, [2]int{row + 1, col})
			if grid[row][col] == "S" {
				possibles = IntersectStringSlices(possibles, strings.Split("|7F", ""))
			}
		}
		// left
		if col > 0 && strings.Contains("S-J7", grid[row][col]) && strings.Contains("-LF", grid[row][col-1]) && !seen[[2]int{row, col - 1}] {
			// log.Printf("col = %d, row = %d", row, col-1)
			seen[[2]int{row, col - 1}] = true
			q = append(q, [2]int{row, col - 1})
			if grid[row][col] == "S" {
				possibles = IntersectStringSlices(possibles, strings.Split("-J7", ""))
			}
		}
		// right
		if col < height-1 && strings.Contains("S-LF", grid[row][col]) && strings.Contains("-J7", grid[row][col+1]) && !seen[[2]int{row, col + 1}] {
			// log.Printf("col = %d, row = %d", row+1, col)
			seen[[2]int{row, col + 1}] = true
			q = append(q, [2]int{row, col + 1})
			if grid[row][col] == "S" {
				possibles = IntersectStringSlices(possibles, strings.Split("-LF", ""))
			}
		}
	}

	grid[sr][sc] = possibles[0] // Replace 'S' with what 'S' actually is

	// | is a vertical pipe connecting north and south.
	// - is a horizontal pipe connecting east and west.
	// L is a 90-degree bend connecting north and east.
	// J is a 90-degree bend connecting north and west.
	// 7 is a 90-degree bend connecting south and west.
	// F is a 90-degree bend connecting south and east.
	// . is ground; there is no pipe in this tile.

	for _, row := range grid {
		log.Println(strings.Join(row, ""))
	}

	// Visual map, just for me - I'm losing my mind by this point
	// for _, row := range grid {
	// 	line := strings.Join(row, "")
	// 	line = strings.Replace(line, "|", "║", -1)
	// 	line = strings.Replace(line, "-", "═", -1)
	// 	line = strings.Replace(line, "L", "╚", -1)
	// 	line = strings.Replace(line, "J", "╝", -1)
	// 	line = strings.Replace(line, "7", "╗", -1)
	// 	line = strings.Replace(line, "F", "╔", -1)
	// 	log.Println(line)
	// }

	outside := map[[2]int]bool{}
	up := false
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			inside := false
			char := grid[row][col]
			switch {
			case char == "|":
				inside = !inside
			case char == "L":
				up = true
			case char == "F":
				up = false
			case char == "7":
				if !up {
					inside = !inside
				}
			case char == "J":
				if up {
					inside = !inside
				}
			}
			if !inside {
				outside[[2]int{row, col}] = true
			}
		}
	}

	// Viz #2
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			grid[row][col] = "."
		}
	}
	for coord := range outside {
		grid[coord[0]][coord[1]] = "#"
	}
	for _, row := range grid {
		line := strings.Join(row, "")
		log.Println(line)
	}

	log.Printf("%+v\n %d", outside, len(loop))
	answer = height*width - len(outside)
	return
}

func TilesEnclosed2(grid [][]string, loop map[[2]int]bool, sr, sc int) (answer int) {
	q := [][2]int{{sr, sc}}
	height := len(grid)
	width := len(grid[0])
	seen := map[[2]int]bool{}
	possibles := strings.Split("|-JL7F", "")
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		row, col := current[0], current[1]
		// up
		if row > 0 && strings.Contains("S|JL", grid[row][col]) && strings.Contains("|7F", grid[row-1][col]) && !seen[[2]int{row - 1, col}] {
			// log.Printf("col = %d, row = %d", row-1, col)
			seen[[2]int{row - 1, col}] = true
			q = append(q, [2]int{row - 1, col})
			if grid[row][col] == "S" {
				possibles = IntersectStringSlices(possibles, strings.Split("|JL", ""))
			}
		}
		// down
		if row < width-1 && strings.Contains("S|7F", grid[row][col]) && strings.Contains("|JL", grid[row+1][col]) && !seen[[2]int{row + 1, col}] {
			// log.Printf("col = %d, row = %d", row+1, col)
			seen[[2]int{row + 1, col}] = true
			q = append(q, [2]int{row + 1, col})
			if grid[row][col] == "S" {
				possibles = IntersectStringSlices(possibles, strings.Split("|7F", ""))
			}
		}
		// left
		if col > 0 && strings.Contains("S-J7", grid[row][col]) && strings.Contains("-LF", grid[row][col-1]) && !seen[[2]int{row, col - 1}] {
			// log.Printf("col = %d, row = %d", row, col-1)
			seen[[2]int{row, col - 1}] = true
			q = append(q, [2]int{row, col - 1})
			if grid[row][col] == "S" {
				possibles = IntersectStringSlices(possibles, strings.Split("-J7", ""))
			}
		}
		// right
		if col < height-1 && strings.Contains("S-LF", grid[row][col]) && strings.Contains("-J7", grid[row][col+1]) && !seen[[2]int{row, col + 1}] {
			// log.Printf("col = %d, row = %d", row+1, col)
			seen[[2]int{row, col + 1}] = true
			q = append(q, [2]int{row, col + 1})
			if grid[row][col] == "S" {
				possibles = IntersectStringSlices(possibles, strings.Split("-LF", ""))
			}
		}
	}
	grid[sr][sc] = possibles[0] // Replace 'S' with what 'S' actually is
	// // remove non-loop pipes
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if !seen[[2]int{row, col}] {
				grid[row][col] = "."
			}
		}
	}
	outside := map[[2]int]bool{}
	for r, row := range grid {
		inside := false // start each line outside
		up := false
		for c, char := range row {
			switch {
			case char == "|":
				inside = !inside
			case strings.Contains("LF", char):
				if char == "L" {
					up = true
				} else {
					up = false
				}
			case strings.Contains("7J", char):
				if up {
					if char != "J" {
						inside = !inside
					}
				} else {
					if char != "7" {
						inside = !inside
					}
				}
				up = false
			}
			if !inside {
				outside[[2]int{r, c}] = true
			}
		}
	}
	answer = width*height - len(OrLocations(outside, loop))
	return
}

func OrLocations(s1, s2 map[[2]int]bool) (either map[[2]int]bool) {
	either = map[[2]int]bool{}
	for key := range s1 {
		either[key] = true
	}
	for key := range s2 {
		either[key] = true
	}
	return
}

func IntersectStringSlices(s1, s2 []string) []string {
	counts := map[string]int{}
	if len(s1) == 0 || len(s2) == 0 {
		return []string{}
	}
	for _, element := range s1 {
		if _, exists := counts[element]; !exists {
			counts[element]++
		}
	}
	for _, element := range s2 {
		if _, exists := counts[element]; exists {
			counts[element]++
		}
	}
	intersection := []string{}
	for word, count := range counts {
		if count >= 2 {
			intersection = append(intersection, word)
		}
	}
	sort.Strings(intersection)
	return intersection
}

func Solve(input string, part int) (answer int) {
	grid, sr, sc := ReadDiagram(input)
	loop, farthest := FarthestSpot(grid, sr, sc)
	if part == 1 {
		answer = farthest
	}
	if part == 2 {
		answer = TilesEnclosed2(grid, loop, sr, sc)
	}
	_ = loop
	return
}
