package adventofcode

import (
	"fmt"
	"image"
	"strings"
)

type y16d24AirDuct struct {
	start image.Point
	stops map[image.Point]bool
	path  map[image.Point]bool
}

func y16d24(input string, part int) (answer int) {
	duct := y16d24ParseInput(input)
	answer = y16d24Spelunk(duct, part)
	return
}

type y16d24State struct {
	pos         image.Point
	steps       int
	checkpoints map[image.Point]bool
}

func (s y16d24State) encode() string {
	checkpoints := []string{}
	for k := range s.checkpoints {
		checkpoints = append(checkpoints, fmt.Sprintf("[%d,%d]", k.X, k.Y))
	}
	return fmt.Sprintf("%d,%d,%s", s.pos.X, s.pos.Y, strings.Join(checkpoints, ","))
}

func y16d24Spelunk(duct y16d24AirDuct, part int) int {
	cache := map[string]bool{}
	q := []y16d24State{{pos: duct.start, steps: 0, checkpoints: map[image.Point]bool{duct.start: true}}}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		encoded := current.encode()
		if cache[encoded] {
			continue
		}
		cache[encoded] = true
		if len(current.checkpoints) == len(duct.stops) {
			if part == 1 {
				return current.steps
			}
			if current.pos == duct.start {
				return current.steps
			}
		}
		for _, dir := range [4]image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			next := current.pos.Add(dir)
			if duct.path[next] {
				np := map[image.Point]bool{}
				for k := range current.checkpoints {
					np[k] = true
				}
				if duct.stops[next] {
					np[next] = true
				}
				q = append(q, y16d24State{pos: next, steps: current.steps + 1, checkpoints: np})
			}
		}
	}
	return -1
}

func y16d24ParseInput(input string) y16d24AirDuct {
	duct := y16d24AirDuct{stops: map[image.Point]bool{}, path: map[image.Point]bool{}}
	for y, line := range strings.Split(input, "\n") {
		for x, cell := range line {
			if cell != '#' {
				duct.path[image.Point{x, y}] = true
				if cell >= '0' && cell <= '9' {
					duct.stops[image.Point{x, y}] = true
				}
				if cell == '0' {
					duct.start = image.Point{x, y}
				}
			}
		}
	}
	return duct
}
