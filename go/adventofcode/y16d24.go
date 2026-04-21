package adventofcode

import (
	"fmt"
	"image"
	"log"
	"strings"
)

func y16d24(input string, part int) (answer int) {
	start, stops, path := y16d24ParseInput(input)
	// log.Println("start", start, "stops", stops)
	answer = y16d24FindShortestPath(start, stops, path, part)
	return
}

type y16d24State struct {
	pos          image.Point
	steps        int
	destinations string
}

func y16d24FindShortestPath(start image.Point, stops, path map[image.Point]bool, part int) (steps int) {
	visited := map[image.Point]bool{}
	destinations := ""
	q := []y16d24State{{pos: start, steps: 0, destinations: destinations}}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		// log.Println("current", current)
		if visited[current.pos] {
			continue
		}
		visited[current.pos] = true
		if len(current.destinations) == len(stops) {
			log.Println("current", current.destinations, "stops", stops, "steps", current.steps)
			steps = current.steps
			break
		}
		for _, dir := range [4]image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			next := current.pos.Add(dir)
			if path[next] {
				destinations := current.destinations
				if stops[next] {
					destinations = fmt.Sprintf("%s%d", destinations, next)
				}
				q = append(q, y16d24State{pos: next, steps: current.steps + 1, destinations: destinations})
			}
		}
	}
	return
}

func y16d24ParseInput(input string) (image.Point, map[image.Point]bool, map[image.Point]bool) {
	stops := map[image.Point]bool{}
	start := image.Point{}
	path := map[image.Point]bool{}
	for y, line := range strings.Split(input, "\n") {
		for x, cell := range line {
			switch cell {
			case '0':
				start = image.Point{x, y}
				stops[image.Point{x, y}] = true
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				stops[image.Point{x, y}] = true
			case '.':
				path[image.Point{x, y}] = true
			}
		}
	}
	return start, stops, path
}
