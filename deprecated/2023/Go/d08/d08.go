package d08

import (
	"aoc2023/common"
	"fmt"
	"strings"
)

const TEST1 string = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const TEST2 string = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

const TEST3 string = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

type Node struct {
	Left, Right string
}

type Network struct {
	LeftRight string
	Nodes     map[string]Node
}

func (n *Network) ParseInput(input string) {
	var step, left, right string
	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			n.LeftRight = line
		}
		if !strings.Contains(line, "=") {
			continue
		}
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &step, &left, &right)
		n.Nodes[step] = Node{Left: left, Right: right}
	}
}

func StepsToZZZ(n Network) int {
	steps := 0
	step := "AAA"
	var node Node
	for {
		for _, char := range strings.Split(n.LeftRight, "") {
			steps++
			node = n.Nodes[step]
			if char == "L" {
				step = node.Left
			} else {
				step = node.Right
			}
			if step == "ZZZ" {
				return steps
			}
		}
	}
}

func StepsToSecondZ(n Network, start string) int {
	found, steps := 1, 0
	step := start
	var node Node
	for {
		for _, char := range strings.Split(n.LeftRight, "") {
			steps++
			node = n.Nodes[step]
			if char == "L" {
				step = node.Left
			} else {
				step = node.Right
			}
			if strings.HasSuffix(step, "Z") {
				found++
			}
			if found == 2 {
				return steps
			}
		}
	}
}

func GhostSteps(n Network) (answer int) {
	cycles := []int{}
	for k := range n.Nodes {
		if strings.HasSuffix(k, "A") {
			cycles = append(cycles, StepsToSecondZ(n, k))
		}
	}
	lcm := cycles[0]
	for _, num := range cycles {
		lcm = lcm * num / common.GCD(lcm, num)
	}
	answer = lcm
	return
}

func Solve(input string, part int) (answer int) {
	network := &Network{
		LeftRight: "",
		Nodes:     map[string]Node{},
	}
	network.ParseInput(input)
	if part == 1 {
		answer = StepsToZZZ(*network)
	} else {
		answer = GhostSteps(*network)
	}
	return
}
