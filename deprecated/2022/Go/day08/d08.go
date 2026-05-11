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
	grid := buildGrid(input)
	visible := countVisibleTrees(grid)
	return visible
}

func d0802(input string) int {
	topScore := 0
	grid := buildGrid(input)
	dimY := len(grid)
	dimX := len(grid[0])
	for y := 0; y < dimY; y++ {
		for x := 0; x < dimX; x++ {
			if x == 0 || x == dimX-1 || y == 0 || y == dimY-1 {
				continue
			}
			score := scenicScore(grid, y, x)
			if score > topScore {
				topScore = score
			}
		}
	}
	return topScore
}

func scenicScore(grid Grid, y int, x int) int {
	score := countDirection(grid, y, x, 0, -1)
	score *= countDirection(grid, y, x, 0, 1)
	score *= countDirection(grid, y, x, -1, 0)
	score *= countDirection(grid, y, x, 1, 0)
	return score
}

func countVisibleTrees(grid Grid) int {
	visible := 0
	for y, row := range grid {
		for x, _ := range row {
			if isVisible(grid, y, x, 0, -1) {
				visible += 1
			} else if isVisible(grid, y, x, 0, 1) {
				visible += 1
			} else if isVisible(grid, y, x, -1, 0) {
				visible += 1
			} else if isVisible(grid, y, x, 1, 0) {
				visible += 1
			}
		}
	}
	return visible
}

func countDirection(grid Grid, y int, x int, dy int, dx int) int {
	count := 0
	height := grid[y][x]
	r := y + dy
	c := x + dx
	for r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) {
		if grid[r][c] < height {
			count++
		} else {
			count++
			break
		}
		r += dy
		c += dx
	}
	return count
}

func isVisible(grid Grid, y int, x int, dy int, dx int) bool {
	if y == 0 || x == 0 || y == len(grid)-1 || x == len(grid[0])-1 {
		return true
	}
	height := grid[y][x]
	r := y + dy
	c := x + dx
	for r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) {
		if height <= grid[r][c] {
			return false
		}
		r += dy
		c += dx
	}
	return true
}

func buildGrid(input string) Grid {
	var grid Grid
	for _, line := range strings.Split(input, "\n") {
		if len(line) < 2 {
			continue
		}
		var row []int
		for _, height := range line {
			row = append(row, castStringToInt(string(height)))
		}
		grid = append(grid, row)
	}
	return grid
}

func castStringToInt(input string) int {
	i, _ := strconv.Atoi(input)
	return i
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
