package adventofcode

import (
	"fmt"
	"strings"
)

type y17d07Node struct {
	weight int
	edges  []string
}

func y17d07(input string, part int) string {
	graph := y17d07ParseInput(input)
	names := y17d07MapNames(graph)
	part1 := y17d07Part1(graph, names)
	if part == 1 {
		return part1
	}
	return y17d07Part2(graph, part1)
}

func y17d07Part1(graph map[string]y17d07Node, names map[string]bool) string {
	for _, node := range graph {
		for _, name := range node.edges {
			delete(names, name)
		}
	}
	// log.Println(names)
	for k := range names {
		return k
	}
	return ""
}

func y17d07CalcWeights(graph map[string]y17d07Node) func(string) int {
	var insFunc func(string) int
	insFunc = func(root string) int {
		sum := graph[root].weight
		for _, dependent := range graph[root].edges {
			sum += insFunc(dependent)
		}
		return sum
	}
	return insFunc
}

func y17d07Part2(graph map[string]y17d07Node, current string) string {
	weightsFunc := y17d07CalcWeights(graph)
	neighbors := []string{}
	for i := 0; i < len(graph); i++ {
		downstreamWeights := map[int][]string{}
		for _, dependent := range graph[current].edges {
			weight := weightsFunc(dependent)
			downstreamWeights[weight] = append(downstreamWeights[weight], dependent)
		}
		if len(downstreamWeights) > 1 {
			neighbors = graph[current].edges
			for _, names := range downstreamWeights {
				if len(names) == 1 {
					current = names[0]
				}
			}
		} else if len(downstreamWeights) == 1 {
			currentWeight := weightsFunc(current)
			for _, neighb := range neighbors {
				if neighb != current {
					wanted := weightsFunc(neighb)
					return fmt.Sprintf("%d", graph[current].weight-(currentWeight-wanted))
				}
			}
		}
	}
	return fmt.Sprintf("%d", -1)
}

func y17d07ParseInput(input string) map[string]y17d07Node {
	graph := map[string]y17d07Node{}
	for line := range strings.SplitSeq(input, "\n") {
		left, right, found := strings.Cut(line, " -> ")
		fields := strings.Fields(left)
		name := fields[0]
		weight := strToInt(strings.Trim(fields[1], "()"))
		edges := []string{}
		if found {
			edges = strings.Split(right, ", ")
		}
		graph[name] = y17d07Node{
			weight: weight,
			edges:  edges,
		}
	}
	return graph
}

func y17d07MapNames(graph map[string]y17d07Node) map[string]bool {
	names := map[string]bool{}
	for name := range graph {
		names[name] = true
	}
	return names
}
