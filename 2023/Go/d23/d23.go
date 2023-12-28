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

// type Point struct {
// 	y, x int
// }

// type Day struct {
// 	grid          []string
// 	start, end    Point
// 	height, width int
// 	crossroads    []Point
// 	neighbors     map[Point][]Point
// 	graph         map[Point][]GraphInfo
// }

// // var visited = make(map[Point]bool)

// func (d *Day) ReadGrid(input string) {
// 	d.grid = strings.Split(input, "\n")
// 	d.height, d.width = len(d.grid), len(d.grid[0])
// 	if d.height < 1 || d.width < 1 {
// 		log.Fatal("Could not read the grid.")
// 	}
// }

// func (d *Day) EndPoints() {
// 	d.start = Point{0, strings.Index(d.grid[0], ".")}
// 	d.end = Point{len(d.grid) - 1, strings.Index(d.grid[len(d.grid)-1], ".")}
// 	if d.start.x < 0 || d.end.x < 0 {
// 		log.Fatalf("Exit: start=%+v, end=%+v", d.start, d.end)
// 	}
// }

// func (d *Day) Prepare() {
// 	d.crossroads = []Point{d.start, d.end}
// 	d.neighbors = map[Point][]Point{}
// 	d.graph = map[Point][]GraphInfo{}
// }

