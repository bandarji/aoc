package d18

import (
	"aoc2023/common"
	"strconv"
	"strings"
)

const TEST1 string = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

type Instruction struct {
	dir, color string
	dist       int
}

type Point struct {
	y, x int
}

var dirs = map[string]Point{
	"U": {-1, 0},
	"D": {1, 0},
	"L": {0, -1},
	"R": {0, 1},
}

var dirs2 = map[byte]string{
	'0': "R",
	'1': "D",
	'2': "L",
	'3': "U",
}

func ReadPlan(plan []string) (instructions []Instruction) {
	instructions = []Instruction{}
	for _, entry := range plan {
		f := strings.Fields(entry)
		instructions = append(instructions, Instruction{f[0], f[2][2 : len(f[2])-1], common.StrToInt(f[1])})
	}
	return
}

func Dig(instructions []Instruction, part int) (answer int) {
	moves := 0
	points := []Point{{0, 0}}
	for _, ins := range instructions {
		dir, color, dist := ins.dir, ins.color, ins.dist
		_ = color
		delta := dirs[dir]
		if part == 2 {
			dist64, _ := strconv.ParseInt(color[0:5], 16, 64)
			dist = int(dist64)
			delta = dirs[dirs2[color[5]]]
			// log.Printf("color=%s %s %d", color, dirs2[color[5]], dist)
		}
		moves += dist
		y, x := points[len(points)-1].y, points[len(points)-1].x
		points = append(points, Point{y + (dist * delta.y), x + (dist * delta.x)})
	}
	inside := Area(points) - (moves / 2) + 1
	answer = inside + moves
	return
}

func Area(points []Point) (area int) {
	coords := len(points) - 1
	for i := range points {
		area += (points[coords].y + points[i].y) * (points[coords].x - points[i].x)
		coords = i
	}
	if area < 0 {
		area = -area
	}
	area /= 2
	return
}

func Solve(input string, part int) (answer int) {
	plan := strings.Split(input, "\n")
	instructions := ReadPlan(plan)
	answer = Dig(instructions, part)
	// log.Printf("ins=%+v", instructions)
	return
}
