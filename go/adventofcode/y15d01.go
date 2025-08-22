package adventofcode

import "fmt"

type Y15D01 struct{}

func (d *Y15D01) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D01) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d01p1(d.GetInput(year, day)))
}

func (d *Y15D01) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d01p2(d.GetInput(year, day)))
}

func y15d01p1(input string) (floor int) {
	for _, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return
}

func y15d01p2(input string) int {
	floor := 0
	for i, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return 0
}
