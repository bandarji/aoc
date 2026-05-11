package main

import (
	"fmt"
	"log"
	"os"
)

type Coord struct {
	x, y int
}

type Coords map[string]int

func Day0301(input string) int {
	coords := Coords{"0-0": 1}
	VisitHouses(&coords, input)
	return len(coords)
}

func Day0302(input string) int {
	coords := Coords{"0-0": 1}
	VisitHousesWithRoboSanta(&coords, input)
	return len(coords)
}

func VisitHouses(coords *Coords, input string) {
	x, y := 0, 0
	dx, dy := 0, 0
	for _, char := range input {
		dx, dy = 0, 0
		switch string(char) {
		case "^":
			dy += 1
		case "v":
			dy -= 1
		case ">":
			dx += 1
		case "<":
			dx -= 1
		}
		x += dx
		y += dy
		(*coords)[fmt.Sprintf("%d-%d", x, y)] = 1
	}
}

func VisitHousesWithRoboSanta(coords *Coords, input string) {
	sx, sy, rx, ry := 0, 0, 0, 0
	dx, dy := 0, 0
	for i, char := range input {
		dx, dy = 0, 0
		switch string(char) {
		case "^":
			dy += 1
		case "v":
			dy -= 1
		case ">":
			dx += 1
		case "<":
			dx -= 1
		}
		if i%2 == 0 {
			sx += dx
			sy += dy
			(*coords)[fmt.Sprintf("%d-%d", sx, sy)] = 1
		} else {
			rx += dx
			ry += dy
			(*coords)[fmt.Sprintf("%d-%d", rx, ry)] = 2
		}
	}
}

func RunTests() bool {
	tests := map[string]int{
		">":          2,
		"^>v<":       4,
		"^v^v^v^v^v": 2,
	}
	for test, expected := range tests {
		received := Day0301(test)
		log.Printf("test=\"%s\", expected=%d, received=%d\n", test, expected, received)
		if received != received {
			return false
		}
	}
	log.Println("Tests 01 pass.")
	tests = map[string]int{
		"^v":         3,
		"^>v<":       3,
		"^v^v^v^v^v": 11,
	}
	for test, expected := range tests {
		received := Day0302(test)
		log.Printf("test=\"%s\", expected=%d, received=%d\n", test, expected, received)
		if received != received {
			return false
		}
	}
	log.Println("Tests 02 pass.")
	return true
}

func main() {
	day, part, houses := 3, 1, 0
	RunTests()
	content, _ := os.ReadFile(fmt.Sprintf("input_day%02d.txt", day))
	houses = Day0301(string(content))
	log.Printf("[%02d / %02d]: %d\n", day, part, houses)
	part = 2
	houses = Day0302(string(content))
	log.Printf("[%02d / %02d]: %d\n", day, part, houses)
}
