package d09

import (
	"aoc2023/common"
	"strings"
)

type History []int

func (h History) AllZeroes() bool {
	for _, n := range h {
		if n != 0 {
			return false
		}
	}
	return true
}

const TEST1 string = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func ReadHistories(input string) (histories []History) {
	var numbers History
	histories = []History{}
	for _, line := range strings.Split(input, "\n") {
		numbers = []int{}
		for _, entry := range strings.Fields(line) {
			numbers = append(numbers, common.StrToInt(entry))
		}
		histories = append(histories, numbers)
	}
	return
}

func MakePrediction(sequence History) (int, int) {
	var prefix, suffix int
	deltas := []int{}
	for i := 1; i < len(sequence); i++ {
		deltas = append(deltas, sequence[i]-sequence[i-1])
	}
	if sequence.AllZeroes() {
		return sequence[0], sequence[len(sequence)-1]
	}
	prefix, suffix = MakePrediction(deltas)
	return sequence[0] - prefix, sequence[len(sequence)-1] + suffix
}

func Solve(input string, part int) (answer int) {
	histories := ReadHistories(input)
	for _, history := range histories {
		prefix, suffix := MakePrediction(history)
		if part == 1 {
			answer += suffix
		} else {
			answer += prefix
		}
	}
	return
}
