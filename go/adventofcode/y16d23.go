package adventofcode

import (
	"log"
	"strings"
)

type y16d23Computer struct {
	registers    map[string]int
	instructions []string
	cursor, size int
}

func (c *y16d23Computer) executeInstruction() {
	instruction := c.instructions[c.cursor]
	fields := strings.Fields(instruction)
	switch fields[0] {
	case "cpy":
		c.cpy(fields[1], fields[2])
	case "inc":
		c.inc(fields[1])
	case "dec":
		c.dec(fields[1])
	case "jnz":
		c.jnz(fields[1], fields[2])
	case "tgl":
		c.tgl(fields[1])
	}
}

func (c *y16d23Computer) tgl(i string) {
	x := y16d23ValueOrRegister(c.registers, i)
	mi := c.cursor + x
	if mi < c.size {
		newOp := ""
		mfields := strings.Fields(c.instructions[mi])
		if len(mfields) == 2 {
			newOp = "inc"
			if mfields[0] == "inc" {
				newOp = "dec"
			}
		}
		if len(mfields) == 3 {
			newOp = "jnz"
			if mfields[0] == "jnz" {
				newOp = "cpy"
			}
		}
		mfields[0] = newOp
		c.instructions[mi] = strings.Join(mfields, " ")
	}
	c.cursor++
}

func (c *y16d23Computer) cpy(i1, i2 string) {
	c.registers[i2] = y16d23ValueOrRegister(c.registers, i1)
	c.cursor++
}

func (c *y16d23Computer) inc(i string) {
	c.registers[i]++
	c.cursor++
}

func (c *y16d23Computer) dec(i string) {
	c.registers[i]--
	c.cursor++
}

func (c *y16d23Computer) jnz(i1, i2 string) {
	if y16d23ValueOrRegister(c.registers, i1) != 0 {
		c.cursor += y16d23ValueOrRegister(c.registers, i2)
	} else {
		c.cursor++
	}
}

var y16d23Registers = map[string]int{"a": 7, "b": 0, "c": 0, "d": 0}

func y16d23(input string, registers map[string]int) int {
	instructions := strings.Split(input, "\n")
	p := &y16d23Computer{
		registers:    registers,
		instructions: instructions,
		cursor:       0,
		size:         len(instructions),
	}
	i := 0
	for p.cursor < p.size {
		i++
		p.executeInstruction()
		if i%50_000 == 0 {
			log.Println("i", i, "cursor", p.cursor, "registers", p.registers)
		}
	}
	return p.registers["a"]
}

func y16d23ValueOrRegister(registers map[string]int, insPart string) (response int) {
	if strings.ContainsAny(insPart, "abcd") {
		response = registers[insPart]
	} else {
		response = strToInt(insPart)
	}
	return
}
