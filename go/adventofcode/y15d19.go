package adventofcode

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

func y15d19(input string, part int) (answer int) {
	replacements, molecule := y15d19ReadInput(input)
	if part == 1 {
		answer = y15d19Part1(replacements, molecule)
	} else {
		answer = y15d19Part2(input)
	}
	return
}

func y15d19Part1(replacements [][2]string, molecule string) int {
	molecules := map[string]bool{}
	for i := 0; i < len(molecule); i++ {
		for _, replacement := range replacements {
			start, generated := replacement[0], replacement[1]
			if strings.HasPrefix(molecule[i:], start) {
				newMolecule := molecule[:i] + generated + molecule[i+len(start):]
				molecules[newMolecule] = true
			}
		}
	}
	// log.Println(molecules)
	return len(molecules)
}

func y15d19Part2(input string) (fewestSteps int) {
	mappings, molecule := y15d19ParseInputPart2(input)
	flipped := y15d19FlipMappings(mappings)
	yields := y15d19GetPossibleYields(flipped)
	fewestSteps = y15d19GetFewestSteps(molecule, flipped, yields)
	return
}

func y15d19GetPossibleYields(flippedMappings map[string]string) (possibleYields []string) {
	for prod := range flippedMappings {
		possibleYields = append(possibleYields, prod)
	}
	return
}

func y15d19FlipMappings(mappings map[string][]string) (flipped map[string]string) {
	flipped = map[string]string{}
	for base, yields := range mappings {
		for _, yield := range yields {
			flipped[yield] = base
		}
	}
	return
}

func y15d19ParseInputPart2(input string) (mappings map[string][]string, original string) {
	mappings = map[string][]string{}
	sections := strings.Split(input, "\n\n")
	original = sections[1]
	for line := range strings.SplitSeq(sections[0], "\n") {
		rule := strings.Split(line, " => ")
		base, yield := rule[0], rule[1]
		mappings[base] = append(mappings[base], yield)
	}
	return
}

func y15d19GetFewestSteps(original string, flipped map[string]string, possibleYields []string) (steps int) {
	startingMolecule := strings.Clone(original)
	changed := false
	for startingMolecule != "e" {
		changed = false
		for _, prod := range possibleYields {
			count := strings.Count(startingMolecule, prod)
			if count <= 0 {
				continue
			}
			changed = true
			steps += count
			// log.Println(startingMolecule, prod, flipped[prod])
			startingMolecule = strings.ReplaceAll(startingMolecule, prod, flipped[prod])
			break
		}
		if !changed {
			randNow := rand.New(rand.NewSource(time.Now().UnixNano()))
			sort.Slice(possibleYields, func(i, j int) bool { return randNow.Intn(2) == 0 })
			startingMolecule = strings.Clone(original)
			steps = 0
		}
	}
	return
}

func y15d19ReadInput(input string) (replacements [][2]string, molecule string) {
	sections := strings.Split(input, "\n\n")
	replacements = [][2]string{}
	for _, line := range strings.Split(sections[0], "\n") {
		parts := strings.Split(line, " => ")
		base, yield := parts[0], parts[1]
		replacements = append(replacements, [2]string{base, yield})
	}
	molecule = sections[1]
	return
}
