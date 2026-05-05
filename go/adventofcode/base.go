package adventofcode

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const VOWELS string = "aeiou"
const ME string = "bandarji"

func inputFound(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}

func NewAOCDay(year, day int) (DayRunner, error) {
	concatDate := fmt.Sprintf("%d-%02d", year, day)
	if !inputFound(fmt.Sprintf("inputs/%s-input.txt", concatDate)) {
		return nil, fmt.Errorf("input not found for year %d, day %02d: skipping", year, day)
	}
	switch year {
	case 2015:
		return NewAOCDay2015(day)
	case 2016:
		return NewAOCDay2016(day)
	case 2017:
		return NewAOCDay2017(day)
	default:
		return nil, fmt.Errorf("no day runner for year %d, day %d", year, day)
	}
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
	return fmt.Sprintf("inputs/%d-%02d-input.txt", year, day)
}

func strToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
