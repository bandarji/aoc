package adventofcode

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const y15d04input string = "yzbqklnj"

type Y15D04 struct{}

func (d *Y15D04) GetInput(year, day int) string {
	return y15d04input
}

func (d *Y15D04) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d04(d.GetInput(year, day), 5))
}

func (d *Y15D04) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d04(d.GetInput(year, day), 6))
}

func y15d04(input string, zeroes int) (number int) {
	for {
		number++
		data := []byte(fmt.Sprintf("%s%d", input, number))
		s := fmt.Sprintf("%x", md5.Sum(data))
		if strings.Repeat("0", zeroes) == s[:zeroes] {
			break
		}
	}
	return
}
