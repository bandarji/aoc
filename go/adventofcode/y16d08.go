package adventofcode

import (
	"fmt"
	"strings"
)

const (
	y16d08Tall = 6
	y16d08Wide = 50
)

type y16d08Screen struct {
	tall, wide int
	grid       [][]bool
}

func y16d08(input string, part, tall, wide int) string {
	answer := ""
	s := &y16d08Screen{}
	s.new(tall, wide)
	for _, cmd := range strings.Split(input, "\n") {
		if strings.HasPrefix(cmd, "rect ") {
			s.light(cmd)
		} else if strings.HasPrefix(cmd, "rotate column ") {
			s.rotateColumn(cmd)
		} else if strings.HasPrefix(cmd, "rotate row ") {
			s.rotateRow(cmd)
		}
	}
	if part == 1 {
		answer = fmt.Sprintf("%d", s.countLit())
	} else {
		answer = s.print()
	}
	return answer
}

func (s *y16d08Screen) new(tall, wide int) {
	s.tall = tall
	s.wide = wide
	for r := 0; r < tall; r++ {
		s.grid = append(s.grid, make([]bool, wide))
	}
}

func (s *y16d08Screen) light(cmd string) {
	var wide, tall int
	fmt.Sscanf(cmd, "rect %dx%d", &wide, &tall)
	for row := 0; row < tall; row++ {
		for col := 0; col < wide; col++ {
			s.grid[row][col] = true
		}
	}
}

func (s *y16d08Screen) rotateRow(cmd string) {
	var row, by int
	fmt.Sscanf(cmd, "rotate row y=%d by %d", &row, &by)
	for count := 0; count < by; count++ {
		tmp := s.grid[row][s.wide-1]
		for i := s.wide - 1; i > 0; i-- {
			s.grid[row][i] = s.grid[row][i-1]
		}
		s.grid[row][0] = tmp
	}
}

func (s *y16d08Screen) rotateColumn(cmd string) {
	var col, by int
	fmt.Sscanf(cmd, "rotate column x=%d by %d", &col, &by)
	for count := 0; count < by; count++ {
		tmp := s.grid[s.tall-1][col]
		for i := s.tall - 1; i > 0; i-- {
			s.grid[i][col] = s.grid[i-1][col]
		}
		s.grid[0][col] = tmp
	}
}

func (s y16d08Screen) print() string {
	sb := strings.Builder{}
	sb.WriteString("\n")
	for _, row := range s.grid {
		for _, pixel := range row {
			if pixel {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString("\n")
	return sb.String()
}

func (s y16d08Screen) countLit() (count int) {
	for _, row := range s.grid {
		for _, pixel := range row {
			if pixel {
				count++
			}
		}
	}
	return
}
