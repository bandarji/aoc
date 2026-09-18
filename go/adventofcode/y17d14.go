package adventofcode

import (
	"fmt"
	"strings"
)

const y17d14KeyString string = "ljoxqyyw"

func y17d14(input string, part int) (answer int) {
	w := []byte(input)
	grid := make([]string, 128)
	for i := range 128 {
		input := []byte(fmt.Sprintf("%s-%d", w, i))
		d := make([]int, 0, len(input)+5)
		for _, b := range input {
			d = append(d, int(b))
		}
		d = append(d, 17, 31, 73, 47, 23)
		grid[i] = y17d14KnotHash(d, 64)
	}
	total := 0
	for _, row := range grid {
		for _, c := range row {
			if c == '1' {
				total++
			}
		}
	}
	if part == 1 {
		answer = total
	} else {
		answer = y17d14CountRegions(grid)
	}
	return
}

func y17d14XorSlice(values []int) (result int) {
	for _, n := range values {
		result ^= n
	}
	return
}

func y17d14KnotHash(d []int, r int) string {
	sb := strings.Builder{}
	m := make([]int, 256)
	for i := range m {
		m[i] = i
	}
	pos, skip := 0, 0
	for round := 0; round < r; round++ {
		for _, length := range d {
			for o := 0; o < length/2; o++ {
				a := (pos + o) % 256
				b := (pos + length - o - 1) % 256
				m[a], m[b] = m[b], m[a]
			}
			pos += length + skip
			skip++
		}
	}
	for i := range 16 {
		block := make([]int, 16)
		for j := range 16 {
			block[j] = m[16*i+j]
		}
		value := y17d14XorSlice(block)
		sb.WriteString(fmt.Sprintf("%08b", value))
	}
	return sb.String()
}

func y17d14CountRegions(input []string) (count int) {
	grid := make([][]string, 128)
	for i, row := range input {
		grid[i] = strings.Split(row, "")
	}
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := range 128 {
		for j := range 128 {
			if grid[i][j] == "1" {
				q := [][2]int{{i, j}}
				for len(q) > 0 {
					cr, cc := q[0][0], q[0][1]
					grid[cr][cc] = "0"
					q = q[1:]
					for _, d := range dirs {
						nr, nc := cr+d[0], cc+d[1]
						if nr >= 0 && nr < 128 && nc >= 0 && nc < 128 && grid[nr][nc] == "1" {
							q = append(q, [2]int{nr, nc})
						}
					}
				}
				count++
			}
		}
	}
	return
}
