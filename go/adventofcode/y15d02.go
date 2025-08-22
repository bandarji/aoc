package adventofcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Y15D02 struct{}

func (d *Y15D02) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D02) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d02p1(d.GetInput(year, day)))
}

func (d *Y15D02) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d02p2(d.GetInput(year, day)))
}

func y15d02p1(input string) (sqft int) {
	sqft = y15d02solve(input, 1)
	return
}

func y15d02p2(input string) (ft int) {
	ft = y15d02solve(input, 2)
	return
}

func y15d02solve(input string, part int) (result int) {
	input = strings.TrimRight(input, "\n")
	dims := make([]int, 3)
	for _, line := range strings.Split(input, "\n") {
		for i, dim := range strings.Split(line, "x") {
			value, _ := strconv.Atoi(dim)
			dims[i] = value
		}
		sort.Ints(dims)
		if part == 1 {
			result += 2*dims[0]*dims[1] + 2*dims[1]*dims[2] + 2*dims[0]*dims[2]
			result += dims[0] * dims[1]
		} else {
			result += 2*dims[0] + 2*dims[1]
			result += dims[0] * dims[1] * dims[2]
		}
	}
	return
}
