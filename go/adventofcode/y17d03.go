package adventofcode

import (
	"image"
	"log"
)

func y17d03(input string, part int) (dist int) {
	if part == 1 {
		dist = y17d03Solve1(strToInt(input))
	}
	return
}

func y17d03Solve1(iNumber int) int {
	log.Println("iNumber", iNumber)
	dirs := []image.Point{image.Point{X: 0, Y: 1}, image.Point{X: -1, Y: 0}, image.Point{X: 0, Y: -1}, image.Point{X: 1, Y: 0}} // R, U, L, D
	d, r, c, n := 0, 0, 0, 2
	getToSquaredNum := map[image.Point]int{image.Point{0, 0}: 1, image.Point{0, 1}: 2}
	for len(getToSquaredNum) < iNumber {
		dl := dirs[(d+1)%4]
		cl := image.Point{Y: r + dl.Y, X: c + dl.X}
		if _, exists := getToSquaredNum[cl]; !exists {
			d = (d + 1) % 4
		}
		dp := dirs[d]
		r += dp.Y
		c += dp.X
		np := image.Point{Y: r, X: c}
		getToSquaredNum[np] = n
		n++
	}
	return y17d03Manhattan(0, 0, r, c)
}

func y17d03Manhattan(x1, y1, x2, y2 int) int {
	dx := x1 - x2
	dy := y1 - y2
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}
	return dx + dy
}
