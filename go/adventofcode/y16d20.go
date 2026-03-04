package adventofcode

import (
	"fmt"
	"sort"
	"strings"
)

type y16d20FirewalledRange struct {
	start, end int
}

func y16d20(input string, ceiling, part int) int {
	ranges := y16d20ReadRanges(input)
	if part == 1 {
		return ranges[0].end + 1
	}

	rangeLength := len(ranges)

	if ranges[rangeLength-1].end != ceiling {
		ranges = append(ranges, y16d20FirewalledRange{start: ranges[rangeLength-2].end + 1, end: ceiling})
	}
	sum := 0
	for i := 1; i < rangeLength; i++ {
		sum += ranges[i].start - ranges[i-1].end - 1
	}

	return sum
}

func y16d20ReadRanges(input string) []y16d20FirewalledRange {
	ranges := []y16d20FirewalledRange{}
	for line := range strings.SplitAfterSeq(input, "\n") {
		var start, end int
		fmt.Sscanf(line, "%d-%d", &start, &end)
		ranges = append(ranges, y16d20FirewalledRange{start: start, end: end})
	}

	// sort the firewalled ranges
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].start != ranges[j].start {
			return ranges[i].start < ranges[j].start
		}
		return ranges[i].end < ranges[j].end
	})

	// merge overlapping address blocks
	m := []y16d20FirewalledRange{{}}
	for _, r := range ranges {
		highest := m[len(m)-1].end
		if highest >= r.start-1 {
			m[len(m)-1].end = max(highest, r.end)
		} else {
			m = append(m, r)
		}
	}

	return m
}
