package adventofcode

import (
	"fmt"
	"strings"
)

func y17d24(input string, part int) (answer int) {
	interposers := y17d24ParseInput(input)
	bridges := y17d24Bridge(0, interposers, part, 0)
	if len(bridges) > 0 {
		answer = bridges[len(bridges)-1][2]
	}
	return
}

func y17d24Bridge(start int, interposers [][2]int, part, score int) (best [][3]int) {
	best = [][3]int{}
	for _, interposer := range interposers {
		if interposer[0] == start || interposer[1] == start {
			bag := y17d24CloneInterposers(interposers)
			bag = y17d24RemoveInterposer(bag, interposer)
			thisScore := interposer[0] + interposer[1] + score
			thisStart := interposer[0]
			if start == interposer[0] {
				thisStart = interposer[1]
			}
			candidate := [][3]int{{interposer[0], interposer[1], thisScore}}
			candidate = append(candidate, y17d24Bridge(thisStart, bag, part, thisScore)...)
			// log.Printf("candidate: %v", candidate)
			if len(best) == 0 {
				best = candidate
			} else {
				if part == 1 {
					if candidate[len(candidate)-1][2] > best[len(best)-1][2] {
						best = candidate
					}
				} else {
					if len(candidate) > len(best) || (len(candidate) == len(best) && candidate[len(candidate)-1][2] > best[len(best)-1][2]) {
						best = candidate
					}
				}
			}
		}
	}
	return
}

func y17d24RemoveInterposer(interposers [][2]int, interposer [2]int) (remaining [][2]int) {
	remaining = [][2]int{}
	searches := y17d24InterposerKeys(interposer)
	for i, element := range interposers {
		keys := y17d24InterposerKeys(element)
		if keys[0] == searches[0] || keys[0] == searches[1] || keys[1] == searches[0] || keys[1] == searches[1] {
			remaining = append(remaining, append(interposers[:i], interposers[i+1:]...)...)
			break
		}
	}
	return
}

func y17d24InterposerKeys(interposer [2]int) (keys [2]string) {
	keys[0] = fmt.Sprintf("%d/%d", interposer[0], interposer[1])
	keys[1] = fmt.Sprintf("%d/%d", interposer[1], interposer[0])
	return
}

func y17d24CloneInterposers(interposers [][2]int) (clone [][2]int) {
	for _, interposer := range interposers {
		clone = append(clone, [2]int{interposer[0], interposer[1]})
	}
	return
}

func y17d24ParseInput(input string) (interposers [][2]int) {
	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Split(line, "/")
		interposers = append(interposers, [2]int{strToInt(parts[0]), strToInt(parts[1])})
	}
	return
}
