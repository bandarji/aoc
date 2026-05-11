package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	TEST_INPUT string = `root: pppw + sjmn
	dbpl: 5
	cczh: sllz + lgvd
	zczc: 2
	ptdq: humn - dvpt
	dvpt: 3
	lfqf: 4
	humn: 5
	ljgn: 2
	sjmn: drzm * dbpl
	sllz: 4
	pppw: cczh / lfqf
	lgvd: ljgn * ptdq
	drzm: hmdt - zczc
	hmdt: 32`
	TEST_ANS1 int = 152
	TEST_ANS2 int = 301
)

type Monkey struct {
	Name, Left, Operand, Right string
	Value                      int
}

func Day2101(input string) (root int) {
	known, monkeys := AssembleMonkeys(input)
	i := 0
	for {
		i++
		// log.Printf("\nStart loop %05d\n", i)
		for name, m := range monkeys {
			ValueLeft, okLeft := known[m.Left]
			ValueRight, okRight := known[m.Right]
			if okLeft && okRight {
				switch m.Operand {
				case "+":
					m.Value = ValueLeft + ValueRight
				case "-":
					m.Value = ValueLeft - ValueRight
				case "*":
					m.Value = ValueLeft * ValueRight
				case "/":
					m.Value = ValueLeft / ValueRight
				}
				// log.Printf("Added %s (%d) to known\n", name, m.Value)
				known[name] = m.Value
			}
		}
		yell, ok := known["root"]
		if ok {
			root = yell
			break
		}
	}
	return
}

func AssembleMonkeys(input string) (map[string]int, map[string]Monkey) {
	known := map[string]int{}
	monkeys := make(map[string]Monkey)
	input = strings.TrimRight(input, "\n")
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		name := string(fields[0][0:4])
		if len(fields) == 2 {
			value, _ := strconv.Atoi(fields[1])
			monkeys[name] = Monkey{Name: name, Left: "", Operand: "", Right: "", Value: value}
			known[name] = value
		} else {
			monkeys[name] = Monkey{Name: name, Left: fields[1], Operand: fields[2], Right: fields[3], Value: 0}
		}
	}
	return known, monkeys
}

func main() {
	day, part, answer := 21, 1, 0
	answer = Day2101(TEST_INPUT)
	if answer != TEST_ANS1 {
		log.Fatalf("Day %02d / %02d test failed: %d vs %d\n", day, part, answer, TEST_ANS1)
	}
	log.Println("Tests pass!")
	content, _ := os.ReadFile(fmt.Sprintf("input_day%02d.txt", day))
	answer = Day2101(string(content))
	log.Printf("Day %02d / %02d answer = %d\n", day, part, answer)
}
