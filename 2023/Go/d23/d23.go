package d23

import (
	"fmt"
	"log"
	"os"
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
	x, y int
}

var visited = make(map[Point]int)

func Endpoints(grid []string) (start, end Point) {
	start = Point{strings.Index(grid[0], "."), 0}
	end = Point{strings.Index(grid[len(grid)-1], "."), len(grid) - 1}
	return
}

func Hike(grid []string, start, facing Point) {
	// log.Printf("visited=%#v", visited)
	opMap := map[Point]byte{
		{1, 0}:  '<',
		{-1, 0}: '>',
		{0, 1}:  '^',
		{0, -1}: 'v',
	}
	current := start
	currStep := visited[current]
	for _, d := range [3]Point{facing, Left(facing), Right(facing)} {
		np := Point{current.x + d.x, current.y + d.y}
		if InGrid(grid, np) && grid[np.y][np.x] != '#' {
			if opMap[d] == grid[np.y][np.x] {
				continue
			}
			if cost, exists := visited[np]; !exists || cost < currStep+1 {
				if cost > np.y*np.x*200 {
					log.Printf("visited=%#v", visited)
					log.Print("WTF")
					os.Exit(1)
				}
				visited[np] = currStep + 1
				Hike(grid, np, d)
			}
		}
	}
}

func InGrid(grid []string, p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(grid[0]) && p.y < len(grid)
}
func Left(p Point) Point  { return Point{p.y, -p.x} }
func Right(p Point) Point { return Point{-p.y, p.x} }

func NoSlopes(grid []string) (ngrid []string) {
	ngrid = []string{}
	tmp := []string{}
	for _, line := range grid {
		for _, ch := range line {
			char := string(ch)
			switch char {
			case "#":
				tmp = append(tmp, char)
			case "^", "v":
				tmp = append(tmp, ".")
			default:
				tmp = append(tmp, ".")
			}
		}
		ngrid = append(ngrid, strings.Join(tmp, ""))
		tmp = []string{}
	}
	return
}

func Solve(input string, part int) (answer int) {
	visited = map[Point]int{}
	grid := strings.Split(input, "\n")
	start, end := Endpoints(grid)
	// log.Printf("start=%#v, end=%#v", start, end)
	visited[start] = 0
	if part == 1 {
		Hike(grid, start, Point{0, 1})
	} else {
		// grid = NoSlopes(grid)
		fmt.Print(strings.Join(grid, "\n"), "\n")
		Hike(grid, start, Point{0, 1})
	}
	answer = visited[end]
	return
}
