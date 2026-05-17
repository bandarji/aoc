package adventofcode

import "math"

func y17d03(input string, part int) (answer int) {
	if part == 1 {
		answer = y17d03Solve1(strToInt(input))
	} else {
		answer = y17d03Solve2(strToInt(input))
	}
	return
}

func y17d03Solve2(target int) int {
	dirs := [4][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
	nd := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	posSum := map[[2]int]int{{0, 0}: 1, {0, 1}: 1}
	facing := 0
	row, col := 0, 1
	for len(posSum) < target {
		ld := dirs[(facing+1)%4]
		left := [2]int{row + ld[0], col + ld[1]}
		if _, ok := posSum[left]; !ok {
			facing = (facing + 1) % 4
		}
		diff := dirs[facing]
		row += diff[0]
		col += diff[1]
		next := [2]int{row, col}
		sum := 0
		for _, d := range nd {
			sum += posSum[[2]int{row + d[0], col + d[1]}]
		}
		if sum > target {
			return sum
		}
		posSum[next] = sum
	}
	return -1
}

func y17d03Solve1(target int) int {
	n, l := 0, 0
	n = int(math.Ceil(math.Sqrt(float64(target))))
	if n%2 == 0 {
		l = n + 1
	} else {
		l = n
	}
	min := l*l - l + 1
	max := l * l
	mid := min + (max-min)/2
	drift := int(math.Abs(float64(target - mid)))
	return l/2 + drift
}
