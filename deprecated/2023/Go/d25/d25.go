package d25

import (
	"strings"
)

const TEST1 string = `jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr`

type Line struct {
	src, dst string
}

func GetNodes(input string) (nodes map[string][]string) {
	nodes = map[string][]string{}
	for _, line := range strings.Split(input, "\n") {
		f := strings.Fields(line)
		source := f[0][0:3]
		if _, exists := nodes[source]; !exists {
			nodes[source] = []string{}
		}
		for _, destination := range f[1:] {
			nodes[source] = append(nodes[source], destination)
			if _, exists := nodes[destination]; !exists {
				nodes[destination] = []string{}
			}
			nodes[destination] = append(nodes[destination], source)
		}
	}
	return
}

func TraceNodes(nodes map[string][]string, source string, counts map[Line]int) {
	visited := map[string]bool{}
	q := []string{source}
	for len(q) > 0 {
		start := q[0]
		q = q[1:]
		for _, next := range nodes[start] {
			if _, exists := visited[next]; exists {
				continue
			}
			q = append(q, next)
			visited[next] = true
			line := Line{}
			if start < next {
				line = Line{start, next}
			} else {
				line = Line{next, start}
			}
			counts[line]++
		}
	}
}

func CountLines(nodes map[string][]string) (counts map[Line]int) {
	counts = map[Line]int{}
	giveUp := len(nodes) * 2
	i := 0
	for source := range nodes {
		TraceNodes(nodes, source, counts)
		i++
		if i > giveUp {
			break
		}
	}
	return

}

func RemoveMax(counts map[Line]int) (maxLine Line) {
	maxLine = Line{}
	hi := 0
	for k, v := range counts {
		if v > hi {
			hi = v
			maxLine = k
		}
	}
	delete(counts, maxLine)
	return
}

func DropLine(nodes map[string][]string, line Line) {
	strings := []string{}
	for _, v := range nodes[line.src] {
		if v != line.dst {
			strings = append(strings, v)
		}
	}
	nodes[line.src] = strings
	strings = []string{}
	for _, v := range nodes[line.dst] {
		if v != line.src {
			strings = append(strings, v)
		}
	}
	nodes[line.dst] = strings
}

func CountNodes(nodes map[string][]string, start string) int {
	visited := map[string]bool{}
	q := []string{start}
	for len(q) > 0 {
		from := q[0]
		q = q[1:]
		for _, to := range nodes[from] {
			if _, exists := visited[to]; exists {
				continue
			}
			q = append(q, to)
			visited[to] = true
		}
	}
	return len(visited)
}

func Solve(input string, part int) (answer int) {
	nodes := GetNodes(input)

	count := CountLines(nodes)
	cut1 := RemoveMax(count)
	DropLine(nodes, cut1)
	count = CountLines(nodes)
	cut2 := RemoveMax(count)
	DropLine(nodes, cut2)
	count = CountLines(nodes)
	cut3 := RemoveMax(count)
	DropLine(nodes, cut3)

	nodeCount := CountNodes(nodes, cut1.src)
	answer = (len(nodes) - nodeCount) * nodeCount

	// for k, v := range nodes {
	// 	log.Printf("k=%s / v=%+v", k, v)
	// }
	return
}
