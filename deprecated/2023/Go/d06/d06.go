package d06

import (
	"aoc2023/common"
	"strings"
)

const TEST1 string = `Time:      7  15   30
Distance:  9  40  200`

func ParseRaces(input string) (races map[int]int) {
	races = map[int]int{}
	times := []int{}
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		if strings.HasPrefix(line, "Time:") {
			for _, time := range fields[1:] {
				times = append(times, common.StrToInt(time))
			}
		}
		if strings.HasPrefix(line, "Distance:") {
			for i, dist := range fields[1:] {
				races[times[i]] = common.StrToInt(dist)
			}
		}
	}
	return
}

func ParseRaces2(input string) (races map[int]int) {
	races = map[int]int{}
	time := 0
	dist := 0
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		if strings.HasPrefix(line, "Time:") {
			time = common.StrToInt(strings.Join(fields[1:], ""))
		}
		if strings.HasPrefix(line, "Distance:") {
			dist = common.StrToInt(strings.Join(fields[1:], ""))
		}
	}
	races[time] = dist
	return
}

func Race(time, win int) (ways int) {
	for t := 1; t < time-1; t++ {
		if (time-t)*t > win {
			ways++
		}
	}
	return
}

func Solve(input string, part int) (answer int) {
	var races map[int]int
	answer++
	if part == 1 {
		races = ParseRaces(input)
	} else {
		races = ParseRaces2(input)
	}
	for time, distance := range races {
		answer *= Race(time, distance)
	}
	return
}
