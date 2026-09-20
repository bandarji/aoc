package adventofcode

import (
	"fmt"
	"sort"
	"strings"
)

type y17d20XYZ struct {
	xyz [3]int
}

func (n y17d20XYZ) manhattan() int {
	return absInt(n.xyz[0]) + absInt(n.xyz[1]) + absInt(n.xyz[2])
}

type y17d20Particle struct {
	i int
	p y17d20XYZ
	v y17d20XYZ
	a y17d20XYZ
}

func y17d20(input string, part int) (answer int) {
	particles := y17d20ParseInput(input)
	if part == 1 {
		answer = y17d20FindClosestParticle(particles)
	} else {
		answer = y17d20CountRemainingParticles(particles)
	}
	return
}

func y17d20CountRemainingParticles(particles []y17d20Particle) int {
	for k := 0; k < 1<<8; k++ {
		// tick
		newParticles := []y17d20Particle{}
		for _, p := range particles {
			for i, a := range p.a.xyz {
				p.v.xyz[i] += a
			}
			for j, v := range p.v.xyz {
				p.p.xyz[j] += v
			}
			newParticles = append(newParticles, p)
		}
		particles = newParticles
		// remove collisions
		set := map[[3]int]int{}
		for _, p := range particles {
			set[p.p.xyz]++
		}
		np := []y17d20Particle{}
		for _, p := range particles {
			if c, ok := set[p.p.xyz]; ok && c == 1 {
				np = append(np, p)
			}
		}
		particles = np
	}
	return len(particles) // TODO: remove collisions
}

func y17d20FindClosestParticle(particles []y17d20Particle) (closest int) {
	sort.Slice(particles, func(i, j int) bool {
		pi, pj := particles[i], particles[j]
		if pi.a.manhattan() != pj.a.manhattan() {
			return pi.a.manhattan() < pj.a.manhattan()
		}
		if pi.v.manhattan() != pj.v.manhattan() {
			return pi.v.manhattan() > pj.v.manhattan()
		}
		return pi.p.manhattan() < pj.p.manhattan()
	})
	closest = particles[0].i
	return
}

func y17d20ParseInput(input string) (particles []y17d20Particle) {
	for i, line := range strings.Split(input, "\n") {
		p := y17d20Particle{i: i}
		fmt.Sscanf(
			line,
			"p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&p.p.xyz[0], &p.p.xyz[1], &p.p.xyz[2],
			&p.v.xyz[0], &p.v.xyz[1], &p.v.xyz[2],
			&p.a.xyz[0], &p.a.xyz[1], &p.a.xyz[2],
		)
		particles = append(particles, p)
	}
	return
}
