package adventofcode

import (
	"crypto/md5"
	"fmt"
	"image"
)

const y16d17Input string = "mmsxrhfx"
const y16d17DestinationX int = 3
const y16d17DestinationY int = 3

type y16d17Room struct {
	point image.Point
	path  string
	steps int
}

func y16d17(input string, part int) string {
	longestPath := ""
	q := []y16d17Room{{point: image.Point{0, 0}, path: input, steps: 0}}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		if current.point.X == y16d17DestinationX && current.point.Y == y16d17DestinationY {
			rightPath := current.path[len(input):]
			if part == 1 {
				return rightPath
			}
			if len(longestPath) < len(rightPath) {
				longestPath = rightPath
			}
			continue
		}
		hash := fmt.Sprintf("%x", md5.Sum([]byte(current.path)))
		for i, dir := range [4]string{"U", "D", "L", "R"} {
			nr, nc := 0, 0
			switch dir {
			case "U":
				nr = current.point.X - 1
				nc = current.point.Y
			case "D":
				nr = current.point.X + 1
				nc = current.point.Y
			case "L":
				nr = current.point.X
				nc = current.point.Y - 1
			case "R":
				nr = current.point.X
				nc = current.point.Y + 1
			}
			if nr >= 0 && nr < 4 && nc >= 0 && nc < 4 {
				if hash[i] >= 'b' && hash[i] <= 'f' {
					q = append(q, y16d17Room{point: image.Point{nr, nc}, path: current.path + dir, steps: current.steps + 1})
				}
			}
		}
	}
	return fmt.Sprintf("%d", len(longestPath))
}
