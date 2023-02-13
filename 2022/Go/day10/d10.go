package main

import (
	"aoc/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const INPUT_FILENAME string = "input_day10.txt"
const TEST_ANS1 int = 13140
const TEST_ANS2 string = "##..##..##..##..##..##..##..##..##..##..###...###...###...###...###...###...###.####....####....####....####....####....#####.....#####.....#####.....#####.....######......######......######......###########.......#######.......#######....."
const TEST_INPUT string = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

func d1001(input string) int {
	var dx int = 0
	var strength int = 0
	var cycle int = 0
	var regX int = 1
	cycles := [6]int{20, 60, 100, 140, 180, 220}
	for _, instruction := range strings.Split(input, "\n") {
		if len(instruction) < 2 {
			continue
		}
		if instruction == "noop" {
			dx = 0
		} else {
			fields := strings.Fields(instruction)
			dx, _ = strconv.Atoi(fields[1])
			cycle += 1
			if isIn(cycle, cycles) {
				strength += regX * cycle
			}
		}
		cycle += 1
		if isIn(cycle, cycles) {
			strength += regX * cycle
		}
		regX += dx
	}
	return strength
}

func isIn(item int, items [6]int) bool {
	for _, test := range items {
		if item == test {
			return true
		}
	}
	return false
}

func d1002(input string) string {
	var dx int = 0
	var dc int = 1
	var cycle int = 0
	var regX int = 1
	display := buildDisplay(".", 240)
	tolerance := withinBounds(regX)
	for _, instruction := range strings.Split(input, "\n") {
		if len(instruction) < 2 {
			continue
		}
		if instruction == "noop" {
			dx = 0
			dc = 1
		} else {
			fields := strings.Fields(instruction)
			v, _ := strconv.Atoi(fields[1])
			dx = v
			dc = 2
		}
		for offset := 0; offset < dc; offset++ {
			position := cycle + offset
			if inTolerance(position%40, tolerance) {
				display[position] = "#"
			}
		}
		cycle += dc
		regX += dx
		tolerance = withinBounds(regX)
	}
	return strings.Join(display, "")
}

func buildDisplay(fillChar string, size int) (d []string) {
	for i := 0; i < size; i++ {
		d = append(d, fillChar)
	}
	return
}

func withinBounds(regX int) (t []int) {
	for p := regX - 1; p <= regX+1; p++ {
		if p >= 0 {
			t = append(t, p)
		}
	}
	return
}

func inTolerance(tolerance int, tolerances []int) bool {
	for _, element := range tolerances {
		if element == tolerance {
			return true
		}
	}
	return false
}

func showDisplay(display string) {
	for i := 0; i < len(display); i++ {
		if i%40 == 0 {
			fmt.Println()
		}
		fmt.Print(string(display[i]))
	}
	fmt.Println()
}

func main() {
	input := common.ReadFile(INPUT_FILENAME)
	result := d1001(TEST_INPUT)
	if result != TEST_ANS1 {
		errMsg := fmt.Sprintf("Failed tests: day 10 01, received %d instead of %d", result, TEST_ANS1)
		log.Fatal(errMsg)
	}
	result2 := d1002(TEST_INPUT)
	if result2 != TEST_ANS2 {
		errMsg := fmt.Sprintf("Failed tests: day 10 02, received \n%s instead of \n%s", result2, TEST_ANS2)
		log.Fatal(errMsg)
	}
	log.Println("Day 10 01:", d1001(input))
	log.Println("Day 10 02:", d1002(input))
	showDisplay(d1002(input))
}
