package d05

import (
	"aoc2023/common"
	"fmt"
	"strings"
)

const TEST1 string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func GetSeeds(input string) (seeds []int) {
	seeds = []int{}
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "seeds: ") {
			sections := strings.Split(line, ":")
			for _, entry := range strings.Fields(sections[1]) {
				seeds = append(seeds, common.StrToInt(entry))
			}
			return seeds
		}
	}
	return
}

type Map struct {
	Source, Destination, Range int
}

type Maps struct {
	Mapping []Map
}

type Almanac struct {
	AllMaps []Maps
}

func (m Maps) FindMapping(src int) int {
	for _, oneMap := range m.Mapping {
		if src < oneMap.Source+oneMap.Range && src >= oneMap.Source {
			return oneMap.Destination + src - oneMap.Source
		}
	}
	return src
}

func (a Almanac) FindMapping(src int) int {
	for _, maps := range a.AllMaps {
		src = maps.FindMapping(src)
	}
	return src
}

func ParseMapping(lines string) (m Maps) {
	var dst, src, length int
	for i, line := range strings.Split(lines, "\n") {
		// skip words line
		if i == 0 {
			continue
		}
		fmt.Sscanf(line, "%d %d %d", &dst, &src, &length)
		m.Mapping = append(m.Mapping, Map{Source: src, Destination: dst, Range: length})
	}
	return
}

func BuildMaps(input string) (a Almanac) {
	a = Almanac{}
	sections := strings.Split(input, "\n\n")
	for i := 1; i < len(sections); i++ {
		a.AllMaps = append(a.AllMaps, ParseMapping(sections[i]))
	}
	return
}

func Part1(seeds []int, almanac Almanac) (answer int) {
	var minimum int = 9223372036854775807
	var location int
	for _, seed := range seeds {
		location = almanac.FindMapping(seed)
		if location < minimum {
			minimum = location
		}
	}
	answer = minimum
	return
}

func Part2(seeds []int, almanac Almanac) (answer int) {
	var minimum int = 9223372036854775807
	var location int
	for i := 0; i < len(seeds); i += 2 {
		for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
			location = almanac.FindMapping(seed)
			if location < minimum {
				minimum = location
			}
		}
	}
	answer = minimum
	return
}

func Solve(input string, part int) (answer int) {
	seeds := GetSeeds(input)
	almanac := BuildMaps(input)
	if part == 1 {
		answer = Part1(seeds, almanac)
	} else {
		answer = Part2(seeds, almanac)
	}
	return
}
