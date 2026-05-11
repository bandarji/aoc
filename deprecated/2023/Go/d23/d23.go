package d23

import (
	"slices"
	"strings"
)

const TEST1 string = `#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#`

type Point struct {
	y, x int
}

type Distance struct {
	p Point
	d int
}

func Endpoints(grid []string) (start, end Point) {
	start = Point{0, strings.Index(grid[0], ".")}
	end = Point{len(grid) - 1, strings.Index(grid[len(grid)-1], ".")}
	return
}

func Dirs(ch rune, part int) (dirs []Point) {
	if part == 2 {
		dirs = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	} else {
		switch ch {
		case '^':
			dirs = []Point{{-1, 0}}
		case 'v':
			dirs = []Point{{1, 0}}
		case '<':
			dirs = []Point{{0, -1}}
		case '>':
			dirs = []Point{{0, 1}}
		case '.':
			dirs = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		}
	}
	return
}

func Hike(grid []string, start, end Point, part int) (answer int) {
	points := []Point{start, end}
	for r, row := range grid {
		for c, ch := range row {
			if ch == '#' {
				continue
			}
			neighbors := 0
			for _, np := range []Point{{r - 1, c}, {r + 1, c}, {r, c - 1}, {r, c + 1}} {
				if np.y >= 0 && np.y < len(grid) && np.x >= 0 && np.x < len(grid[0]) && grid[np.y][np.x] != '#' {
					neighbors++
				}
			}
			if neighbors > 2 {
				points = append(points, Point{r, c})
			}
		}
	}
	graph := map[Point]map[Point]int{}
	for _, p := range points {
		graph[p] = map[Point]int{}
	}

	for _, p := range points {
		stack := []Distance{{p, 0}}
		visited := map[Point]bool{p: true}
		for len(stack) > 0 {
			plate := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n, r, c := plate.d, plate.p.y, plate.p.x
			if n != 0 && slices.Contains(points, Point{r, c}) {
				graph[p][Point{r, c}] = n
				continue
			}
			for _, np := range Dirs(rune(grid[r][c]), part) {
				nr, nc := r+np.y, c+np.x
				if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) && grid[nr][nc] != '#' && !visited[Point{nr, nc}] {
					stack = append(stack, Distance{Point{nr, nc}, n + 1})
					visited[Point{nr, nc}] = true
				}
			}
		}
	}
	// fmt.Println("Points")
	// for _, p := range points {
	// 	fmt.Printf("  %#v\n", p)
	// }
	// fmt.Println("\nGraph")
	// for k, v := range graph {
	// 	fmt.Printf("  %#v: %#v\n", k, v)
	// }
	seen := map[Point]bool{}
	answer = DFS(graph, start, end, seen)
	return
}

func DFS(graph map[Point]map[Point]int, point, end Point, seen map[Point]bool) (answer int) {
	if point == end {
		return 0
	}
	seen[point] = true
	m := -1000000
	for k, v := range graph[point] {
		if !seen[k] {
			m = max(m, DFS(graph, k, end, seen)+v)
		}
	}
	delete(seen, point)
	return m
}

func Solve(input string, part int) (answer int) {
	grid := strings.Split(input, "\n")
	start, end := Endpoints(grid)
	answer = Hike(grid, start, end, part)
	return
}
