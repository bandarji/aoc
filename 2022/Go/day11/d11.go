package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	ID          int
	Inspections int
	Items       []int
	Left        string
	Right       string
	Operand     string
	Divisor     int
	True        int
	False       int
}

type Monkeys []Monkey

const INPUT_FILENAME string = "input_day11.txt"
const TEST_ANS1 int = 10605
const TEST_ANS2 int = 2713310158
const TEST_INPUT string = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`

func Day11(input string, rounds int, divide bool) (monkeyBiz int) {
	worry := 0
	destination := 0
	// msg := ""
	monkeys := ParseMonkeys(input)
	bigModulo := CalcBigModulo(monkeys)
	fmt.Println(monkeys)
	for r := 1; r <= rounds; r++ {
		for _, monkey := range monkeys {
			monkeys[monkey.ID].Inspections += len(monkey.Items)
			fmt.Println(fmt.Sprintf("Monkey %d:", monkey.ID))
			for _, item := range monkey.Items {
				worry = CalcWorry(monkey, item, divide, bigModulo)
				if worry%monkey.Divisor == 0 {
					destination = monkey.True
					// msg = "is"
				} else {
					destination = monkey.False
					// msg = "is not"
				}
				// fmt.Println(fmt.Sprintf("    Current worry level %s divisible by %d.", msg, monkey.Divisor))
				// fmt.Println(fmt.Sprintf("    Item with worry level %d is thrown to monkey %d.", worry, destination))
				monkeys[destination].Items = append(monkeys[destination].Items, worry)
			}
			monkeys[monkey.ID].Items = []int{}
		}
		ShowItems(monkeys, r)
	}
	monkeyBiz = ProductOfHighestTwoInspections(monkeys)
	return
}

func CalcBigModulo(monkeys Monkeys) int {
	n := 1
	for _, m := range monkeys {
		n *= m.Divisor
	}
	return n
}

func ParseMonkeys(input string) Monkeys {
	monkeys := make(Monkeys, 0)
	for _, monkeyInfo := range strings.Split(input, "\n\n") {
		monkeys = append(monkeys, parseMonkey(monkeyInfo))
	}
	return monkeys
}

func parseMonkey(input string) (m Monkey) {
	for _, line := range strings.Split(input, "\n") {
		if len(line) < 2 {
			continue
		}
		fields := strings.Fields(line)
		switch fields[0] {
		case "Monkey":
			m.Inspections = 0
			m.ID = atoi(fields[1])
		case "Starting":
			m.Items = assembleItems(fields[2:])
		case "Operation:":
			m.Left = fields[3]
			m.Operand = fields[4]
			m.Right = fields[5]
		case "Test:":
			m.Divisor = atoi(fields[len(fields)-1])
		case "If":
			if fields[1] == "true:" {
				m.True = atoi(fields[len(fields)-1])
			}
			if fields[1] == "false:" {
				m.False = atoi(fields[len(fields)-1])
			}

		}
	}
	return
}

func assembleItems(f []string) []int {
	items := make([]int, 0)
	for _, item := range f {
		items = append(items, atoi(item))
	}
	return items
}

func atoi(s string) int {
	ns := "" // number string
	for i := 0; i < len(s); i++ {
		if s[i] <= 57 && s[i] >= 48 {
			ns += string(s[i])
		}
	}
	num, _ := strconv.Atoi(ns)
	return num
}

func CalcWorry(m Monkey, item int, divide bool, bigModulo int) (value int) {
	var left, right int
	// fmt.Println(fmt.Sprintf("  Monkey inspects an item with a worry level of %d.", item))
	if m.Left == "old" {
		left = item
	} else {
		left = atoi(m.Left)
	}
	if m.Right == "old" {
		right = item
	} else {
		right = atoi(m.Right)
	}
	if m.Operand == "+" {
		value = left + right
		// fmt.Println(fmt.Sprintf("    Worry level increases to %d.", value))
	} else {
		value = left * right
		// fmt.Println(fmt.Sprintf("    Worry level is multiplied to %d.", value))
	}
	if divide {
		value /= 3
		// fmt.Println(fmt.Sprintf("    Monkey gets bored with item. Worry level is divided by 3 to %d.", value))
	} else {
		value %= bigModulo
	}
	return
}

func ShowItems(monkeys Monkeys, r int) {
	fmt.Println("\nAfter round", r, "the monkeys are holding items with these worry levels:")
	for _, m := range monkeys {
		fmt.Println("Monkey", m.ID, m.Inspections, "|", m.Items)
	}
}

func ProductOfHighestTwoInspections(monkeys Monkeys) (product int) {
	numbers := make([]int, len(monkeys))
	for i := 0; i < len(monkeys); i++ {
		numbers[i] = monkeys[i].Inspections
	}
	sort.Ints(numbers)
	product = numbers[len(numbers)-1] * numbers[len(numbers)-2]
	return
}

func RunTests(day int) {
	result := 0
	if day == 1 {
		result = Day11(TEST_INPUT, 20, true)
		if result != TEST_ANS1 {
			log.Fatal(fmt.Sprintf("Day 11 01 Test failed: %d instead of %d.", result, TEST_ANS1))
		}
	} else {
		result = Day11(TEST_INPUT, 10000, false)
		if result != TEST_ANS2 {
			log.Fatal(fmt.Sprintf("Day 11 01 Test failed: %d instead of %d.", result, TEST_ANS2))
		}
	}
}

func main() {
	fmt.Println("Day 11")
	RunTests(1)
	RunTests(2)
	input, _ := os.ReadFile(INPUT_FILENAME)
	fmt.Println("\nDay 11 01:", Day11(string(input), 20, true))
	fmt.Println("\nDay 11 02:", Day11(string(input), 10000, false))
}
