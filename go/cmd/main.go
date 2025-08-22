package main

import (
	"github.com/bandarji/aoc/adventofcode"
)

const (
	firstYear   int = 2015
	currentYear int = 2024
)

func main() {
	for year := firstYear; year <= currentYear; year++ {
		for day := 1; day <= 25; day++ {
			if data, err := adventofcode.NewAOCDay(year, day); err == nil {
				adventofcode.RunDay(data, year, day)
			}
		}
	}
}
