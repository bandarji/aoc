package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	TEST_INPUT string = "London to Dublin = 464\nLondon to Belfast = 518\nDublin to Belfast = 141"
	TEST_ANS1  int    = 605
	TEST_ANS2  int    = 982
)

type Map map[string]map[string]int
type Places map[string]bool

func (m *Map) AddDistance(input string) {
	from, to := "", ""
	dist := 0
	fmt.Sscanf(input, "%s to %s = %d", &from, &to, &dist)
	_, hasFrom := (*m)[from]
	if !hasFrom {
		(*m)[from] = map[string]int{}
	}
	(*m)[from][to] = dist
}

func (ps *Places) AddPlace(p string) {
	(*ps)[p] = true
}

func Day09(input string, part int) (dist int) {
	distances := AssembleDistances(input)
	places := AssemblePlaces(input)
	var perms [][]string
	Permutations(places, len(places), &perms)
	if part == 1 {
		dist = ShortestDistance(perms, distances)
	} else {
		dist = LongestDistance(perms, distances)
	}
	return
}

func FindDistance(distances Map, p1, p2 string) int {
	dist, ok := distances[p1][p2]
	if !ok {
		dist, ok := distances[p2][p1]
		if !ok {
			return -1
		} else {
			return dist
		}
	}
	return dist
}

func LongestDistance(perms [][]string, distances Map) (longest int) {
	for _, perm := range perms {
		total := 0
		for i := 1; i < len(perm); i++ {
			from, to := perm[i-1], perm[i]
			dist := FindDistance(distances, from, to)
			if dist < 0 {
				total -= 100000
			} else {
				total += dist
			}
		}
		if total > longest {
			longest = total
		}
	}
	return
}

func ShortestDistance(perms [][]string, distances Map) int {
	shortest := 1 << 31
	for _, perm := range perms {
		total := 0
		// log.Printf("%s -> %s\n", perm[0], perm[len(perm)-1])
		for i := 1; i < len(perm); i++ {
			from, to := perm[i-1], perm[i]
			dist := FindDistance(distances, from, to)
			if dist < 0 {
				total += 100000
			} else {
				total += dist
			}
		}
		if total < 100000 {
			log.Printf("%s -> %s: %d\n", perm[0], perm[len(perm)-1], total)
		}
		if total < shortest {
			shortest = total
		}
	}
	return shortest
}

func AssemblePlaces(input string) (spots []string) {
	places := &Places{}
	input = strings.TrimRight(input, "\n")
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		places.AddPlace(fields[0])
		places.AddPlace(fields[2])
	}
	for place, _ := range *places {
		spots = append(spots, place)
	}
	return
}

func AssembleDistances(input string) map[string]map[string]int {
	distances := Map{}
	input = strings.TrimRight(input, "\n")
	for _, line := range strings.Split(input, "\n") {
		distances.AddDistance(line)
	}
	return distances
}

func AddDistance(distances *Map, input string) {
	from, to := "", ""
	dist := 0
	fmt.Sscanf(input, "%s to %s = %d", &from, &to, &dist)
	_, hasFrom := (*distances)[from]
	if !hasFrom {
		(*distances)[from] = map[string]int{}
	}
	(*distances)[from][to] = dist
}

func Permutations(slice []string, n int, result *[][]string) {
	if n == 1 {
		dst := make([]string, len(slice))
		copy(dst, slice[:])
		*result = append(*result, dst)
	} else {
		for i := 0; i < n; i++ {
			Permutations(slice, n-1, result)
			if n%2 == 0 {
				slice[0], slice[n-1] = slice[n-1], slice[0]
			} else {
				slice[i], slice[n-1] = slice[n-1], slice[i]
			}
		}
	}
}

func RunTests(day int) {
	dist := Day09(TEST_INPUT, 1)
	log.Printf("Test %02d 01: expect=%d, received=%d\n", day, TEST_ANS1, dist)
	dist = Day09(TEST_INPUT, 2)
	log.Printf("Test %02d 01: expect=%d, received=%d\n", day, TEST_ANS2, dist)
}

func main() {
	day, part := 9, 1
	RunTests(day)
	content, _ := os.ReadFile("input_day09.txt")
	answer := Day09(string(content), part)
	log.Printf("[%02d / %02d]: %d\n", day, part, answer)
	part = 2
	answer = Day09(string(content), part)
	log.Printf("[%02d / %02d]: %d\n", day, part, answer)
}
