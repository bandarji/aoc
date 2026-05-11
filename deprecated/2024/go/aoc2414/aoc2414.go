package aoc2414

import (
	"fmt"
	"strings"
	"time"
)

type Robot struct {
	x, y, vx, vy int
}
type Robots []Robot

const AOC2414_TEST string = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func ReadRobots(input string) (robots Robots) {
	for _, line := range strings.Split(input, "\n") {
		var x, y, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
		robots = append(robots, Robot{x, y, vx, vy})
	}
	return
}

func MoveRobots(robots Robots, width, height, seconds int) {
	for s := 0; s < seconds; s++ {
		for i, r := range robots {
			robots[i].x = (r.x + r.vx + width) % width
			robots[i].y = (r.y + r.vy + height) % height
		}
	}
}

func CountRobots(robots Robots, sx, ex, sy, ey int) (count int) {
	for x := sx; x < ex; x++ {
		for y := sy; y < ey; y++ {
			for _, r := range robots {
				if r.x == x && r.y == y {
					count++
				}
			}
		}
	}
	return
}

func EnoughRobotsInColumn(robots Robots, height, x, threshold int) bool {
	count := 0
	slots := make([]bool, height)
	for _, r := range robots {
		if r.x == x {
			slots[r.y] = true
		}
	}
	for _, s := range slots {
		if s {
			count++
			if count >= threshold {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

func DisplayRobots(robots Robots, width, height int) {
	g := make([][]string, height)
	for i := range g {
		g[i] = make([]string, width)
	}
	for w := range width {
		for h := range height {
			g[h][w] = " "
		}
	}
	for _, r := range robots {
		g[r.y][r.x] = "#"
	}
	for _, row := range g {
		fmt.Println(strings.Join(row, ""))
	}
}

func CalcSafetyFactor(robots Robots, width, height int) (safetyFactor int) {
	quads := [4]int{}
	hw, hh := width/2, height/2
	quads[0] = CountRobots(robots, 0, hw, 0, hh)
	quads[1] = CountRobots(robots, hw+1, width, 0, hh)
	quads[2] = CountRobots(robots, 0, hw, hh+1, height)
	quads[3] = CountRobots(robots, hw+1, width, hh+1, height)
	safetyFactor = quads[0] * quads[1] * quads[2] * quads[3]
	return
}

func Aoc241401(input string, width, height, seconds int) (safetyFactor int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	robots := ReadRobots(input)
	MoveRobots(robots, width, height, seconds)
	safetyFactor = CalcSafetyFactor(robots, width, height)
	return
}

func Aoc241402(input string, width, height int) int {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	robots := ReadRobots(input)
	s := 0
treefound:
	for {
		s++
		MoveRobots(robots, width, height, 1)
		for col := range width {
			if EnoughRobotsInColumn(robots, height, col, 8) {
				break treefound
			}
		}
	}
	return s
}
