package main

import (
	"aoc/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Grid [][]int
type Dimensions struct {
	rows int
	cols int
}

const INPUT_FILENAME = "input_day08.txt"
const TEST_INPUT string = `30373
25512
65332
33549
35390`
const TEST_ANS1 int = 21
const TEST_ANS2 int = 8

func d0801(input string) int {
	visible := 0
	dims := getDimensions(input)
	grid := buildGrid(input, dims)
	for row, trees := range grid {
		if row == 0 || row == len(grid)-1 {
			continue
		}
		visible += countVisible(grid, row, trees)
	}
	return 0
}

func d0802(input string) int {
	return 0
}

func countVisible(grid Grid, row int, trees []int) int {
	visible := 0
	for column, tree := range trees {
		if column == 0 || column == len(trees)-1 {
			continue
		}
		if treeVisible(grid, row, column, tree) {
			visible++
		}
	}
}

func getDimensions(input string) Dimensions {
	dims := Dimensions{rows: 0, cols: 0}
	for _, line := range strings.Split(input, "\n") {
		if len(line) > 2 {
			dims.rows++
			dims.cols = len(line)
		}
	}
	return dims
}

func buildGrid(input string, dims Dimensions) Grid {
	grid := make(Grid, dims.rows)
	for i, row := range strings.Split(input, "\n") {
		if len(row) < 2 {
			continue
		}
		grid[i] = make([]int, dims.cols)
		for j, height := range row {
			grid[i][j], _ = strconv.Atoi(string(height))
		}
	}
	return grid
}

func main() {
	var result int = 0
	input := common.ReadFile(INPUT_FILENAME)
	result = d0801(TEST_INPUT)
	if result != TEST_ANS1 {
		errMsg := fmt.Sprintf("Failed tests: day 08 01, received %d instead of %d", result, TEST_ANS1)
		log.Fatal(errMsg)
	}
	result = d0802(TEST_INPUT)
	if result != TEST_ANS2 {
		errMsg := fmt.Sprintf("Failed tests: day 08 01, received %d instead of %d", result, TEST_ANS2)
		log.Fatal(errMsg)
	}
	log.Println("Day 08 01:", d0801(input))
	log.Println("Day 08 02:", d0802(input))
}
