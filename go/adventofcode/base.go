package adventofcode

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const VOWELS string = "aeiou"
const ME string = "bandarji"

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
	switch concatDate {
	case "2015-01":
		return &Y15D01{}, nil
	case "2015-02":
		return &Y15D02{}, nil
	case "2015-03":
		return &Y15D03{}, nil
	case "2015-04":
		return &Y15D04{}, nil
	case "2015-05":
		return &Y15D05{}, nil
	case "2015-06":
		return &Y15D06{}, nil
	case "2015-07":
		return &Y15D07{}, nil
	case "2015-08":
		return &Y15D08{}, nil
	case "2015-09":
		return &Y15D09{}, nil
	case "2015-10":
		return &Y15D10{}, nil
	case "2015-11":
		return &Y15D11{}, nil
	case "2015-12":
		return &Y15D12{}, nil
	case "2015-13":
		return &Y15D13{}, nil
	case "2015-14":
		return &Y15D14{}, nil
	case "2015-15":
		return &Y15D15{}, nil
	case "2015-16":
		return &Y15D16{}, nil
	case "2015-17":
		return &Y15D17{}, nil
	case "2015-18":
		return &Y15D18{}, nil
	case "2015-19":
		return &Y15D19{}, nil
	case "2015-20":
		return &Y15D20{}, nil
	case "2015-21":
		return &Y15D21{}, nil
	case "2015-22":
		return &Y15D22{}, nil
	case "2015-23":
		return &Y15D23{}, nil
	case "2015-24":
		return &Y15D24{}, nil
	case "2015-25":
		return &Y15D25{}, nil
	case "2016-01":
		return &Y16D01{}, nil
	case "2016-02":
		return &Y16D02{}, nil
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
