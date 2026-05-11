package d02

import (
	"aoc2023/common"
	"strings"
)

type Counts struct {
	Red, Green, Blue int
}

const TEST1 string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func GetGameID(s string) (id int) {
	fields := strings.Fields(s)
	id = common.StrToInt(strings.Trim(fields[1], ":"))
	return
}

func PossibleGame(s string, maxRed, maxGreen, maxBlue int) bool {
	var counts Counts
	var count, start int
	for _, grab := range strings.Split(s, ";") {
		counts = Counts{0, 0, 0}
		fields := strings.Fields(grab)
		if strings.Contains(grab, ":") {
			start = 2
		} else {
			start = 0
		}
		for i := start; i < len(fields); i += 2 {
			count = common.StrToInt(fields[i])
			switch {
			case strings.HasPrefix(fields[i+1], "red"):
				counts.Red += count
			case strings.HasPrefix(fields[i+1], "green"):
				counts.Green += count
			case strings.HasPrefix(fields[i+1], "blue"):
				counts.Blue += count
			}
		}
		if counts.Red > maxRed || counts.Green > maxGreen || counts.Blue > maxBlue {
			return false
		}
	}
	return true
}

func Power(s string) (power int) {
	var counts Counts
	var count, start int
	maxCounts := Counts{0, 0, 0}
	for _, grab := range strings.Split(s, ";") {
		counts = Counts{0, 0, 0}
		fields := strings.Fields(grab)
		if strings.Contains(grab, ":") {
			start = 2
		} else {
			start = 0
		}
		for i := start; i < len(fields); i += 2 {
			count = common.StrToInt(fields[i])
			switch {
			case strings.HasPrefix(fields[i+1], "red"):
				counts.Red += count
			case strings.HasPrefix(fields[i+1], "green"):
				counts.Green += count
			case strings.HasPrefix(fields[i+1], "blue"):
				counts.Blue += count
			}
		}
		if counts.Red > maxCounts.Red {
			maxCounts.Red = counts.Red
		}
		if counts.Green > maxCounts.Green {
			maxCounts.Green = counts.Green
		}
		if counts.Blue > maxCounts.Blue {
			maxCounts.Blue = counts.Blue
		}
	}
	power = maxCounts.Red * maxCounts.Green * maxCounts.Blue
	return
}

func Solve(input string, part int) (total int) {
	for _, line := range strings.Split(input, "\n") {
		if part == 1 {
			if PossibleGame(line, 12, 13, 14) {
				total += GetGameID(line)
			}
		} else {
			total += Power(line)
		}
	}
	return
}
