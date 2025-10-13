package adventofcode

import (
	"sort"
	"strings"
)

func y15d24(input string, part int) (answer int) {
	weights, targetWeight := y15d24ReadWeights(input, part)
	answer = y15d24FindQuantumEntanglement(weights, targetWeight)
	return
}

func y15d24ReadWeights(input string, part int) (weights []int, targetWeight int) {
	totalWeight := 0
	weights = []int{}
	for _, line := range strings.Split(input, "\n") {
		weight := strToInt(line)
		totalWeight += weight
		weights = append(weights, weight)
	}
	targetWeight = totalWeight / (part + 2)
	return
}

func y15d24FindQuantumEntanglement(weights []int, target int) (qe int) {
	groupings := [][]int{}
	for gLength := 2; len(groupings) == 0; gLength++ {
		groupings = combinationsInt(weights, gLength)
		goodGroupings := [][]int{}
		for _, group := range groupings {
			if y15d24AddWeights(group) == target {
				goodGroupings = append(goodGroupings, group)
			}
		}
		groupings = goodGroupings
	}
	qeGroup := y15d24FindLowestQuantumEntanglement(groupings)
	qe = y15d24CalcQuantumEntanglement(qeGroup)
	return
}

func y15d24AddWeights(group []int) (sum int) {
	for _, weight := range group {
		sum += weight
	}
	return
}

func y15d24CalcQuantumEntanglement(group []int) (qe int) {
	qe = 1
	for _, weight := range group {
		qe *= weight
	}
	return
}

func y15d24FindLowestQuantumEntanglement(groupings [][]int) (qeGroup []int) {
	sort.Slice(groupings, func(i, j int) bool {
		return y15d24CalcQuantumEntanglement(groupings[i]) < y15d24CalcQuantumEntanglement(groupings[j])
	})
	qeGroup = groupings[0]
	return
}
