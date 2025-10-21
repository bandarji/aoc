package adventofcode

import "strings"

type y16d02Code struct {
	code     strings.Builder
	pad      []string
	row, col int
}

func (c *y16d02Code) add(key rune) {
	c.code.WriteRune(key)
}

func (c *y16d02Code) move(direction rune) {
	if direction == 'U' && c.row > 0 && c.pad[c.row-1][c.col] != ' ' {
		c.row--
	} else if direction == 'D' && c.row < len(c.pad)-1 && c.pad[c.row+1][c.col] != ' ' {
		c.row++
	} else if direction == 'L' && c.col > 0 && c.pad[c.row][c.col-1] != ' ' {
		c.col--
	} else if direction == 'R' && c.col < len(c.pad[0])-1 && c.pad[c.row][c.col+1] != ' ' {
		c.col++
	}
}

var (
	y16d02Pad1 = []string{
		"     ",
		" 123 ",
		" 456 ",
		" 789 ",
		"     ",
	}
	y16d02Pad2 = []string{
		"       ",
		"   1   ",
		"  234  ",
		" 56789 ",
		"  ABC ",
		"   D   ",
		"      ",
	}
)

func y16d02(input string, part int) string {
	keypad := y16d02Code{code: strings.Builder{}, pad: y16d02Pad1, row: 2, col: 2}
	if part == 2 {
		keypad = y16d02Code{code: strings.Builder{}, pad: y16d02Pad2, row: 3, col: 1}
	}
	for _, line := range strings.Split(input, "\n") {
		for _, direction := range line {
			keypad.move(direction)
		}
		keypad.add(rune(keypad.pad[keypad.row][keypad.col]))
	}
	return keypad.code.String()
}
