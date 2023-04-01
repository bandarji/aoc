package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coords struct {
	sx, sy, bx, by int
}

const (
	INPUT      string = "input_day15.txt"
	TEST_ANS1  int    = 26
	TEST_ANS2  int    = 56000011
	TEST_INPUT string = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	Sensor at x=9, y=16: closest beacon is at x=10, y=16
	Sensor at x=13, y=2: closest beacon is at x=15, y=3
	Sensor at x=12, y=14: closest beacon is at x=10, y=16
	Sensor at x=10, y=20: closest beacon is at x=10, y=16
	Sensor at x=14, y=17: closest beacon is at x=10, y=16
	Sensor at x=8, y=7: closest beacon is at x=2, y=10
	Sensor at x=2, y=0: closest beacon is at x=2, y=10
	Sensor at x=0, y=11: closest beacon is at x=2, y=10
	Sensor at x=20, y=14: closest beacon is at x=25, y=17
	Sensor at x=17, y=20: closest beacon is at x=21, y=22
	Sensor at x=16, y=7: closest beacon is at x=15, y=3
	Sensor at x=14, y=3: closest beacon is at x=15, y=3
	Sensor at x=20, y=1: closest beacon is at x=15, y=3`
)

func ParseInput(input string) (coords []Coords) {
	input = strings.TrimRight(input, "\n")
	for _, entry := range strings.Split(input, "\n") {
		words := strings.Split(entry, " ")
		coords = append(coords, Coords{sx: atoi(words[2]), sy: atoi(words[3]), bx: atoi(words[8]), by: atoi(words[9])})
	}
	return
}

func atoi(s string) int {
	ns := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 45 || (s[i] <= 57 && s[i] >= 48) {
			ns += string(s[i])
		}
	}
	num, _ := strconv.Atoi(ns)
	return num
}

func Manhattan(x1, y1, x2, y2 int) (dist int) {
	dist += abs(x1 - x2)
	dist += abs(y1 - y2)
	return
}

func abs(value int) int {
	if value < 0 {
		value *= -1
	}
	return value
}

func Day1501(coords []Coords, row int) (answer int) {
	cant := map[int]int{}
	for _, pairs := range coords {
		dist := Manhattan(pairs.sx, pairs.sy, pairs.bx, pairs.by)
		offset := dist - abs(pairs.sy-row)
		if offset < 0 {
			continue
		}
		lox := pairs.sx - offset
		hix := pairs.sx + offset
		for x := lox; x <= hix; x++ {
			cant[x] = 1
		}
		if pairs.by == row {
			delete(cant, pairs.bx)
		}
	}
	answer = len(cant)
	return
}

func Day1502(coords []Coords, row int) int {
	// cant := map[int]int{}
	// known := map[int]int{}
	// intervals := [][2]int{}
	for y := 0; y <= 4000000; y++ {
		intervals := [][2]int{}
		for _, pairs := range coords {
			dist := Manhattan(pairs.sx, pairs.sy, pairs.bx, pairs.by)
			offset := dist - abs(pairs.sy-y)
			if offset < 0 {
				continue
			}
			lox := pairs.sx - offset
			hix := pairs.sx + offset
			intervals = append(intervals, [2]int{lox, hix})
		}
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i][1] < intervals[j][1]
		})
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})
		overlaps := [][2]int{}
		for i, interval := range intervals {
			if i == 0 {
				overlaps = append(overlaps, interval)
				continue
			}
			lastinterval := overlaps[len(overlaps)-1]
			if interval[0] > lastinterval[1] {
				overlaps = append(overlaps, interval)
				continue
			}
			overlaps[len(overlaps)-1][1] = max(lastinterval[1], interval[1])
		}
		x := 0
		for _, overlap := range overlaps {
			loo, hio := overlap[0], overlap[1]
			if x < loo {
				return y + x*4000000
			}
			x = max(x, hio+1)
			if x > 4000000 {
				break
			}
		}
	}
	return -1
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func RunTest(input string, part, row int) (result int) {
	coords := ParseInput(input)
	if part == 1 {
		result = Day1501(coords, row)
	} else {
		result = Day1502(coords, row)
	}
	return
}

func main() {
	day, part, row := 15, 1, 10
	result := RunTest(TEST_INPUT, part, row)
	if result != TEST_ANS1 {
		log.Fatalf("Day %02d / %02d test failed: %d vs %d\n", day, part, result, TEST_ANS1)
	} else {
		log.Printf("Day %02d / %02d test passed: %d vs %d\n", day, part, result, TEST_ANS1)
	}
	part = 2
	result = RunTest(TEST_INPUT, part, row)
	if result != TEST_ANS2 {
		log.Fatalf("Day %02d / %02d test failed: %d vs %d\n", day, part, result, TEST_ANS2)
	} else {
		log.Printf("Day %02d / %02d test passed: %d vs %d\n", day, part, result, TEST_ANS2)
	}
	content, _ := os.ReadFile(INPUT)
	coords := ParseInput(string(content))
	log.Printf("Day %02d / %02d answer: %d\n", day, part, Day1501(coords, 2000000))
	log.Printf("Day %02d / %02d answer: %d\n", day, part, Day1502(coords, 2000000))
}
