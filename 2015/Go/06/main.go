package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Day06(input string, part int) (brightness int) {
	grid := ProcessInstructions(input, part)
	brightness = CountLit(&grid)
	return
}

func CountLit(grid *[][]int) (count int) {
	for _, row := range *grid {
		for _, cell := range row {
			count += cell
		}
	}
	return
}

func ProcessInstructions(input string, part int) [][]int {
	var sy, sx, ey, ex int
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}
	input = strings.TrimRight(input, "\n")
	for _, instruction := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(instruction, "toggle"):
			fmt.Sscanf(instruction, "toggle %d,%d through %d,%d", &sy, &sx, &ey, &ex)
			for y := sy; y <= ey; y++ {
				for x := sx; x <= ex; x++ {
					if part == 1 {
						if grid[y][x] == 0 {
							grid[y][x] = 1
						} else {
							grid[y][x] = 0
						}
					} else {
						grid[y][x] += 2
					}
				}
			}
		case strings.HasPrefix(instruction, "turn on"):
			fmt.Sscanf(instruction, "turn on %d,%d through %d,%d", &sy, &sx, &ey, &ex)
			for y := sy; y <= ey; y++ {
				for x := sx; x <= ex; x++ {
					if part == 1 {
						grid[y][x] = 1
					} else {
						grid[y][x]++
					}
				}
			}
		case strings.HasPrefix(instruction, "turn off"):
			fmt.Sscanf(instruction, "turn off %d,%d through %d,%d", &sy, &sx, &ey, &ex)
			for y := sy; y <= ey; y++ {
				for x := sx; x <= ex; x++ {
					if part == 1 {
						grid[y][x] = 0
					} else {
						grid[y][x]--
						if grid[y][x] < 0 {
							grid[y][x] = 0
						}
					}
				}
			}
		}
	}
	return grid
}

func main() {
	day, part := 6, 1
	content, _ := os.ReadFile("input_day06.txt")
	lit := Day06(string(content), part)
	log.Printf("[%02d / %02d]: %d\n", day, part, lit)
	part = 2
	brightness := Day06(string(content), part)
	log.Printf("[%02d / %02d]: %d\n", day, part, brightness)
}
