package aoc2416

import (
	"fmt"
	"image"
	"maps"
	"slices"
	"strings"
	"time"
)

type Deer struct {
	Location, Facing image.Point
}

type Item struct {
	Deer Deer
	Cost int
	Path map[image.Point]bool
}

const AOC2416_TEST string = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

func ReadMaze(input string) (maze map[image.Point]string, S, E image.Point) {
	maze = map[image.Point]string{}
	for y, line := range strings.Split(input, "\n") {
		for x, s := range strings.Split(line, "") {
			maze[image.Pt(x, y)] = s
			if s == "S" {
				S = image.Pt(x, y)
			}
			if s == "E" {
				E = image.Pt(x, y)
			}
		}
	}
	return
}

func ProcessMaze(input string) (ans1, ans2 int) {
	ans1 = 1 << 20
	p2 := map[image.Point]bool{}
	maze, S, E := ReadMaze(input)
	distances := map[Deer]int{}
	queue := []Item{
		{
			Deer: Deer{S, image.Pt(1, 0)},
			Cost: 0,
			Path: map[image.Point]bool{S: true},
		},
	}
	for len(queue) > 0 {
		// fmt.Printf("Queue: %+v\n", queue)
		slices.SortFunc(queue, func(i, j Item) int {
			if i.Cost == j.Cost {
				return 0
			}
			if i.Cost > j.Cost {
				return 1
			}
			return -1
		})
		i := queue[0]
		queue = queue[1:]
		if cost, ok := distances[i.Deer]; ok && cost < i.Cost {
			continue
		}
		distances[i.Deer] = i.Cost
		if i.Deer.Location.X == E.X && i.Deer.Location.Y == E.Y && i.Cost <= ans1 {
			ans1 = i.Cost
			maps.Copy(p2, i.Path)

		}
		for dir, cost := range map[image.Point]int{
			i.Deer.Facing:                       1,
			{-i.Deer.Facing.Y, i.Deer.Facing.X}: 1001,
			{i.Deer.Facing.Y, -i.Deer.Facing.X}: 1001,
		} {
			nd := Deer{i.Deer.Location.Add(dir), dir}
			if maze[nd.Location] == "#" {
				continue
			}
			path := maps.Clone(i.Path)
			path[nd.Location] = true
			queue = append(queue, Item{nd, i.Cost + cost, path})
		}
	}
	ans2 = len(p2)
	return
}

func Aoc241601(input string) (score int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	score, _ = ProcessMaze(input)
	return
}

func Aoc241602(input string) (tiles int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	_, tiles = ProcessMaze(input)
	return
}
