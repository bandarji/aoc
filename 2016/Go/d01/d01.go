package d01

import (
	"fmt"
	"strings"
)

var Directions = [4]int{0, 1, 2, 3}
var Deltas = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func NormalizePosition(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func ParseInstruction(instruction string) (turn string, steps int) {
	fmt.Sscanf(instruction, "%1s%d", &turn, &steps)
	return
}

func Turn(turn string, facing int) (direction int) {
	delta := 0
	if turn == "R" {
		delta = 1
	} else if turn == "L" {
		delta = 3
	}
	direction = Directions[(facing+delta)%4]
	return
}

func Distance(sx, sy, ex, ey int) int {
	return Abs(sx, ex) + Abs(sy, ey)
}

func Abs(v1, v2 int) (a int) {
	a = v2 - v1
	if a < 0 {
		a *= -1
	}
	return
}

func Day(input string, part int) int {
	visited := map[string]bool{}
	facing := 0 // 0 -> N, 1 -> E, 2 -> S, 3 -> W
	x, y := 0, 0
	position := NormalizePosition(x, y)
	for _, instruction := range strings.Split(input, ", ") {
		turn, steps := ParseInstruction(instruction)
		facing = Turn(turn, facing)
		for i := 0; i < steps; i++ {
			x += Deltas[facing][0]
			y += Deltas[facing][1]
			position = NormalizePosition(x, y)
			if part == 2 && visited[position] {
				return Distance(0, 0, x, y)
			}
			visited[position] = true
		}
	}
	return Distance(0, 0, x, y)
}
