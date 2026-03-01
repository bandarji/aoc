package adventofcode

func y16d18(input string, rowCount int) (safeTiles int) {
	row := y16d18MakeRow(input)
	for i := 0; i < rowCount; i++ {
		safeTiles += y16d18CountSafeTiles(row)
		row = y16d18NextRow(row)
	}
	return
}

func y16d18CountSafeTiles(row []bool) (safeTiles int) {
	for _, trap := range row {
		if !trap {
			safeTiles++
		}
	}
	return
}

func y16d18NextRow(row []bool) (nextRow []bool) {
	nextRow = []bool{}
	for i := range row {
		l := !(i == 0 || !row[i-1])
		c := row[i]
		r := !(i == len(row)-1 || !row[i+1])
		eval := (l && c && !r) || (!l && c && r) || (l && !c && !r) || (!l && !c && r)
		nextRow = append(nextRow, eval)
	}
	return
}

func y16d18MakeRow(input string) (row []bool) {
	row = []bool{}
	for _, char := range input {
		row = append(row, char == '^')
	}
	return
}
