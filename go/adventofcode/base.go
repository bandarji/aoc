package adventofcode

import (
	"fmt"
	"os"
	"strings"
)

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

func NewAOCDay(year, day int) (DayRunner, error) {
	return nil, fmt.Errorf("no day runner for year %d, day %d", year, day)
}

func RunDay(dr DayRunner, year, day int) {
	fmt.Println(dr.Part1(year, day))
	fmt.Println(dr.Part2(year, day))
}

func readContent(filename string) string {
	if bytes, err := os.ReadFile(filename); err == nil {
		return strings.TrimSpace(string(bytes))
	}
	return ""
}

func formatFilename(year, day int) string {
	return fmt.Sprintf("%d-%02d-input.txt", year, day)
}
