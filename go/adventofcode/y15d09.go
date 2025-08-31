package adventofcode

import (
	"fmt"
	"strconv"
	"strings"
)

type Y15D09 struct{}

func (d *Y15D09) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D09) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d09(d.GetInput(year, day), 1))
}

func (d *Y15D09) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d09(d.GetInput(year, day), 2))
}

func y15d09(input string, part int) int {
	places := y15d09AssemblePlaces(input)
	perms := [][]string{}
	y15d09Permutations(places, len(places), &perms)
	short, long := y15d09CalcDistances(perms, y15d09AssembleDistances(input), part)
	if part == 1 {
		return short
	}
	return long
}

func y15d09CalcDistances(perms [][]string, distances map[string]int, part int) (short, long int) {
	short = 1 << 31
	for _, perm := range perms {
		total := 0
		for i := 1; i < len(perm); i++ {
			from, to := perm[i-1], perm[i]
			dist := y15d09FindDistance(distances, from, to)
			if part == 1 {
				if dist < 0 {
					total += 100000
				} else {
					total += dist
				}
			} else {
				if dist < 0 {
					total -= 100000
				} else {
					total += dist
				}
			}
		}
		if part == 1 && total < short {
			short = total
		} else if part == 2 && total > long {
			long = total
		}
	}
	return
}

func y15d09FindDistance(distances map[string]int, from, to string) int {
	if dist, ok := distances[y15d09EncKey(from, to)]; ok {
		return dist
	}
	return -1
}

func y15d09AssembleDistances(input string) map[string]int {
	distances := map[string]int{}
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		fields := strings.Fields(line)
		dist, _ := strconv.Atoi(fields[4])
		key1 := y15d09EncKey(fields[0], fields[2])
		key2 := y15d09EncKey(fields[2], fields[0])
		distances[key1] = dist
		distances[key2] = dist
	}
	return distances
}

func y15d09EncKey(p1, p2 string) string {
	return fmt.Sprintf("%s-%s", p1, p2)
}

func y15d09AssemblePlaces(input string) (spots []string) {
	places := map[string]bool{}
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		fields := strings.Fields(line)
		places[fields[0]] = true
		places[fields[2]] = true
	}
	for place := range places {
		spots = append(spots, place)
	}
	return
}

func y15d09Permutations(places []string, n int, result *[][]string) {
	if n == 1 {
		dst := make([]string, len(places))
		copy(dst, places[:])
		*result = append(*result, dst)
	} else {
		for i := 0; i < n; i++ {
			y15d09Permutations(places, n-1, result)
			if n%2 == 0 {
				places[0], places[n-1] = places[n-1], places[0]
			} else {
				places[i], places[n-1] = places[n-1], places[i]
			}
		}
	}
}
