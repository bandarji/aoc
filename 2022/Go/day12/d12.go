package main

import (
	"os"
	"fmt"
	"log"
	"strings"
)

const INPUT_FILENAME string = "input_day12.txt"
const TEST_ANS1 int = 31
const TEST_ANS2 int = 29
const TEST_INPUT string = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func Day1201(input string) (steps int) {
	var diffs = [4][2]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	visited := map[[2]int]bool{}
	queue := [][3]int{}
	grid := buildGrid(input)
	r, c := findChar(grid, "S")
	queue = append(queue, [3]int{r, c, 0})
	start := queue[0]
	for len(queue) > 0 {
		start = queue[0]
		sRow, sCol, cost := start[0], start[1], start[2]
		queue = queue[1:]
		if visited[[2]int{sRow, sCol}] { continue }
		visited[[2]int{sRow, sCol}] = true
		if grid[sRow][sCol] == "E" { return cost }
		for _, diff := range diffs {
			dRow, dCol := diff[0], diff[1]
			nRow, nCol := sRow + dRow, sCol + dCol
			if nRow >= 0 && nRow < len(grid) && nCol >= 0 && nCol < len(grid[0]) {
				charDelta := calcDistance(grid[sRow][sCol], grid[nRow][nCol])
				if charDelta <= 1 {
					queue = append(queue, [3]int{nRow, nCol, start[2] + 1})
				}
			}
		}
	}
	return
}

func Day1202(input string) (steps int) {
	var diffs = [4][2]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	grid := buildGrid(input)
	r, c := findChar(grid, "E")
	visited := map[[2]int]bool{}
	queue := [][3]int{}
	queue = append(queue, [3]int{r, c, 0})
	for len(queue) > 0 {
		start := queue[0]
		sRow, sCol, cost := start[0], start[1], start[2]
		queue = queue[1:]
		if visited[[2]int{sRow, sCol}] { continue }
		visited[[2]int{sRow, sCol}] = true
		if grid[sRow][sCol] == "a" { return cost }
		for _, diff := range diffs {
			dRow, dCol := diff[0], diff[1]
			nRow, nCol := sRow + dRow, sCol + dCol
			if nRow >= 0 && nRow < len(grid) && nCol >= 0 && nCol < len(grid[0]) {
				charDelta := calcDistance(grid[sRow][sCol], grid[nRow][nCol])
				if charDelta >= -1 {
					queue = append(queue, [3]int{nRow, nCol, start[2] + 1})
				}
			}
		}
	}
	return
}

func findChar(grid [][]string, search string) (r, c int) {
	for r, row := range grid {
		for c, spot := range row {
			if spot == search {
				return r, c
			}
		}
	}
	return
}

func calcDistance(charA, charB string) int {
	if charA == "S" {charA = "a"}
	if charA == "E" {charA = "z"}
	if charB == "S" {charB = "a"}
	if charB == "E" {charB = "z"}
	return charToOrd(charB) - charToOrd(charA)
}

func charToOrd(c string) (ord int) {
	ords := []byte(c)
	if len(ords) > 0 {
		ord = int(ords[0])
	} else {
		ord = -1
	}
	return
}

func RunTests(day int) {
	result := 0
	if day == 1 {
		result = Day1201(TEST_INPUT)
		if result != TEST_ANS1 {
			log.Fatal(fmt.Sprintf("Day 12 01 test failed: %d instead of %d.", result, TEST_ANS1))
		} else {
			fmt.Println("Day 12 01 test passed.")
		}
	}
	if day == 2 {
		result = Day1202(TEST_INPUT)
		if result != TEST_ANS2 {
			log.Fatal(fmt.Sprintf("Day 12 02 test failed: %d instead of %d.", result, TEST_ANS2))
		} else {
			fmt.Println("Day 12 02 test passed.")
		}
	}
}

func buildGrid(input string) (grid [][]string) {
	input = strings.TrimRight(input, "\n")
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}
	return
}

func main() {
	fmt.Println("Day 12")
	RunTests(1)
	RunTests(2)
	input, _ := os.ReadFile(INPUT_FILENAME)
	fmt.Println("Day 12 01:", Day1201(string(input)))
	fmt.Println("Day 12 02:", Day1202(string(input)))
}