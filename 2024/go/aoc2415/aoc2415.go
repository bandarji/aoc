package aoc2415

import (
	"fmt"
	"image"
	"strings"
	"time"
)

const AOC2415_TEST_1 string = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

const AOC2415_TEST_2 string = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

const AOC2415_TEST_3 string = `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

func ReadData(input string) (grid [][]string, movements string, robot image.Point) {
	grid = [][]string{}
	sb := strings.Builder{}
	sections := strings.Split(input, "\n\n")
	for _, line := range strings.Split(sections[0], "\n") {
		grid = append(grid, strings.Split(line, ""))
	}
	for _, line := range strings.Split(sections[1], "\n") {
		sb.WriteString(line)
	}
	movements = sb.String()
	for y, row := range grid {
		for x, cell := range row {
			if cell == "@" {
				robot = image.Point{x, y}
			}
		}
	}
	return
}

func PushUp(grid [][]string, robot image.Point) int {
	Os := 0
	for y := robot.Y - 1; y >= 0; y-- {
		if grid[y][robot.X] == "O" {
			Os++
		} else if grid[y][robot.X] == "." {
			return Os
		} else {
			break
		}
	}
	return 0
}

func PushDown(grid [][]string, robot image.Point) int {
	Os := 0
	for y := robot.Y + 1; y < len(grid); y++ {
		if grid[y][robot.X] == "O" {
			Os++
		} else if grid[y][robot.X] == "." {
			return Os
		} else {
			break
		}
	}
	return 0
}

func PushLeft(grid [][]string, robot image.Point) int {
	Os := 0
	for x := robot.X - 1; x >= 0; x-- {
		if grid[robot.Y][x] == "O" {
			Os++
		} else if grid[robot.Y][x] == "." {
			return Os
		} else {
			break
		}
	}
	return 0
}

func PushRight(grid [][]string, robot image.Point) int {
	Os := 0
	for x := robot.X + 1; x < len(grid[0]); x++ {
		if grid[robot.Y][x] == "O" {
			Os++
		} else if grid[robot.Y][x] == "." {
			return Os
		} else {
			break
		}
	}
	return 0
}

func MoveRobot(grid [][]string, movements string, robot image.Point) {
	positions := 0
	_ = positions
	for _, move := range movements {
		// fmt.Printf("Move %c:\n", move)
		switch move {
		case '^':
			if grid[robot.Y-1][robot.X] == "." {
				grid[robot.Y][robot.X] = "."
				grid[robot.Y-1][robot.X] = "@"
				robot = image.Point{robot.X, robot.Y - 1}
			} else {
				if grid[robot.Y-1][robot.X] == "O" {
					positions = PushUp(grid, robot)
					if positions > 0 {
						grid[robot.Y-positions-1][robot.X] = "O"
						grid[robot.Y][robot.X] = "."
						grid[robot.Y-1][robot.X] = "@"
						robot = image.Point{robot.X, robot.Y - 1}
					}
				}
			}
		case 'v':
			if grid[robot.Y+1][robot.X] == "." {
				grid[robot.Y][robot.X] = "."
				grid[robot.Y+1][robot.X] = "@"
				robot = image.Point{robot.X, robot.Y + 1}
			} else {
				if grid[robot.Y+1][robot.X] == "O" {
					positions = PushDown(grid, robot)
					if positions > 0 {
						grid[robot.Y+positions+1][robot.X] = "O"
						grid[robot.Y][robot.X] = "."
						grid[robot.Y+1][robot.X] = "@"
						robot = image.Point{robot.X, robot.Y + 1}
					}
				}
			}
		case '<':
			if grid[robot.Y][robot.X-1] == "." {
				grid[robot.Y][robot.X] = "."
				grid[robot.Y][robot.X-1] = "@"
				robot = image.Point{robot.X - 1, robot.Y}
			} else {
				if grid[robot.Y][robot.X-1] == "O" {
					positions = PushLeft(grid, robot)
					if positions > 0 {
						grid[robot.Y][robot.X-positions-1] = "O"
						grid[robot.Y][robot.X] = "."
						grid[robot.Y][robot.X-1] = "@"
						robot = image.Point{robot.X - 1, robot.Y}
					}
				}
			}
		case '>':
			if grid[robot.Y][robot.X+1] == "." {
				grid[robot.Y][robot.X] = "."
				grid[robot.Y][robot.X+1] = "@"
				robot = image.Point{robot.X + 1, robot.Y}
			} else {
				if grid[robot.Y][robot.X+1] == "O" {
					positions = PushRight(grid, robot)
					if positions > 0 {
						grid[robot.Y][robot.X+positions+1] = "O"
						grid[robot.Y][robot.X] = "."
						grid[robot.Y][robot.X+1] = "@"
						robot = image.Point{robot.X + 1, robot.Y}
					}
				}
			}
		}
		// DisplayGrid(grid)
		// fmt.Println()
	}
}

func AddGPS(grid [][]string) (total int) {
	for y, row := range grid {
		for x, cell := range row {
			if cell == "O" {
				total += 100*y + x
			}
		}
	}
	return
}

func DisplayGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(strings.Join(row, ""))
	}
}

func Aoc241501(input string) (ans int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	grid, instructions, robot := ReadData(input)
	MoveRobot(grid, instructions, robot)
	// DisplayGrid(grid)
	ans = AddGPS(grid)
	return
}

func Aoc241502(input string) (ans int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	return
}
