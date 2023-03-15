package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	Part, Grains, Floor, Minx, Maxx int
	Start                           [2]int
	Mapping                         map[[2]int]string
}

const EXTEND int = 10000
const TEST_INPUT string = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`
const TEST_ANS1 int = 24
const TEST_ANS2 int = 93

func (g *Grid) Render() {
	lowx, highx := 0, 0
	if g.Part == 1 {
		lowx, highx = g.Minx-2, g.Maxx+2
	} else {
		lowx, highx = g.Minx-EXTEND, g.Maxx+EXTEND
	}
	for y := 0; y <= g.Floor+2; y++ {
		for x := lowx; x < highx; x++ {
			if x == g.Start[0] && y == g.Start[1] {
				fmt.Print("+")
			} else if Mapped(g, x, y) {
				fmt.Print(g.Mapping[[2]int{x, y}])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (g *Grid) ConstructCave(s string) {
	x, y := 0, 0
	points := make([][2]int, 0)
	for _, rawCoord := range strings.Split(s, " -> ") {
		rawPoint := strings.Split(rawCoord, ",")
		x, y = atoi(rawPoint[0]), atoi(rawPoint[1])
		if x < g.Minx {
			g.Minx = x
		}
		if x > g.Maxx {
			g.Maxx = x
		}
		points = append(points, [2]int{x, y})
	}
	x, y = points[0][0], points[0][1]
	if y > g.Floor {
		g.Floor = y
	}
	g.Mapping[[2]int{x, y}] = "#"
	for i := 1; i < len(points); i++ {
		nx, ny := points[i][0], points[i][1]
		for x != nx || y != ny {
			dx, dy := 0, 0
			if nx > x {
				dx = 1
			}
			if x > nx {
				dx -= 1
			}
			if ny > y {
				dy = 1
			}
			if y > ny {
				dy -= 1
			}
			x += dx
			y += dy
			g.Mapping[[2]int{x, y}] = "#"
			if y > g.Floor {
				g.Floor = y
			}
		}
	}
}

func (g *Grid) PourSand() {
	if g.Part == 1 {
		for v := 0; v < g.Floor*(g.Maxx-g.Minx)+1; v++ {
			x, y := g.Start[0], g.Start[1]
			for {
				// g.Render()
				if y > g.Floor {
					break
				}
				if !Mapped(g, x, y+1) {
					y += 1
				} else if !Mapped(g, x-1, y+1) {
					x -= 1
					y += 1
				} else if !Mapped(g, x+1, y+1) {
					x += 1
					y += 1
				} else {
					g.Mapping[[2]int{x, y}] = "o"
					break
				}
			}
		}
	} else {
		for {
			x, y := g.Start[0], g.Start[1]
			for {
				// g.Render()
				if y > g.Floor {
					break
				}
				if !Mapped(g, x, y+1) {
					y += 1
				} else if !Mapped(g, x-1, y+1) {
					x -= 1
					y += 1
				} else if !Mapped(g, x+1, y+1) {
					x += 1
					y += 1
				} else {
					g.Mapping[[2]int{x, y}] = "o"
					break
				}
			}
			cell, ok := g.Mapping[[2]int{g.Start[0], g.Start[1]}]
			if ok && cell == "o" {
				break
			}
		}
	}
}

func (g *Grid) CountGrains() {
	for _, cell := range g.Mapping {
		if cell == "o" {
			g.Grains++
		}
	}
}

func (g *Grid) AddSubFloor() {
	g.Floor += 2
	for x := g.Minx - EXTEND; x <= g.Maxx+EXTEND; x++ {
		g.Mapping[[2]int{x, g.Floor}] = "#"
	}
}

func Mapped(grid *Grid, px, py int) bool {
	_, ok := grid.Mapping[[2]int{px, py}]
	if ok {
		return true
	} else {
		return false
	}
}

func Day14(input string, part int) int {
	grid := &Grid{part, 0, 0, 1 << 15, -1 * 1 << 15, [2]int{500, 0}, make(map[[2]int]string)}
	input = strings.TrimRight(input, "\n")
	for _, line := range strings.Split(input, "\n") {
		grid.ConstructCave(line)
	}
	if grid.Part == 2 {
		grid.AddSubFloor()
	}
	grid.PourSand()
	grid.CountGrains()
	// log.Println(grid.Floor, grid.Grains)
	return grid.Grains
}

func RunTests(part int) bool {
	if part == 1 {
		return Day14(TEST_INPUT, part) == TEST_ANS1
	} else {
		return Day14(TEST_INPUT, part) == TEST_ANS2
	}
}

func atoi(s string) int {
	ns := "" // number string
	for i := 0; i < len(s); i++ {
		if s[i] <= 57 && s[i] >= 48 {
			ns += string(s[i])
		}
	}
	num, _ := strconv.Atoi(ns)
	return num
}

func main() {
	day, part := 14, 1
	if !RunTests(part) {
		log.Fatal("Day", day, "part", part, "failed")
	}
	part = 2
	if !RunTests(part) {
		log.Fatal("Day", day, "part", part, "failed")
	}
	content, _ := os.ReadFile("input_day14.txt")
	grains := Day14(string(content), 1)
	fmt.Println(grains)
	grains = Day14(string(content), 2)
	fmt.Println(grains)
}
