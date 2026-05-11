package d02

import "strings"

func Day(input string, part int) (bathroomCode string) {
	var keypad [7][7]string
	x, y := 2, 4
	if part == 2 {
		x = 1
		y = 3
	}
	if part == 1 {
		keypad = [7][7]string{
			{".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "."},
			{".", "1", "2", "3", ".", ".", "."},
			{".", "4", "5", "6", ".", ".", "."},
			{".", "7", "8", "9", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "."},
		}
	} else {
		keypad = [7][7]string{
			{".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", "1", ".", ".", "."},
			{".", ".", "2", "3", "4", ".", "."},
			{".", "5", "6", "7", "8", "9", "."},
			{".", ".", "A", "B", "C", ".", "."},
			{".", ".", ".", "D", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "."},
		}
	}
	input = strings.TrimRight(input, "\n")
	for _, line := range strings.Split(input, "\n") {
		for _, char := range line {
			switch string(char) {
			case "U":
				if keypad[y-1][x] != "." {
					y--
				}
			case "D":
				if keypad[y+1][x] != "." {
					y++
				}
			case "L":
				if keypad[y][x-1] != "." {
					x--
				}
			case "R":
				if keypad[y][x+1] != "." {
					x++
				}
			}
		}
		bathroomCode += keypad[y][x]
	}
	return
}
