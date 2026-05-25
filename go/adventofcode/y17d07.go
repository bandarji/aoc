package adventofcode

import (
	"strings"
)

type y17d07Node struct {
	weight int
	edges  []string
}

func y17d07(input string, part int) (name string) {
	graph := y17d07ParseInput(input)
	names := y17d07MapNames(graph)
	if part == 1 {
		name = y17d07Part1(graph, names)
	} else {
		name = ""
	}
	return
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
