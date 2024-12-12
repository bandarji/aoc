package aoc2412

import (
	"fmt"
	"image"
	"strings"
	"time"
)

const AOC2412_TEST_1 string = `AAAA
BBCD
BBCC
EEEC`

const AOC2412_TEST_2 string = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

const AOC2412_TEST_3 string = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func CalcFencingCost(gardens map[image.Point]string, part int) (cost int) {
	visited := map[image.Point]bool{}
	for plot := range gardens {
		if visited[plot] {
			continue
		}
		visited[plot] = true
		area := 1
		perimeter, sides := 0, 0
		q := []image.Point{plot}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				if np := p.Add(d); gardens[np] != gardens[p] {
					perimeter++
					c := p.Add(image.Point{-d.Y, -d.X})
					if gardens[c] != gardens[p] || gardens[c.Add(d)] == gardens[p] {
						sides++
					}
				} else if !visited[np] {
					visited[np] = true
					q = append(q, np)
					area++
				}
			}
		}
		if part == 1 {
			cost += area * perimeter
		} else {
			cost += area * sides
		}
	}
	return
}

func ReadGardens(input string) (gardens map[image.Point]string) {
	gardens = map[image.Point]string{}
	for y, row := range strings.Split(input, "\n") {
		for x, garden := range strings.Split(row, "") {
			gardens[image.Point{X: x, Y: y}] = garden
		}
	}
	return
}

func Aoc241201(input string) (cost int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	gardens := ReadGardens(input)
	cost = CalcFencingCost(gardens, 1)
	return
}

func Aoc241202(input string) (cost int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	gardens := ReadGardens(input)
	cost = CalcFencingCost(gardens, 2)
	return
}
