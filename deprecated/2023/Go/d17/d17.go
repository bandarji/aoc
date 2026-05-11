package d17

import (
	"strings"
)

const TEST1 string = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

const TEST2 string = `111111111111
999999999991
999999999991
999999999991
999999999991`

type Point struct {
	x, y int
}

type State struct {
	point, dir Point
	steps      int
}

func InGrid(grid []string, pos Point) bool {
	return pos.x >= 0 && pos.x < len(grid[0]) && pos.y >= 0 && pos.y < len(grid)
}

func TurnB(p Point) Point {
	return Point{p.y, -p.x}
}

func TurnC(p Point) Point {
	return Point{-p.y, p.x}
}

func HeatLoss(grid []string, start, end Point, loSteps, hiSteps int) int {
	points := []State{{start, Point{1, 0}, 0}, {start, Point{0, 1}, 0}}
	visited := map[State]int{{start, Point{0, 0}, 0}: 0}
	mini := 1<<63 - 1

	for len(points) > 0 {
		current := points[0]
		points = points[1:]

		if current.point == end && current.steps >= loSteps {
			mini = min(mini, visited[current])
		}

		aState := State{Point{current.point.x + current.dir.x, current.point.y + current.dir.y}, current.dir, current.steps + 1}
		if InGrid(grid, aState.point) && current.steps < hiSteps {
			loss := visited[current] + int(grid[aState.point.y][aState.point.x]-'0')
			if val, found := visited[aState]; !found || val > loss {
				visited[aState] = loss
				points = append(points, aState)
			}
		}

		bDir := TurnB(current.dir)
		bState := State{Point{current.point.x + bDir.x, current.point.y + bDir.y}, bDir, 1}
		if InGrid(grid, bState.point) && current.steps >= loSteps {
			loss := visited[current] + int(grid[bState.point.y][bState.point.x]-'0')
			if val, found := visited[bState]; !found || val > loss {
				visited[bState] = loss
				points = append(points, bState)
			}
		}

		cDir := TurnC(current.dir)
		cState := State{Point{current.point.x + cDir.x, current.point.y + cDir.y}, cDir, 1}
		if InGrid(grid, cState.point) && current.steps >= loSteps {
			loss := visited[current] + int(grid[cState.point.y][cState.point.x]-'0')
			if val, found := visited[cState]; !found || val > loss {
				visited[cState] = loss
				points = append(points, cState)
			}
		}
	}
	return mini
}

func Solve(input string, part int) (answer int) {
	grid := strings.Split(input, "\n")
	start, end := Point{0, 0}, Point{len(grid[0]) - 1, len(grid) - 1}
	if part == 1 {
		answer = HeatLoss(grid, start, end, 1, 3)
	} else {
		answer = HeatLoss(grid, start, end, 4, 10)
	}
	return
}
