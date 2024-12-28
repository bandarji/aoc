package aoc2417

import (
	"fmt"
	"strings"
	"time"
)

const AOC2417_TEST string = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

type Computer struct {
	A, B, C, Ptr int
	Ins          []int
	O            []string
}

func BuildComputer(input string) (computer *Computer) {
	a, b, c, instructions := ReadProgram(input)
	return &Computer{A: a, B: b, C: c, Ptr: 0, Ins: instructions, O: []string{}}
}

func ReadProgram(input string) (a, b, c int, program []int) {
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "Register A: ") {
			fmt.Sscanf(line, "Register A: %d", &a)
		}
		if strings.HasPrefix(line, "Register B: ") {
			fmt.Sscanf(line, "Register B: %d", &b)
		}
		if strings.HasPrefix(line, "Register C: ") {
			fmt.Sscanf(line, "Register C: %d", &c)
		}
		if strings.HasPrefix(line, "Program: ") {
			sections := strings.Split(line, ": ")
			for _, v := range strings.Split(sections[1], ",") {
				var ins int
				fmt.Sscanf(v, "%d", &ins)
				program = append(program, ins)
			}
		}
	}
	return
}

func (computer *Computer) Run() string {
	opcode, operand, combo := 0, 0, 0
	a, b, c := computer.A, computer.B, computer.C
	ins := computer.Ins
	o := []string{}
	for p := 0; p < len(ins); p += 2 {
		opcode = ins[p]
		operand = ins[p+1]
		combo = operand
		switch operand {
		case 4:
			combo = a
		case 5:
			combo = b
		case 6:
			combo = c
		}
		switch opcode {
		case 0:
			a >>= combo
		case 1:
			b ^= operand
		case 2:
			b = combo % 8
		case 3:
			if a != 0 {
				p = operand - 2
			}
		case 4:
			b ^= c
		case 5:
			o = append(o, fmt.Sprintf("%d", combo%8))
		case 6:
			b = a >> combo
		case 7:
			c = a >> combo
		}
	}
	fmt.Printf("A: %d, B: %d, C: %d, O: %v\n", a, b, c, o)
	return strings.Join(o, ",")
}

func Aoc241701(input string) (output string) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	computer := BuildComputer(input)
	output = computer.Run()
	return
}

func Aoc241702(input string) (output string) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	return
}
