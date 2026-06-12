package adventofcode

import "strings"

var y17d11Directions = map[string][2]float64{
	"n":  {0, 1},
	"s":  {0, -1},
	"ne": {0.5, 0.5},
	"se": {0.5, -0.5},
	"nw": {-0.5, 0.5},
	"sw": {-0.5, -0.5},
}

func y17d11(input string, part int) (answer int) {
	x, y, farthest := 0.0, 0.0, 0.0
	for move := range strings.SplitSeq(input, ",") {
		x += y17d11Directions[move][0]
		y += y17d11Directions[move][1]
		distance := absFloat(x) + absFloat(y)
		if distance > farthest {
			farthest = distance
		}
	}
	if part == 1 {
		answer = int(absFloat(x) + absFloat(y))
	} else {
		answer = int(farthest)
	}
	return
}
