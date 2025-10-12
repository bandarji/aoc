package adventofcode

import (
	"fmt"
)

type y15d22Spell struct {
	index, cost, turns, dmg, eDmg, iHeal, heal, armor, recharge int
}

type y15d22Status struct {
	hp, mana, bossHP, bossDmg int
	effects                   [5]int
	myTurn                    bool
}

var y15d22Spells = map[string]y15d22Spell{
	"Magic Missile": {index: 0, cost: 53, dmg: 4},
	"Drain":         {index: 1, cost: 73, dmg: 2, heal: 2},
	"Shield":        {index: 2, cost: 113, turns: 6, armor: 7},
	"Poison":        {index: 3, cost: 173, turns: 6, eDmg: 3},
	"Recharge":      {index: 4, cost: 229, turns: 5, recharge: 101},
}

type Y15D22 struct{}

func (d *Y15D22) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D22) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d22(d.GetInput(year, day), 1, 50, 500))
}

func (d *Y15D22) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d22(d.GetInput(year, day), 2, 50, 500))
}

func y15d22(input string, part, hp, mana int) (answer int) {
	bossHP, bossDmg := y15d22ReadBoss(input)
	status := y15d22Status{hp: hp, mana: mana, bossHP: bossHP, bossDmg: bossDmg, effects: [5]int{}, myTurn: true}
	answer = y15d22Battle(status, part)
	return
}

func y15d22ReadBoss(input string) (hp, dmg int) {
	fmt.Sscanf(input, "Hit Points: %d\nDamage: %d", &hp, &dmg)
	return
}

func y15d22Battle(status y15d22Status, part int) (answer int) {
	answer = 2_000_000_000
	if part == 2 && status.myTurn {
		status.hp--
	}
	if status.hp < 1 {
		return 2_000_000_000
	}
	var armor int
	for _, spell := range y15d22Spells {
		if status.effects[spell.index] > 0 {
			status.effects[spell.index]--
			status.bossHP -= spell.eDmg
			status.hp += spell.heal
			armor += spell.armor
			status.mana += spell.recharge
		}
	}
	if status.bossHP < 1 {
		return 0
	}
	if status.myTurn {
		var castSpell bool
		for _, spell := range y15d22Spells {
			if status.effects[spell.index] == 0 {
				if status.mana >= spell.cost {
					castSpell = true
					updatedTurns := status.effects
					updatedTurns[spell.index] += spell.turns
					updatedStatus := y15d22Status{
						hp:      status.hp + spell.iHeal,
						mana:    status.mana - spell.cost,
						bossHP:  status.bossHP - spell.dmg,
						bossDmg: status.bossDmg,
						effects: updatedTurns,
						myTurn:  false,
					}
					result := spell.cost + y15d22Battle(updatedStatus, part)
					answer = min(answer, result)
				}
			}
		}
		if !castSpell {
			return 2_000_000_000
		}
	} else {
		bossAttack := status.bossDmg - armor
		if bossAttack < 1 {
			bossAttack = 1
		}
		updatedStatus := y15d22Status{
			hp:      status.hp - bossAttack,
			mana:    status.mana,
			bossHP:  status.bossHP,
			bossDmg: status.bossDmg,
			effects: status.effects,
			myTurn:  true,
		}
		result := y15d22Battle(updatedStatus, part)
		answer = min(answer, result)
	}
	return
}
