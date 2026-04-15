package adventofcode

import (
	"log"
	"strings"
)

func y16d24(input string, part int) int {
	grid := y16d24ParseInput(input)
	graph := y16d24DetermineDistances(grid)
	_ = graph
	// y16d24DFS(graph, map[int]bool{0: true}, 0, part)
	return 0
}

func y16d24DFS(graph [][]int, visited map[int]bool, current, part int) int {
	if len(visited) == len(graph) {
		if part == 2 {
			return graph[current][0]
		}
		return 0
	}
	minDist := 1 << 31
	for i, value := range graph[current] {
		if !visited[i] {
			visited[i] = true
			dist := value + y16d24DFS(graph, visited, i, part)
			minDist = min(minDist, dist)
			delete(visited, i)
		}
	}
	return minDist
}

func y16d24ParseInput(input string) [][]string {
	grid := [][]string{}
	for line := range strings.SplitSeq(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func y16d24DetermineDistances(grid [][]string) (graph [][]int) {
	graph = [][]int{}
	for r, row := range grid {
		for c, cell := range row {
			if strings.ContainsAny(cell, "0123456789") {
				robotPoint := cell
				distancesFromRobot := y16d24BFSRobots(grid, [2]int{r, c})
				if graph == nil {
					for i := range len(distancesFromRobot) {
						_ = i
						graph = append(graph, make([]int, len(distancesFromRobot)))
					}
				}
				graph[strToInt(robotPoint)] = distancesFromRobot
			}
		}
	}
	return
}

func y16d24BFSRobots(grid [][]string, start [2]int) map[string]int {
	toDistance := map[string]int{
		grid[start[0]][start[1]]: 0,
	}
	visited := map[[2]int]bool{}
	q := [][3]int{{start[0], start[1], 0}}
	for len(q) > 0 {
		log.Printf("map: %v", toDistance)
		curr := q[0]
		q = q[1:]
		r, c, d := curr[0], curr[1], curr[2]
		if visited[[2]int{r, c}] {
			continue
		}
		visited[[2]int{r, c}] = true
		if strings.ContainsAny(grid[r][c], "0123456789") {
			toDistance[grid[r][c]] = d
		}
		for _, dx := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nr, nc := r+dx[0], c+dx[1]
			if grid[nr][nc] != "#" {
				q = append(q, [3]int{nr, nc, d + 1})
			}
		}
	}
	distances := map[string]int{}
	for robot, distance := range toDistance {
		distances[robot] = distance
	}
	return distances
}
