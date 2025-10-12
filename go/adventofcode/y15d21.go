package adventofcode

import (
	"fmt"
	"strings"
)

const y15d21StartingHP = 100
const y15d21Shop string = `Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3`

type y15d21Item struct {
	name                string
	cost, damage, armor int
}

type y15d21Fighter struct {
	hp, damage, armor, cost int
}

type y15d21ShopInventory struct {
	weapons, armor, rings []y15d21Item
}

func y15d21(input string, startingHP, part int) (answer int) {
	shop := y15d21ReadShop()
	boss := y15d21ReadBoss(input)
	equips := y15d21EquipConfigs(shop)
	mostGold, leastGold := y15d21DoBattle(boss, startingHP, equips)
	switch part {
	case 1:
		answer = leastGold
	case 2:
		answer = mostGold
	}
	return
}

func y15d21DoBattle(boss y15d21Fighter, startingHP int, equips [][]y15d21Item) (mostGold, leastGold int) {
	mostGold, leastGold = 0, 2_000_000_000
	for _, equip := range equips {
		me := y15d21Fighter{hp: startingHP, damage: 0, armor: 0, cost: 0}
		for _, item := range equip {
			me.damage += item.damage
			me.armor += item.armor
			me.cost += item.cost
		}
		if y15d21WinOrLose(boss, me) {
			leastGold = min(leastGold, me.cost)
		} else {
			mostGold = max(mostGold, me.cost)
		}
	}
	return
}

func y15d21WinOrLose(boss, me y15d21Fighter) bool {
	attackBoss := max(1, me.damage-boss.armor)
	attackMe := max(1, boss.damage-me.armor)
	for boss.hp > 0 && me.hp > 0 {
		boss.hp -= attackBoss
		me.hp -= attackMe
	}
	return boss.hp <= 0
}

func y15d21EquipConfigs(shop y15d21ShopInventory) (equips [][]y15d21Item) {
	for weapon := 0; weapon < len(shop.weapons); weapon++ {
		for armor := -1; armor < len(shop.armor); armor++ {
			for ring1 := -1; ring1 < len(shop.rings); ring1++ {
				for ring2 := -1; ring2 < len(shop.rings); ring2++ {
					equip := []y15d21Item{shop.weapons[weapon]}
					if armor > -1 {
						equip = append(equip, shop.armor[armor])
					}
					if ring1 > -1 {
						equip = append(equip, shop.rings[ring1])
					}
					if ring2 > -1 && ring2 != ring1 {
						equip = append(equip, shop.rings[ring2])
					}
					equips = append(equips, equip)
				}
			}
		}
	}
	return
}

func y15d21ReadBoss(input string) y15d21Fighter {
	hp, damage, armor := 0, 0, 0
	fmt.Sscanf(input, "Hit Points: %d\nDamage: %d\nArmor: %d", &hp, &damage, &armor)
	return y15d21Fighter{hp: hp, damage: damage, armor: armor}
}

func y15d21ReadShop() (shop y15d21ShopInventory) {
	sections := strings.Split(y15d21Shop, "\n\n")
	shop.weapons = y15d21ReadShopSection(sections[0])
	shop.armor = y15d21ReadShopSection(sections[1])
	shop.rings = y15d21ReadShopSection(strings.ReplaceAll(sections[2], " +", "+"))
	return
}

func y15d21ReadShopSection(section string) (items []y15d21Item) {
	for i, line := range strings.Split(section, "\n") {
		if i == 0 {
			continue
		}
		var name string
		var cost, damage, armor int
		fmt.Sscanf(line, "%s %d %d %d", &name, &cost, &damage, &armor)
		items = append(items, y15d21Item{name: name, cost: cost, damage: damage, armor: armor})
	}
	return
}
