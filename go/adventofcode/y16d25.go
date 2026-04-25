package adventofcode

import (
	"fmt"
	"log"
	"strings"
)

func y16d25(input string) int {
	for a := 1; a < 1_000_000; a++ {
		if y16d25Run(input, map[string]int{"a": a}) {
			return a
		}
	}
	return -1
}

func y16d25ConstructDesiredOutput() string {
	sb := strings.Builder{}
	for range 100 {
		sb.WriteString("01")
	}
	return sb.String()
}

func y16d25Run(input string, registers map[string]int) bool {
	outs := ""
	outsDesired := y16d25ConstructDesiredOutput()
	outsMinLen := len(outsDesired)
	instructions := strings.Split(input, "\n")
	cursor := 0
	for cursor < len(instructions) {
		instruction := instructions[cursor]
		fields := strings.Fields(instruction)
		switch fields[0] {
		case "cpy":
			registers[fields[2]] = y16d23ValueOrRegister(registers, fields[1])
		case "inc":
			registers[fields[1]]++
		case "dec":
			registers[fields[1]]--
		case "jnz":
			if y16d23ValueOrRegister(registers, fields[1]) != 0 {
				cursor += y16d23ValueOrRegister(registers, fields[2])
			} else {
				cursor++
			}
			continue
		case "out":
			reg := y16d23ValueOrRegister(registers, fields[1])
			if reg != 0 && reg != 1 {
				return false
			}
			outs += fmt.Sprintf("%d", reg)
			if len(outs) > 2 && outs[0] != '0' && outs[1] != '1' {
				return false
			}
			if len(outs) == outsMinLen {
				if outs == outsDesired {
					return true
				} else {
					return false
				}
			}
		case "tgl":
			target := cursor + y16d23ValueOrRegister(registers, fields[1])
			if target < len(instructions) {
				targetFields := strings.Fields(instructions[target])
				switch len(targetFields) {
				case 2:
					if targetFields[0] == "inc" {
						targetFields[0] = "dec"
					} else {
						targetFields[0] = "inc"
					}
				case 3:
					if targetFields[0] == "jnz" {
						targetFields[0] = "cpy"
					} else {
						targetFields[0] = "jnz"
					}
				}
				instructions[target] = strings.Join(targetFields, " ")
			}
		default:
			log.Println("unknown instruction", instruction)
		}
		cursor++
	}
	return false
}
