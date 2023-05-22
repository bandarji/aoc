package d01

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Santa struct {
	Facing  int // N = 0, E = 1, S = 2, W = 3
	X, Y    int
	Visited map[string]bool
	Part    int
}

func (s Santa) HasVisited(x, y int) bool {
	if x == 0 && y == 0 {
		return false
	}
	key := NormalizePosition(x, y)
	if _, found := s.Visited[key]; found {
		return true
	} else {
		return false
	}
}

func (s *Santa) Turn(turn string) {
	if turn == "L" {
		switch s.Facing {
		case 0:
			s.Facing = 3
		case 1:
			s.Facing = 0
		case 2:
			s.Facing = 1
		case 3:
			s.Facing = 2
		}
	} else {
		switch s.Facing {
		case 0:
			s.Facing = 1
		case 1:
			s.Facing = 2
		case 2:
			s.Facing = 3
		case 3:
			s.Facing = 0
		}
	}
}

func (s *Santa) Move() {
	x, y := s.X, s.Y
	switch s.Facing {
	case 0:
		y += 1
	case 1:
		x += 1
	case 2:
		y -= 1
	case 3:
		x -= 1
	}
	s.X = x
	s.Y = y
	key := NormalizePosition(x, y)
	if _, found := s.Visited[key]; found {
		if s.Part == 2 {
			log.Printf("Insection: %d, %d: ", x, y)
		}
	}
	s.Visited[key] = true
}

func (s Santa) Distance() int {
	return Abs(0, s.X) + Abs(0, s.Y)
}

func Abs(n1, n2 int) (d int) {
	d = n1 - n2
	if d < 0 {
		d *= -1
	}
	return
}

func NormalizePosition(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func Day(part int) int {
	content, err := os.ReadFile("inputs/input_day01.txt")
	if err != nil {
		log.Fatalln("I/O", err)
	}
	return WalkCity(string(content), part)
}

func WalkCity(input string, part int) int {
	santa := &Santa{0, 0, 0, map[string]bool{"0,0": true}, part}
	for _, i := range strings.Fields(input) {
		santa.Turn(string(i[0]))
		for m := 0; m < atoi(i); m++ {
			santa.Move()
		}
	}
	return santa.Distance()
}

func atoi(s string) (n int) {
	ns := ""
	for _, c := range s {
		if c >= 48 && c <= 57 {
			ns += string(c)
		}
	}
	n, _ = strconv.Atoi(ns)
	return
}
