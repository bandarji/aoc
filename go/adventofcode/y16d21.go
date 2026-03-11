package adventofcode

import (
	"strings"
)

const y16d21Initial string = "abcdefgh"
const y16d21ScrabbledPassword string = "fbgdceah"

func y16d21(toScramble, instructions string, part int) string {
	password := []byte(toScramble)
	if part == 1 {
		return y16d21Process(password, instructions)
	} else {
		memo := map[string]bool{}
		for {
			password = []byte{}
			shuffledIndexes, key := shuffleIndexes(len(toScramble))
			// log.Println("shuffledIndexes", shuffledIndexes, "key", key)
			if memo[key] {
				continue
			}
			memo[key] = true
			for _, i := range shuffledIndexes {
				password = append(password, toScramble[i])
			}
			answer := y16d21Process(password, instructions)
			if answer == y16d21ScrabbledPassword {
				return string(password)
			}
		}
	}
}

func y16d21Process(password []byte, instructions string) string {
	for instruction := range strings.SplitSeq(instructions, "\n") {
		fields := strings.Fields(instruction)
		switch fields[0] {
		case "move":
			password = y16d21Move(password, fields)
		case "reverse":
			password = y16d21Reverse(password, fields)
		case "rotate":
			password = y16d21Rotate(password, fields)
		case "swap":
			password = y16d21Swap(password, fields)
		}
	}
	return string(password)
}

func y16d21CopySlice(slice []byte) []byte {
	return append([]byte{}, slice...)
}

func y16d21Move(password []byte, fields []string) []byte {
	temp := y16d21CopySlice(password)
	start, end := strToInt(fields[2]), strToInt(fields[5])
	temp[end] = password[start]
	offset := 0
	for i, c := range password {
		if i == start {
			offset--
			continue
		} else if i+offset == end {
			offset++
		}
		temp[i+offset] = c
	}
	return temp
}

func y16d21Reverse(password []byte, fields []string) []byte {
	temp := y16d21CopySlice(password)
	start, end := strToInt(fields[2]), strToInt(fields[4])
	k := end
	for i := start; i <= end; i++ {
		temp[i] = password[k]
		k--
	}
	return temp
}

func y16d21Rotate(password []byte, fields []string) []byte {
	temp := y16d21CopySlice(password)
	steps := 0
	switch fields[1] {
	case "left":
		steps = len(password) - strToInt(fields[2])
	case "right":
		steps = strToInt(fields[2])
	case "based":
		steps = 0
		for i, c := range password {
			if string(c) == fields[6] {
				steps = i
			}
		}
		if steps >= 4 {
			steps++
		}
		steps++
	}
	for i, c := range password {
		temp[(i+steps)%len(password)] = c
	}
	return temp
}

func y16d21Swap(password []byte, fields []string) []byte {
	temp := y16d21CopySlice(password)
	switch fields[1] {
	case "position":
		x, y := strToInt(fields[2]), strToInt(fields[5])
		temp[x], temp[y] = temp[y], temp[x]
	case "letter":
		x, y := fields[2][0], fields[5][0]
		for i, c := range password {
			switch c {
			case x:
				temp[i] = y
			case y:
				temp[i] = x
			default:
				temp[i] = c
			}
		}
	}
	return temp
}
