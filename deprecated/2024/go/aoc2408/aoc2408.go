package aoc2408

import (
	"image"
	"strings"
)

type PointSet map[image.Point]bool
type AntennaeMap map[string][]image.Point

const AOC2408_TEST string = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func BuildGrid(input string) (grid [][]string) {
	grid = [][]string{}
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}
	return
}

func MapAntennae(grid [][]string) (points PointSet, antennae AntennaeMap) {
	antennae = AntennaeMap{}
	points = PointSet{}
	for y, row := range grid {
		for x, cell := range row {
			points[image.Point{X: x, Y: y}] = true
			if cell != "." {
				antennae[cell] = append(antennae[cell], image.Point{X: x, Y: y})
			}
		}
	}
	return
}

func CountAntinodes(part int, points PointSet, antennae AntennaeMap) (count int) {
	antinodes := PointSet{}
	for _, antenna := range antennae {
		for _, a1 := range antenna {
			for _, a2 := range antenna {
				if a1 == a2 {
					continue
				}
				switch part {
				case 1:
					if a := a2.Add(a2.Sub(a1)); points[a] {
						antinodes[a] = true
					}
				case 2:
					for delta := a2.Sub(a1); points[a2]; a2 = a2.Add(delta) {
						antinodes[a2] = true
					}
				}
			}
		}
	}
	count = len(antinodes)
	return
}

func Aoc240801(input string) (antinodes int) {
	grid := BuildGrid(input)
	points, antennaeMap := MapAntennae(grid)
	// fmt.Printf("Points: %v\nAntennaeMap: %+v\n", points, antennaeMap)
	antinodes = CountAntinodes(1, points, antennaeMap)
	return
}

func Aoc240802(input string) (antinodes int) {
	grid := BuildGrid(input)
	points, antennaeMap := MapAntennae(grid)
	antinodes = CountAntinodes(2, points, antennaeMap)
	return
}
