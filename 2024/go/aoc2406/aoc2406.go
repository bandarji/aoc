package aoc2406

import (
	"strings"
)

const MAXMOVES int = 10_000
const AOC2406_TEST string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

type Map struct {
	Grid         [][]string
	X, Y, Facing int // 0 = N, 1 = E, 2 = S, 3 = W
	Visited      map[[2]int]bool
}

func NewMap(input string) *Map {
	m := &Map{
		Visited: map[[2]int]bool{},
		Facing:  0,
	}
	for _, line := range strings.Split(input, "\n") {
		m.Grid = append(m.Grid, strings.Split(line, ""))
	}
	for y, row := range m.Grid {
		for x, cell := range row {
			if cell == "^" {
				m.X, m.Y = x, y
				m.Visited[[2]int{x, y}] = true
				break
			}
		}
	}
	return m
}

func (m *Map) PatrolLoopCheck(x, y int) bool {
	moves := 0
	for {
		moves++
		if moves > MAXMOVES {
			return true
		}
		switch m.Facing {
		case 0:
			if m.Y == 0 {
				return false
			}
			if m.Grid[m.Y-1][m.X] == "#" {
				m.Facing++
			} else {
				m.Y--
			}
		case 1:
			if m.X == len(m.Grid[0])-1 {
				return false
			}
			if m.Grid[m.Y][m.X+1] == "#" {
				m.Facing++
			} else {
				m.X++
			}
		case 2:
			if m.Y == len(m.Grid)-1 {
				return false
			}
			if m.Grid[m.Y+1][m.X] == "#" {
				m.Facing++
			} else {
				m.Y++
			}
		case 3:
			if m.X == 0 {
				return false
			}
			if m.Grid[m.Y][m.X-1] == "#" {
				m.Facing = 0
			} else {
				m.X--
			}
		}
		m.Visited[[2]int{m.X, m.Y}] = true
	}
}

func (m *Map) Patrol() {
	for {
		// fmt.Printf("X: %d, Y: %d, Facing: %d\n", m.X, m.Y, m.Facing)
		switch m.Facing {
		case 0:
			if m.Y == 0 {
				return
			}
			if m.Grid[m.Y-1][m.X] == "#" {
				m.Facing++
			} else {
				m.Y--
			}
		case 1:
			if m.X == len(m.Grid[0])-1 {
				return
			}
			if m.Grid[m.Y][m.X+1] == "#" {
				m.Facing++
			} else {
				m.X++
			}
		case 2:
			if m.Y == len(m.Grid)-1 {
				return
			}
			if m.Grid[m.Y+1][m.X] == "#" {
				m.Facing++
			} else {
				m.Y++
			}
		case 3:
			if m.X == 0 {
				return
			}
			if m.Grid[m.Y][m.X-1] == "#" {
				m.Facing = 0
			} else {
				m.X--
			}
		}
		m.Visited[[2]int{m.X, m.Y}] = true
	}
}

func Aoc240601(input string) (count int) {
	m := NewMap(input)
	m.Patrol()
	count = len(m.Visited)
	return
}

func Aoc240602(input string) (loops int) {
	m := NewMap(input)
	for y, row := range m.Grid {
		for x := range row {
			n := NewMap(input)
			n.Grid[y][x] = "#"
			if n.PatrolLoopCheck(x, y) {
				loops++
			}
		}
	}
	return
}
