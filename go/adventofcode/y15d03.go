package adventofcode

import "fmt"

type y15d03coords map[string]int

type Y15D03 struct{}

func (d *Y15D03) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D03) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d03p1(d.GetInput(year, day)))
}

func (d *Y15D03) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d03p2(d.GetInput(year, day)))
}

func y15d03p1(input string) (houses int) {
	coords := y15d03coords{"0-0": 1}
	visitHouses(&coords, input)
	houses = len(coords)
	return
}

func y15d03p2(input string) (houses int) {
	coords := y15d03coords{"0-0": 1}
	visitHousesWithRoboSanta(&coords, input)
	houses = len(coords)
	return
}

func visitHouses(coords *y15d03coords, input string) {
	x, y := 0, 0
	dx, dy := 0, 0
	for _, char := range input {
		dx, dy = 0, 0
		switch string(char) {
		case "^":
			dy += 1
		case "v":
			dy -= 1
		case ">":
			dx += 1
		case "<":
			dx -= 1
		}
		x += dx
		y += dy
		(*coords)[fmt.Sprintf("%d-%d", x, y)] = 1
	}
}

func visitHousesWithRoboSanta(coords *y15d03coords, input string) {
	sx, sy, rx, ry := 0, 0, 0, 0
	dx, dy := 0, 0
	for i, char := range input {
		dx, dy = 0, 0
		switch string(char) {
		case "^":
			dy += 1
		case "v":
			dy -= 1
		case ">":
			dx += 1
		case "<":
			dx -= 1
		}
		if i%2 == 0 {
			sx += dx
			sy += dy
			(*coords)[fmt.Sprintf("%d-%d", sx, sy)] = 1
		} else {
			rx += dx
			ry += dy
			(*coords)[fmt.Sprintf("%d-%d", rx, ry)] = 2
		}
	}
}
