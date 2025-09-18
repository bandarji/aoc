package adventofcode

import (
	"fmt"
	"strings"
)

const y15d15Input string = `Sprinkles: capacity 5, durability -1, flavor 0, texture 0, calories 5
PeanutButter: capacity -1, durability 3, flavor 0, texture 0, calories 1
Frosting: capacity 0, durability -1, flavor 4, texture 0, calories 6
Sugar: capacity -1, durability 0, flavor 0, texture 2, calories 8`

type Y15D15 struct{}

func (d *Y15D15) GetInput(year, day int) string {
	return y15d15Input
}

func (d *Y15D15) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d15(d.GetInput(year, day), 1))
}

func (d *Y15D15) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d15(d.GetInput(year, day), 2))
}

func y15d15(input string, part int) (answer int) {
	ingredients := y15d15ReadIngredients(input)
	bestScore, bestCalScore := 0, 0
	for i1 := range 100 {
		for i2 := range 100 {
			for i3 := range 100 {
				i4 := 100 - i1 - i2 - i3
				capacity := max(0, i1*ingredients[0][0]+i2*ingredients[1][0]+i3*ingredients[2][0]+i4*ingredients[3][0])
				durability := max(0, i1*ingredients[0][1]+i2*ingredients[1][1]+i3*ingredients[2][1]+i4*ingredients[3][1])
				flavor := max(0, i1*ingredients[0][2]+i2*ingredients[1][2]+i3*ingredients[2][2]+i4*ingredients[3][2])
				texture := max(0, i1*ingredients[0][3]+i2*ingredients[1][3]+i3*ingredients[2][3]+i4*ingredients[3][3])
				calories := max(0, i1*ingredients[0][4]+i2*ingredients[1][4]+i3*ingredients[2][4]+i4*ingredients[3][4])
				score := capacity * durability * flavor * texture
				if calories == 500 {
					bestCalScore = max(bestCalScore, score)
				}
				bestScore = max(bestScore, score)
			}
		}
	}
	if part == 1 {
		answer = bestScore
	} else {
		answer = bestCalScore
	}
	return
}

func y15d15ReadIngredients(input string) [][]int {
	ingredients := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		ingredients = append(ingredients, y15d15ReadIngredientsLine(line))
	}
	return ingredients
}

func y15d15ReadIngredientsLine(line string) []int {
	var name string
	var capacity, durability, flavor, texture, calories int
	fmt.Sscanf(
		line,
		"%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
		&name, &capacity, &durability, &flavor, &texture, &calories,
	)
	return []int{capacity, durability, flavor, texture, calories}
}
