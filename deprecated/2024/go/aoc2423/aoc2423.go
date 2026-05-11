package aoc2423

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

const AOC2423_TEST string = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

type Connections map[string][]string

func ReadInput(input string) (adjacencies Connections) {
	adjacencies = Connections{}
	for _, line := range strings.Split(input, "\n") {
		computers := strings.Split(line, "-")
		c1, c2 := computers[0], computers[1]
		adjacencies[c1] = append(adjacencies[c1], c2)
		adjacencies[c2] = append(adjacencies[c2], c1)
	}
	return
}

func Aoc242301(input string) (count int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	adjacencies := ReadInput(input)
	// fmt.Printf("%+v\n", adjacencies)
	connections := Connections{}
	for k := range adjacencies {
		if !strings.HasPrefix(k, "t") {
			continue
		}
		for _, v1 := range adjacencies[k] {
			for _, v2 := range adjacencies[v1] {
				if slices.Contains(adjacencies[v2], k) {
					connection := []string{k, v1, v2}
					slices.Sort(connection)
					interconn := fmt.Sprintf("%s-%s-%s", connection[0], connection[1], connection[2])
					connections[interconn] = connection
				}
			}
		}
	}
	// fmt.Printf("%+v\n", connections)
	count = len(connections)
	return
}

func Aoc242302(input string) (pw string) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	adjacencies := ReadInput(input)
	_ = adjacencies
	return
}