// func (d *Day) ProcessCrossroadsAndNeighbors(part int) {
// 	var (
// 		count int
// 		coord Point
// 	)
// 	for y := 0; y < d.height; y++ {
// 		for x := 0; x < d.width; x++ {
// 			char := d.grid[y][x]
// 			if char != '#' {
// 				count = 0
// 				coord = Point{}
// 				if part == 1 {
// 					switch char {
// 					case 'v':
// 						coord = Point{y + 1, x}
// 					case '^':
// 						coord = Point{y - 1, x}
// 					case '>':
// 						coord = Point{y, x + 1}
// 					case '<':
// 						coord = Point{y, x - 1}
// 					default:
// 						for _, dir := range []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
// 							ny, nx := y+dir.y, x+dir.x
// 							if (ny >= 0 && ny < d.height && nx >= 0 && ny < d.width) && d.grid[ny][nx] != '#' {
// 								count++
// 								d.neighbors[Point{y, x}] = append(d.neighbors[Point{y, x}], Point{ny, nx})
// 							}
// 						}
// 					}
// 					if coord.y >= 0 && coord.x >= 0 && coord.y < d.height && coord.x < d.width {
// 						d.neighbors[Point{y, x}] = append(d.neighbors[Point{y, x}], coord)
// 					}
// 					if count > 2 {
// 						d.crossroads = append(d.crossroads, Point{y, x})
// 					}
// 				} else {
// 					count = 0
// 					for _, dir := range []Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
// 						ny, nx := y+dir.y, x+dir.x
// 						if (ny >= 0 && ny < d.height && nx >= 0 && ny < d.width) && d.grid[ny][nx] != '#' {
// 							count++
// 							d.neighbors[Point{y, x}] = append(d.neighbors[Point{y, x}], Point{ny, nx})
// 						}
// 					}
// 					if count > 2 {
// 						d.crossroads = append(d.crossroads, Point{y, x})
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func (d *Day) Graph() {
// 	for _, cr := range d.crossroads {
// 		for _, nbr := range d.neighbors[cr] {
// 			gi := CalcDist(d.crossroads, d.neighbors, nbr, 1, map[Point]bool{cr: true})
// 			d.graph[cr] = append(d.graph[cr], gi)
// 		}
// 	}
// }

// func CalcDist(crossroads []Point, neighbors map[Point][]Point, start Point, dist int, visited map[Point]bool) GraphInfo {
// 	for _, cr := range crossroads {
// 		if cr == start {
// 			return GraphInfo{start, dist}
// 		}
// 	}
// 	for _, nbr := range neighbors[start] {
// 		if !visited[nbr] {
// 			visited[nbr] = true
// 			return CalcDist(crossroads, neighbors, nbr, dist+1, visited)
// 		}
// 	}
// 	return GraphInfo{Point{}, 0}
// }

// type GraphInfo struct {
// 	p Point
// 	d int
// }

// func (d *Day) LongestPath() (answer int) {
// 	distances := BFS(d.graph, d.start, d.end, 0, map[Point]bool{d.start: true})
// 	for _, d := range distances {
// 		if d > answer {
// 			answer = d
// 		}
// 	}
// 	return
// }

// func BFS(graph map[Point][]GraphInfo, start, end Point, length int, visited map[Point]bool) []int {
// 	if start == end {
// 		return []int{length}
// 	}
// 	paths := []int{}
// 	for _, gi := range graph[start] {
// 		current := gi.p
// 		dist := gi.d
// 		if !visited[current] {
// 			visited[current] = true
// 			paths = append(paths, BFS(graph, current, end, length+dist, visited)...)
// 			delete(visited, current)
// 		}
// 	}
// 	return paths
// }

type Point struct {
	y, x int
}

type GraphPoint struct {
	dist, y, x int
}

func inGrid(grid []string, y, x int) bool {
	return y >= 0 && x >= 0 && y < len(grid) && x < len(grid[0])
}

func ThreeNeighborPoints(grid []string, start, end Point) (points []Point) {
	points = []Point{start, end}
	for y, row := range grid {
		for x, char := range strings.Split(row, "") {
			if char == "#" {
				continue
			}
			neighbors := 0
			for _, np := range [4]Point{{y, x - 1}, {y, x + 1}, {y - 1, x}, {y + 1, x}} {
				if inGrid(grid, np.y, np.x) && grid[np.y][np.x] != '#' {
					neighbors++
				}
			}
			if neighbors > 2 {
				points = append(points, Point{y, x})
			}
		}
	}
	return
}

func Pop(s *[]GraphPoint) GraphPoint {
	l := len(*s)
	e := (*s)[l-1]
	*s = (*s)[:l-1]
	return e
}

func BuildGraph(grid []string, points []Point) (g map[Point]map[Point]int) {
	dirs := map[byte][]Point{
		'^': {{-1, 0}},
		'v': {{1, 0}},
		'<': {{0, -1}},
		'>': {{0, 1}},
		'.': {{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
	}
	g = map[Point]map[Point]int{}
	stack := []GraphPoint{}
	visited := map[Point]bool{}
	for _, p := range points {
		g[p] = map[Point]int{}
	}

	for _, p := range points {
		sy, sx := p.y, p.x
		stack = []GraphPoint{{0, sy, sx}}
		visited[p] = true
		for len(stack) > 0 {
			plate := Pop(&stack)
			d, y, x := plate.dist, plate.y, plate.x
			if d != 0 && inPoints(points, Point{y, x}) {
				g[Point{sy, sx}][Point{y, x}] = d
				continue
			}
			for _, dir := range dirs[grid[y][x]] {
				ny := y + dir.y
				nx := x + dir.x
				if _, exists := visited[Point{ny, nx}]; !exists && inGrid(grid, ny, nx) && grid[ny][nx] != '#' {
					stack = append(stack, GraphPoint{d + 1, ny, nx})
					visited[Point{ny, nx}] = true
				}
			}
		}
	}
	return
}

func inPoints(points []Point, search Point) bool {
	for _, p := range points {
		if p.y == search.y && p.x == search.x {
			return true
		}
	}
	return false
}

var dfsVisited = make(map[Point]bool)
var dfsGraph = make(map[Point]map[Point]int)
var dfsEnd = Point{}

func DFS(point Point) int {
	if point.y == dfsEnd.y && point.x == dfsEnd.x {
		return 0
	}
	hi := -99
	dfsVisited[point] = true
	for np, d := range dfsGraph[point] {
		if _, exists := dfsVisited[np]; !exists {
			hi = max(hi, DFS(np)+d)
		}
	}
	delete(dfsVisited, point)
	log.Printf("hi=%d", hi)
	return hi
}

func Solve(input string, part int) (answer int) {
	grid := strings.Split(input, "\n")
	start := Point{0, strings.Index(grid[0], ".")}
	end := Point{len(grid) - 1, strings.Index(grid[len(grid)-1], ".")}
	points := ThreeNeighborPoints(grid, start, end)
	dfsGraph = BuildGraph(grid, points)
	dfsEnd = end
	answer = DFS(start)
	return
}
