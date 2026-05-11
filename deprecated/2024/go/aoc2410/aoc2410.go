package aoc2410

import (
	"aoc24/aoc2400"
	"fmt"
	"strings"
	"time"
)

const AOC2410_TEST string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

type Location struct {
	X, Y, H int
	V       bool
}

type Heights map[Location]int
type TrailHeads []Location

func (l Location) Next(topo [][]int) (dirs []string) {
	dirs = []string{}
	if l.X > 0 && topo[l.Y][l.X-1]-1 == l.H {
		dirs = append(dirs, "left")
	}
	if l.X+1 < len(topo[0]) && topo[l.Y][l.X+1]-1 == l.H {
		dirs = append(dirs, "right")
	}
	if l.Y > 0 && topo[l.Y-1][l.X]-1 == l.H {
		dirs = append(dirs, "up")
	}
	if l.Y+1 < len(topo) && topo[l.Y+1][l.X]-1 == l.H {
		dirs = append(dirs, "down")
	}
	return
}

func (l Location) Move(dir string) (nextLoc Location) {
	nextLoc = Location{}
	nH := l.H + 1
	switch dir {
	case "left":
		nextLoc = Location{X: l.X - 1, Y: l.Y, H: nH}
	case "right":
		nextLoc = Location{X: l.X + 1, Y: l.Y, H: nH}
	case "up":
		nextLoc = Location{X: l.X, Y: l.Y - 1, H: nH}
	case "down":
		nextLoc = Location{X: l.X, Y: l.Y + 1, H: nH}
	}
	if nextLoc.H == 9 {
		nextLoc.V = true
	}
	return
}

func (l Location) Paths(topo [][]int, heights Heights) {
	for _, dir := range l.Next(topo) {
		loc := l.Move(dir)
		if loc.V {
			heights[loc]++
		}
		loc.Paths(topo, heights)
	}
}

func ReadMap(input string) (topo [][]int, theads TrailHeads) {
	topo = [][]int{}
	theads = TrailHeads{}
	for y, line := range strings.Split(input, "\n") {
		row := []int{}
		for x, c := range strings.Split(line, "") {
			val := aoc2400.StrToInt(c)
			row = append(row, val)
			if val == 0 {
				theads = append(
					theads,
					Location{X: x, Y: y, H: val},
				)
			}
		}
		topo = append(topo, row)
	}
	return
}

func Aoc241001(input string) (score int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	topo, theads := ReadMap(input)
	for _, loc := range theads {
		heights := Heights{}
		loc.Paths(topo, heights)
		score += len(heights)
	}
	return
}

func Aoc241002(input string) (score int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	topo, theads := ReadMap(input)
	for _, loc := range theads {
		heights := Heights{}
		loc.Paths(topo, heights)
		for _, v := range heights {
			score += v
		}
	}
	return
}
