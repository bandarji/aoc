package adventofcode

import (
	"fmt"
	"math"
)

const y15d20Input string = "34000000"

type Y15D20 struct{}

func (d *Y15D20) GetInput(year, day int) string {
	return y15d20Input
}

func (d *Y15D20) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d20(d.GetInput(year, day), 1))
}

func (d *Y15D20) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d20(d.GetInput(year, day), 2))
}

func y15d20(input string, part int) (house int) {
	leastPresents := strToInt(input)
	for house = 1; house < 2_000_000_000; house++ {
		var presents int
		for _, factor := range y15d20Factors(house) {
			switch part {
			case 1:
				presents += factor * 10
			case 2:
				if house/factor <= 50 {
					presents += factor * 11
				}
			}
		}
		if presents >= leastPresents {
			break
		}
	}
	return
}

func y15d20Factors(n int) (factors []int) {
	highest := int(math.Sqrt(float64(n))) + 1
	for i := 1; i <= highest; i++ {
		if n%i == 0 {
			factors = append(factors, i, n/i)
		}
	}
	return
}
