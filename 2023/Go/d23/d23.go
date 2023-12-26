package d23

import (
	"log"
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

var visited = make(map[Point]bool)

type Point struct {
	y, x, d int
}

func Pop(l *[]Point) Point {
	c := len(*l)
	e := (*l)[c-1]
	*l = (*l)[:c-1]
	return e
}

func GetEndPoints(grid []string) (start, end Point) {
	start = Point{y: 0, x: strings.Index(grid[0], "."), d: 0}
	end = Point{y: len(grid) - 1, x: strings.Index(grid[len(grid)-1], "."), d: 0}
	return
}

func FindDecisionPoints(grid []string, start, end Point) (points []Point) {
	points = []Point{start, end}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				continue
			}
			neighbors := 0
			for _, np := range []Point{{y - 1, x, 0}, {y + 1, x, 0}, {y, x - 1, 0}, {y, x + 1, 0}} {
				if np.x >= 0 && np.y >= 0 && np.x < len(grid[0]) && np.y < len(grid) && grid[np.y][np.x] != '#' {
					neighbors++
				}
			}
			if neighbors >= 3 {
				points = append(points, Point{y, x, 0})
			}
		}
	}
	return
}

func BuildGraph(grid []string, points []Point) {
	dirs := map[byte][]Point{
		'^': {{-1, 0, 0}},
		'v': {{1, 0, 0}},
		'<': {{0, -1, 0}},
		'>': {{0, 1, 0}},
		'.': {{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}},
		'#': {},
	}
	graph := map[Point]map[Point]int{}
	for _, point := range points {
		graph[point] = map[Point]int{}
	}
	for _, spoint := range points {
		stack := []Point{spoint}
		visited[Point{y: spoint.y, x: spoint.x, d: spoint.d}] = true
		for len(stack) > 0 {
			p := Pop(&stack)
			if p.d != 0 && visited[Point{p.y, p.x, 0}] {
				graph[Point{spoint.y, spoint.x, spoint.d}][Point{p.y, p.x, p.d}] = p.d
				continue
			}
			for _, dp := range dirs[grid[p.y][p.x]] {
				np := Point{y: p.y + dp.y, x: p.x + dp.x, d: 0}
				if np.x >= 0 && np.y >= 0 && np.x < len(grid[0]) && np.y < len(grid) && grid[np.y][np.x] != '#' && !visited[Point{y: np.y, x: np.x, d: 0}] {
					stack = append(stack, Point{y: np.y, x: np.x, d: spoint.d + 1})
					visited[np] = true
				}
			}
		}

	}
	for k, v := range graph {
		log.Printf("k=%+v v=%+v", k, v)
	}
	// log.Printf("graph=%+v", graph)
}

func Solve(input string, part int) (answer int) {
	grid := strings.Split(input, "\n")
	start, end := GetEndPoints(grid)
	decisionPoints := FindDecisionPoints(grid, start, end)
	BuildGraph(grid, decisionPoints)
	log.Printf("start, end = %+v %+v\npoints = %+v", start, end, decisionPoints)
	return
}
