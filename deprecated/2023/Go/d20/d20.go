package d20

import (
	"strings"
)

const TEST1 string = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

const TEST2 string = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

type Module struct {
	kind    byte
	outputs []string
	state   bool
	memory  map[string]bool
}

func ReadConfiguration(input string) (modules map[string]Module) {
	modules = map[string]Module{}
	modules["button"] = Module{'.', []string{"roadcaster"}, false, map[string]bool{}}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		modules[parts[0][1:]] = Module{parts[0][0], strings.Split(parts[1], ", "), false, map[string]bool{}}
	}
	for name, mod := range modules {
		for _, t := range mod.outputs {
			dst := modules[t]
			if dst.kind == '&' {
				dst.memory[name] = false
			}
		}
	}
	return
}

type Pulse struct {
	src, dst string
	hi       bool
}

func (p Pulse) Process(modules map[string]Module) []Pulse {
	out := false
	m := modules[p.dst]
	if m.kind == '%' {
		if p.hi {
			return []Pulse{}
		} else {
			m.state = !m.state
			out = m.state
		}
	}
	if m.kind == '&' {
		m.memory[p.src] = p.hi
		out = false
		for _, v := range m.memory {
			if !v {
				out = true
				break
			}
		}
	}
	if m.kind == 'b' {
		out = p.hi
	}
	outputs := []Pulse{}
	for _, output := range m.outputs {
		outputs = append(outputs, Pulse{p.dst, output, out})
	}
	modules[p.dst] = m
	return outputs
}

func Button(modules map[string]Module, cycles int) (answer int) {
	lo, hi := 0, 0
	q := []Pulse{}
	for i := 0; i < cycles; i++ {
		q = append(q, Pulse{"button", "roadcaster", false})
		for len(q) > 0 {
			pulse := q[0]
			q = q[1:]
			if pulse.hi {
				hi++
			} else {
				lo++
			}
			q = append(q, pulse.Process(modules)...)
		}
	}
	answer = lo * hi
	return
}

func Inputs(modules map[string]Module, dst string) (inputs []string) {
	inputs = []string{}
	for name, mod := range modules {
		for _, o := range mod.outputs {
			if o == dst {
				inputs = append(inputs, name)
			}
		}
	}
	return
}

func FirstRX(modules map[string]Module) (answer int) {
	rxInputs := Inputs(modules, "rx")
	inputs := Inputs(modules, rxInputs[0])
	factors := map[string]int{}
	for i := 1; len(factors) != len(inputs); i++ {
		q := []Pulse{{"button", "roadcaster", false}}
		for len(q) > 0 {
			pulse := q[0]
			q = q[1:]
			q = append(q, pulse.Process(modules)...)
			for _, v := range inputs {
				if _, exists := factors[v]; !exists && modules[rxInputs[0]].memory[v] {
					factors[v] = i
				}
			}
		}
	}
	answer = 1
	for _, v := range factors {
		answer *= v
	}
	return
}

func Solve(input string, part int) (answer int) {
	modules := ReadConfiguration(input)
	if part == 1 {
		answer = Button(modules, 1_000)
	} else {
		answer = FirstRX(modules)
	}
	return
}
