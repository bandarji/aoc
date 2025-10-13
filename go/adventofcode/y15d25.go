package adventofcode

import (
	"fmt"
	"log"
)

const (
	y15d25Message string = "To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d."
	y15d25Mult    int    = 252533
	y15d25Mod     int    = 33554393
)

func y15d25(input string, code int) int {
	row, col := y15d25GetRowCol(input)
	before := y15d25GetBefore(row, col)
	log.Println("row", row, "col", col, "before", before)
	for range before + col - 1 {
		code = (code * y15d25Mult) % y15d25Mod
	}
	return code
}

func y15d25GetRowCol(input string) (row, col int) {
	fmt.Sscanf(input, y15d25Message, &row, &col)
	return
}

func y15d25GetBefore(row, col int) (before int) {
	for i := 1; i <= row+col-2; i++ {
		before += i
	}
	return
}
