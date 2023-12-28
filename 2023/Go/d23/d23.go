package d23

import (
	"log"
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

type Plate struct {
	d, y, x int
}

var (
	grid       []string
	start, end Point
	points     []Point
	graph      map[Point]map[Point]int
	visited    map[Point]bool
)

func InGrid(p Point) bool {
	return p.y >= 0 && p.x >= 0 && p.y < len(grid) && p.x < len(grid[0])
}

func ThreeNeighborPoints() {
	points = []Point{start}
	for y, row := range grid {
		for x, char := range strings.Split(row, "") {
			if char == "#" {
				continue
			}
			nbr := 0
			for _, np := range [4]Point{{y - 1, x}, {y + 1, x}, {y, x - 1}, {y, x + 1}} {
				if InGrid(np) && grid[np.y][np.x] != '#' {
					nbr++
				}
			}
			if nbr > 2 {
				points = append(points, Point{y, x})
			}
		}
	}
	points = append(points, end)
}

func BuildGraph() {
	visited = map[Point]bool{}
	graph = map[Point]map[Point]int{}
	dirs := map[byte][]Point{
		'^': {{-1, 0}},
		'v': {{1, 0}},
		'<': {{0, -1}},
		'>': {{0, 1}},
		'.': {{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
	}
	for _, point := range points {
		graph[point] = map[Point]int{}
	}
	for _, p := range points {
		sy, sx := p.y, p.x
		stack := []Plate{{0, sy, sx}}
		visited[Point{sy, sx}] = true
		for len(stack) > 0 {
			var plate Plate
			sl := len(stack)
			if sl == 1 {
				plate = stack[0]
				stack = []Plate{}
			} else {
				plate = stack[sl-1]
				stack = stack[0 : sl-2]
			}
			d, y, x := plate.d, plate.y, plate.x
			if d != 0 && slices.Contains(points, Point{y, x}) {
				graph[Point{sy, sx}][Point{y, x}] = d
				continue
			}
			for _, dir := range dirs[grid[y][x]] {
				ny, nx := y+dir.y, x+dir.x
				if InGrid(Point{ny, nx}) && grid[ny][nx] != '#' && !visited[Point{ny, nx}] {
					stack = append(stack, Plate{d + 1, ny, nx})
					visited[Point{ny, nx}] = true
				}
			}
		}
	}
}

func DFS(p Point) int {
	if p.y == end.y && p.x == end.x {
		return 0
	}
	m := 0
	visited[p] = true
	for np, d := range graph[p] {
		if !visited[np] {
			if _, exists := graph[p][np]; !exists {
				d = 0
			}
			m = max(m, DFS(np)+d)
		}
	}
	delete(visited, p)
	return m
}

func Solve(input string, part int) (answer int) {
	grid = strings.Split(input, "\n")
	start = Point{0, strings.Index(grid[0], ".")}
	end = Point{len(grid) - 1, strings.Index(grid[len(grid)-1], ".")}
	ThreeNeighborPoints()
	BuildGraph()
	log.Printf("graph=%+v", graph)
	visited = map[Point]bool{}
	answer = DFS(start)
	return
}
