package adventofcode

import (
	"fmt"
	"strings"
)

var (
	y16d01Directions = [4]int{0, 1, 2, 3}
	y16d01Deltas     = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func y16d01(input string, part int) int {
	visited := map[string]bool{}
	facing, x, y, steps := 0, 0, 0, 0
	position, turn := "", ""
	for _, ins := range strings.Split(input, ", ") {
		fmt.Sscanf(ins, "%1s%d", &turn, &steps)
		facing = y16d01Turn(turn, facing)
		for i := 0; i < steps; i++ {
			x += y16d01Deltas[facing][0]
			y += y16d01Deltas[facing][1]
			position = fmt.Sprintf("%d,%d", x, y)
			if part == 2 && visited[position] {
				return y16d01Abs(0, x) + y16d01Abs(0, y)
			}
			visited[position] = true
		}
	}
	return y16d01Abs(0, x) + y16d01Abs(0, y)
}

func y16d01Abs(x, y int) (a int) {
	a = x - y
	if a < 0 {
		a *= -1
	}
	return
}

func y16d01Turn(turn string, facing int) (dir int) {
	delta := 0
	switch turn {
	case "L":
		delta = 1
	case "R":
		delta = 3
	}
	dir = y16d01Directions[(facing+delta)%4]
	return
}
