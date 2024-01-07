package d20

import (
	"encoding/json"
	"fmt"
	"log"
	"slices"
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

const (
	BCAST int = iota
	FLIP
	CONJ
)

const (
	LOPULSE int = iota
	HIPULSE
)

type Module struct {
	Name     string
	Kind     int
	On       bool
	Remember map[string]int
	Outputs  []string
}

type Puzzle struct {
	Modules             map[string]Module
	Names, BcastOutputs []string
	Lo, Hi              int
}

func (p *Puzzle) ReadConfiguration(input string) {
	for _, entry := range strings.Split(input, "\n") {
		halves := strings.Split(entry, " -> ")
		outputs := strings.Split(halves[1], ", ")
		name := halves[0][1:]
		switch string(halves[0][0]) {
		case "b":
			p.BcastOutputs = outputs
		case `%`:
			p.Names = append(p.Names, name)
			p.Modules[name] = Module{name, FLIP, false, map[string]int{}, outputs}
		case `&`:
			p.Names = append(p.Names, name)
			p.Modules[name] = Module{name, CONJ, false, map[string]int{}, outputs}
		}
	}
	for name, mod := range p.Modules {
		for _, o := range mod.Outputs {
			if slices.Contains(p.Names, o) && p.Modules[o].Kind == CONJ {
				p.Modules[o].Remember[name] = LOPULSE
			}
		}
	}
}

type Signal struct {
	Origin, Target string
	Pulse          int
}

func (p *Puzzle) PressButton() {
	send := LOPULSE
	q := []Signal{}
	p.Lo++
	log.Printf("button -%d-> broadcaster", send)
	for _, bt := range p.BcastOutputs {
		log.Printf("broadcaster -%d-> %s", send, bt)
		q = append(q, Signal{"broadcaster", bt, send})
	}
	for len(q) > 0 {
		this := q[0]
		q = q[1:]
		if this.Pulse == LOPULSE {
			p.Lo++
		} else {
			p.Hi++
		}
		if _, exists := p.Modules[this.Target]; !exists {
			continue
		}
		module := p.Modules[this.Target]
		if module.Kind == FLIP {
			if this.Pulse == LOPULSE {
				if !module.On {
					module.On = true
				}
				if module.On {
					send = HIPULSE
				}
				for _, o := range module.Outputs {
					log.Printf("%s -%d-> %s", module.Name, send, o)
					q = append(q, Signal{module.Name, o, send})
					// log.Printf("q = %#v", q)
				}
			}
		} else {
			module.Remember[this.Origin] = this.Pulse
			if AllHi(module.Remember) {
				send = LOPULSE
			} else {
				send = HIPULSE
			}
			for _, o := range module.Outputs {
				log.Printf("%s -%d-> %s", module.Name, send, o)
				q = append(q, Signal{module.Name, o, send})
				// log.Printf("q = %#v", q)
			}
		}
		// p.Modules[this.Target] = module
	}
}

func AllHi(mem map[string]int) bool {
	log.Printf("mem = %#v", mem)
	for _, v := range mem {
		if v == LOPULSE {
			return false
		}
	}
	return true
}

func Solve(input string, part int) (answer int) {
	p := &Puzzle{map[string]Module{}, []string{}, []string{}, 0, 0}
	p.ReadConfiguration(input)
	for i := 0; i < 1_000; i++ {
		p.PressButton()
	}

	info, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		log.Fatal("Broken JSON")
	}
	fmt.Println(string(info))

	// for i := 0; i < 1; i++ {
	// 	mods.PressButton()
	// }

	log.Printf("lo=%d, hi=%d", p.Lo, p.Hi)
	answer = p.Lo * p.Hi
	return
}
