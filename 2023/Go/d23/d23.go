package d23

import (
	"fmt"
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

type Point struct {
	y, x int
}

type Place struct {
	d int
	p Point
}

func InGrid(grid *[]string, y, x int) bool {
	return y >= 0 && y < len(*grid) && x >= 0 && x < len((*grid)[0])
}

func LongestHike(grid []string, start, end Point) (answer int) {
	steps := []int{}
	visited := map[Point]bool{}
	dirs := map[byte][]Point{
		'^': {{-1, 0}},
		'v': {{1, 0}},
		'<': {{0, -1}},
		'>': {{0, 1}},
		'.': {{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		'#': {},
	}
	q := []Place{{0, start}}
	for len(q) > 0 {
		place := q[0]
		q = q[1:]
		d, sy, sx := place.d, place.p.y, place.p.x
		for _, dp := range dirs[grid[sy][sx]] {
			ny, nx := sy+dp.y, sx+dp.x
			if !InGrid(&grid, ny, nx) || grid[ny][nx] == '#' || visited[Point{ny, nx}] {
				continue
			}
			DisplayGrid(grid, visited)
			q = append(q, Place{d + 1, Point{ny, nx}})
			if ny == end.y && nx == end.x {
				steps = append(steps, d+1)
			}
		}
		visited[Point{sy, sx}] = true
	}
	log.Printf("steps=%+v", steps)
	return
}

func DisplayGrid(grid []string, visited map[Point]bool) {
	fmt.Println("")
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if visited[Point{y, x}] {
				fmt.Print("O")
			} else {
				fmt.Print(string(grid[y][x]))
			}
		}
		fmt.Println("")
	}
	log.Print("")
}

func Solve(input string, part int) (answer int) {
	grid := strings.Split(input, "\n")
	start := Point{0, strings.Index(grid[0], ".")}
	end := Point{len(grid) - 1, strings.Index(grid[len(grid)-1], ".")}
	answer = LongestHike(grid, start, end)
	return
}
