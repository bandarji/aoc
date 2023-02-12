package main

import (
	"aoc/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

const INPUT_FILENAME string = "input_day09.txt"
const TEST_INPUT1 string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

const TEST_INPUT2 string = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

const TEST_ANS1 int = 13
const TEST_ANS2 int = 36

func d0901(input string) int {
	points := make([]Point, 1)
	head := Point{x: 0, y: 0}
	tail := Point{x: 0, y: 0}
	points[0] = Point{x: 0, y: 0}
	for _, line := range strings.Split(input, "\n") {
		if len(line) < 2 {
			continue
		}
		fields := strings.Fields(line)
		dx, dy := setDeltas(fields[0])
		count, _ := strconv.Atoi(fields[1])
		for i := 0; i < count; i++ {
			head.x += dx
			head.y += dy
			tail = moveTail(head, tail)
			uniqueAdd(&points, tail)
		}
	}
	return len(points)
}

func d0902(input string) int {
	points := make([]Point, 1)
	tailSpots := make([]Point, 9)
	head := Point{x: 0, y: 0}
	points[0] = Point{x: 0, y: 0}
	for _, line := range strings.Split(input, "\n") {
		if len(line) < 2 {
			continue
		}
		fields := strings.Fields(line)
		count, _ := strconv.Atoi(fields[1])
		for i := 0; i < count; i++ {
			dx, dy := setDeltas(fields[0])
			head.x += dx
			head.y += dy
			tailSpots[0] = moveTail(head, tailSpots[0])
			for m := 1; m < 9; m++ {
				tailSpots[m] = moveTail(tailSpots[m-1], tailSpots[m])
				uniqueAdd(&points, tailSpots[8])
			}
		}
	}
	return len(points)
}

func moveTail(head, tail Point) Point {
	dx := abs(head.x, tail.x)
	dy := abs(head.y, tail.y)
	if dx > 1 && dy > 1 {
		if tail.x < head.x {
			tail.x += 1
		} else {
			tail.x -= 1
		}
		if tail.y < head.y {
			tail.y += 1
		} else {
			tail.y -= 1
		}
	} else if dx > 1 {
		tail.y = head.y
		if tail.x < head.x {
			tail.x += 1
		} else {
			tail.x -= 1
		}
	} else if dy > 1 {
		tail.x = head.x
		if tail.y < head.y {
			tail.y += 1
		} else {
			tail.y -= 1
		}
	} else {
		return tail
	}
	return tail
}

func uniqueAdd(points *[]Point, point Point) {
	for _, aPoint := range *points {
		if point.x == aPoint.x && point.y == aPoint.y {
			return
		}
	}
	*points = append(*points, point)
}

func abs(v1, v2 int) (diff int) {
	diff = v1 - v2
	if diff < 0 {
		diff *= -1
	}
	return
}

func setDeltas(dir string) (dx, dy int) {
	switch dir {
	case "L":
		dx = -1
	case "R":
		dx = 1
	case "U":
		dy = 1
	case "D":
		dy = -1
	}
	return
}

func main() {
	var result int = 0
	input := common.ReadFile(INPUT_FILENAME)
	result = d0901(TEST_INPUT1)
	if result != TEST_ANS1 {
		errMsg := fmt.Sprintf("Failed tests: day 09 01, received %d instead of %d", result, TEST_ANS1)
		log.Fatal(errMsg)
	}
	result = d0902(TEST_INPUT2)
	if result != TEST_ANS2 {
		errMsg := fmt.Sprintf("Failed tests: day 09 02, received %d instead of %d", result, TEST_ANS2)
		log.Fatal(errMsg)
	}
	log.Println("Day 09 01:", d0901(input))
	log.Println("Day 09 02:", d0902(input))
}
