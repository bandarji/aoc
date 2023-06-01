package d08

import (
	"fmt"
	"strings"
)

type Screen struct {
	Grid          [][]bool
	Width, Height int
	Lit           int
}

func (s *Screen) CountLit() {
	s.Lit = 0
	for _, row := range s.Grid {
		for _, pixel := range row {
			if pixel {
				s.Lit++
			}
		}
	}
}

func (s *Screen) Make(rows, cols int) {
	s.Width = cols
	s.Height = rows
	for r := 0; r < rows; r++ {
		s.Grid = append(s.Grid, make([]bool, cols))
	}
}

func (s *Screen) LightBox(cmd string) {
	var width, height int
	fmt.Sscanf(cmd, "rect %dx%d", &width, &height)
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			s.Grid[row][col] = true
		}
	}
}

func (s *Screen) RotateRow(cmd string) {
	var row, by int
	fmt.Sscanf(cmd, "rotate row y=%d by %d", &row, &by)
	for count := 0; count < by; count++ {
		tmp := s.Grid[row][s.Width-1]
		for i := s.Width - 1; i > 0; i-- {
			s.Grid[row][i] = s.Grid[row][i-1]
		}
		s.Grid[row][0] = tmp
	}
}

func (s *Screen) RotateCol(cmd string) {
	var col, by int
	fmt.Sscanf(cmd, "rotate column x=%d by %d", &col, &by)
	for count := 0; count < by; count++ {
		tmp := s.Grid[s.Height-1][col]
		for r := s.Height - 1; r > 0; r-- {
			s.Grid[r][col] = s.Grid[r-1][col]
		}
		s.Grid[0][col] = tmp
	}
}

func (s Screen) Display() {
	var char rune
	for _, row := range s.Grid {
		for _, pixel := range row {
			if pixel {
				char = '#'
			} else {
				char = '.'
			}
			fmt.Print(string(char))
		}
		fmt.Println()
	}
	fmt.Println()
}

func Day(input string, part int) int {
	s := &Screen{}
	s.Make(6, 50)
	for _, cmd := range strings.Split(input, "\n") {
		if strings.HasPrefix(cmd, "rect") {
			s.LightBox(cmd)
		} else if strings.HasPrefix(cmd, "rotate column") {
			s.RotateCol(cmd)
		} else if strings.HasPrefix(cmd, "rotate row") {
			s.RotateRow(cmd)
		}
		s.Display()
	}
	s.CountLit()
	return s.Lit
}
