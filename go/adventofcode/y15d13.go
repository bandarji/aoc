package adventofcode

import (
	"fmt"
	"strings"
)

func y15d13(input string, part int) (totalChangeInHappiness int) {
	seating := y15d13ParseInput(input)
	people := y15d13ListPeople(seating)
	perms := [][]string{}
	if part == 2 {
		people = append(people, ME)
		seating[ME] = map[string]int{}
	}
	permutations(people, len(people), &perms)
	for _, persons := range perms {
		totalChangeInHappiness = max(totalChangeInHappiness, y15d13CalculateChangeInHappiness(persons, seating))
	}
	return
}

func y15d13CalculateChangeInHappiness(persons []string, seating map[string]map[string]int) int {
	changeInHappiness := 0
	for i := 0; i < len(persons); i++ {
		changeInHappiness += seating[persons[i]][persons[(i+1)%len(persons)]]
		changeInHappiness += seating[persons[i]][persons[(i-1+len(persons))%len(persons)]]
	}
	return changeInHappiness
}

func y15d13ListPeople(seating map[string]map[string]int) []string {
	people := []string{}
	for person := range seating {
		people = append(people, person)
	}
	return people
}

func y15d13ParseInput(input string) (seating map[string]map[string]int) {
	seating = map[string]map[string]int{}
	for _, line := range strings.Split(input, "\n") {
		person, nextTo, amount := y15d13ParseLine(line)
		if _, ok := seating[person]; !ok {
			seating[person] = map[string]int{}
		}
		seating[person][nextTo] = amount
	}
	return
}

func y15d13ParseLine(line string) (person, nextTo string, amount int) {
	gainLoss := ""
	fmt.Sscanf(strings.Trim(line, "."), "%s would %s %d happiness units by sitting next to %s", &person, &gainLoss, &amount, &nextTo)
	if gainLoss == "lose" {
		amount = -amount
	}
	return
}
