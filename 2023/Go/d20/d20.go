package d20

import (
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

type Module struct {
	Name     string
	Tipe     int
	Outputs  []string
	Memory   bool
	Memories map[string]string
}

type Modules struct {
	Modules      map[string]*Module
	BcastOutputs []string
	Lo, Hi       int
}

func (m *Modules) PressButton() {
	q := [][3]string{}
	m.Lo++
	for _, bo := range m.BcastOutputs {
		q = append(q, [3]string{"broadcaster", bo, "lo"})
	}
	for len(q) > 0 {
		q0 := q[0]
		q = q[1:]
		origin, target, pulse := q0[0], q0[1], q0[2]
		if pulse == "lo" {
			m.Lo++
		} else {
			m.Hi++
		}
		send := "lo"
		if mod, exists := m.Modules[target]; !exists {
			continue
		} else {
			if mod.Tipe == FLIP {
				if pulse == "lo" {
					if !mod.Memory {
						mod.Memory = true
					}
					send = "lo"
					if mod.Memory {
						send = "hi"
					}
					for _, d := range mod.Outputs {
						q = append(q, [3]string{mod.Name, d, send})
					}
				}
			} else {
				mod.Memories[origin] = pulse
				send = PulseToSend(mod.Memories)
				for _, d := range mod.Outputs {
					q = append(q, [3]string{mod.Name, d, send})
				}
			}
		}
	}
}

func PulseToSend(memories map[string]string) string {
	for _, v := range memories {
		if v != "hi" {
			return "hi"
		}
	}
	return "lo"
}

func (m *Modules) AddBroadcastOutputs(outputs []string) {
	m.BcastOutputs = outputs
}

func (m *Modules) AddModule(name string, tipe int, outputs []string) {
	m.Modules[name] = &Module{name, tipe, outputs, false, map[string]string{}}
}

func ReadConfiguration(input string) (m *Modules) {
	m = &Modules{
		Modules:      map[string]*Module{},
		BcastOutputs: []string{},
	}
	names := []string{}
	for _, entry := range strings.Split(input, "\n") {
		tipe := BCAST
		parts := strings.Split(entry, " -> ")
		outputs := strings.Split(parts[1], ", ")
		if parts[0] == "broadcaster" {
			m.AddBroadcastOutputs(outputs)
		} else {
			name := parts[0][1:]
			switch parts[0][0] {
			case '%':
				tipe = FLIP
			case '&':
				tipe = CONJ
			}
			m.AddModule(name, tipe, outputs)
			names = append(names, name)
		}
	}
	for name, mod := range m.Modules {
		for _, o := range mod.Outputs {
			if slices.Contains(names, o) && m.Modules[o].Tipe == CONJ {
				m.Modules[o].Memories[name] = o
			}
		}
	}
	return
}

func Solve(input string, part int) (answer int) {
	mods := ReadConfiguration(input)

	// info, err := json.MarshalIndent(mods, "", "    ")
	// if err != nil {
	// 	log.Fatal("Broken JSON")
	// }
	// fmt.Println(string(info))

	for i := 0; i < 1_000; i++ {
		mods.PressButton()
	}

	log.Printf("lo=%d, hi=%d", mods.Lo, mods.Hi)
	answer = mods.Lo * mods.Hi
	return
}
