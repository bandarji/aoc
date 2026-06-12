package adventofcode

import "strings"

func y17d12(input string, part int) (answer int) {
	if part == 1 {
		answer = y17d12p1(y17d12ParseInput(input))
	} else {
		answer = y17d12p2(y17d12ParseInput(input))
	}
	return
}

func y17d12ParseInput(input string) map[int][]int {
	mappings := map[int][]int{}
	for line := range strings.SplitSeq(input, "\n") {
		pre, post, _ := strings.Cut(line, " <-> ")
		connections := []int{}
		for connection := range strings.SplitSeq(post, ", ") {
			connections = append(connections, strToInt(connection))
		}
		mappings[strToInt(pre)] = connections
	}
	return mappings
}

func y17d12p1(mappings map[int][]int) (count int) {
	programs := map[int]bool{}
	q := []int{0}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		_, ok := programs[p]
		if !ok {
			programs[p] = true
			q = append(q, mappings[p]...)
		}
	}
	count = len(programs)
	return
}

func y17d12p2(mappings map[int][]int) (groups int) {
	for program := range mappings {
		programs := map[int]bool{}
		groups++
		q := []int{program}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			_, ok := programs[p]
			if !ok {
				programs[p] = true
				q = append(q, mappings[p]...)
				delete(mappings, p)
			}
		}
	}
	return
}
