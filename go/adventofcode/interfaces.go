package adventofcode

import "fmt"

type DayRunner interface {
	GetInput(year, day int) string
	Part1(year, day int) string
	Part2(year, day int) string
}

type AOCDayNoInput struct{}

func (d *AOCDayNoInput) GetInput() string {
	return ""
}

func (d *AOCDayNoInput) Part1(year, day int) string {
	return fmt.Sprintf("Year %d Day %02d Part 1: Not yet implemented", year, day)
}

func (d *AOCDayNoInput) Part2(year, day int) string {
	return fmt.Sprintf("Year %d Day %02d Part 2: Not yet implemented", year, day)
}
