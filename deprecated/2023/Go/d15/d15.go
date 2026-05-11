package d15

import (
	"aoc2023/common"
	"strings"
)

const TEST1 string = `HASH`
const TEST2 string = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func ProcessSteps(input string, part int) (answer int) {
	steps := []string{}
	for _, step := range strings.Split(input, ",") {
		if part == 1 {
			answer += Hash(step)
		} else {
			steps = append(steps, step)
		}
	}
	if part == 2 {
		answer = ProcessLenses(steps)
	}
	return
}

func Hash(step string) (hash int) {
	for _, char := range step {
		hash += int(char)
		hash *= 17
		hash %= 256
	}
	return
}

func ProcessLenses(steps []string) (answer int) {
	boxes := make([][]string, 256)
	lengths := map[string]int{}
	label := ""
	index, length := 0, 0
	for _, step := range steps {
		if strings.ContainsRune(step, '-') {
			label = step[:len(step)-1]
			index = Hash(label)
			for i, l := range boxes[index] {
				if l == label {
					boxes[index] = append(boxes[index][:i], boxes[index][i+1:]...)
					break
				}
			}
		} else {
			sections := strings.Split(step, "=")
			label = sections[0]
			length = common.StrToInt(sections[1])
			index = Hash(label)
			if !common.SliceContainsString(boxes[index], label) {
				boxes[index] = append(boxes[index], label)
			}
			lengths[label] = length
		}
	}
	for number, box := range boxes {
		for slot, label := range box {
			answer += (number + 1) * (slot + 1) * lengths[label]
		}
	}
	return
}

func Solve(input string, part int) (answer int) {
	answer = ProcessSteps(input, part)
	return
}
