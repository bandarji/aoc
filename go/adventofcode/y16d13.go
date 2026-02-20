package adventofcode

import "strconv"

const (
	y16d13DestinationX int = 31
	y16d13DestinationY int = 39
)

func y16d13(input string, destinationX, destinationY, part int) (answer int) {
	favoriteNumber := strToInt(input)
	answer = y16d13CubiclesMaze(favoriteNumber, destinationX, destinationY, part)
	return
}

func y16d13CubiclesMaze(favoriteNumber, destinationX, destinationY, part int) int {
	visited := map[[2]int]bool{}
	visited2 := map[[2]int]bool{}
	q := [][3]int{{1, 1, 0}} // coordinates and distance
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		x, y, dist := current[0], current[1], current[2]
		if part == 1 && x == destinationX && y == destinationY {
			return dist
		}
		if !visited[[2]int{x, y}] {
			if dist <= 50 {
				visited2[[2]int{x, y}] = true
			}
			if part == 2 && dist > 50 {
				return len(visited2)
			}
			for _, diff := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nx, ny := x+diff[0], y+diff[1]
				if nx >= 0 && ny >= 0 && isOpen(nx, ny, favoriteNumber) {
					q = append(q, [3]int{nx, ny, dist + 1})
				}
			}
		}
		visited[[2]int{x, y}] = true
	}
	return -1
}

func isOpen(x, y, favoriteNumber int) bool {
	ones := 0
	num := (x * x) + (3 * x) + (2 * x * y) + y + (y * y) + favoriteNumber
	for _, char := range strconv.FormatInt(int64(num), 2) {
		if char == '1' {
			ones++
		}
	}
	return ones%2 == 0
}
