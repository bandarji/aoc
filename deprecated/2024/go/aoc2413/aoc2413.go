package aoc2413

import (
	"fmt"
	"image"
	"strings"
	"time"
)

const PART2_OFFSET = 10_000_000_000_000

const AOC2413_TEST string = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

type Games []Game

type Game struct {
	A, B, P image.Point // Button A, Button B, Prize
}

func ReadGames(input string) (games Games) {
	games = Games{}
	var ax, ay, bx, by, px, py int
	for _, gameInput := range strings.Split(input, "\n\n") {
		for _, line := range strings.Split(gameInput, "\n") {
			if strings.HasPrefix(line, "Button A") {
				fmt.Sscanf(line, "Button A: X+%d, Y+%d", &ax, &ay)
			} else if strings.HasPrefix(line, "Button B") {
				fmt.Sscanf(line, "Button B: X+%d, Y+%d", &bx, &by)
			} else if strings.HasPrefix(line, "Prize") {
				fmt.Sscanf(line, "Prize: X=%d, Y=%d", &px, &py)
			}
		}
		games = append(games, Game{A: image.Pt(ax, ay), B: image.Pt(bx, by), P: image.Pt(px, py)})
	}
	return
}

func TokensSpent(game Game, part int) (tokens int, possible bool) {
	A, B, P := game.A, game.B, game.P // Need some shorter names!!!
	if part == 2 {
		P.X += PART2_OFFSET
		P.Y += PART2_OFFSET
	}
	if (((B.X*(-P.Y))-(B.Y*(-P.X)))%((A.X*B.Y)-(A.Y*B.X)) == 0) && ((((-P.X)*A.Y)-((-P.Y)*A.X))%((A.X*B.Y)-(A.Y*B.X)) == 0) {
		x := ((B.X * (-P.Y)) - (B.Y * (-P.X))) / ((A.X * B.Y) - (A.Y * B.X))
		y := (((-P.X) * A.Y) - ((-P.Y) * A.X)) / ((A.X * B.Y) - (A.Y * B.X))
		return x*3 + y, true
	}
	return 0, false
}

func Aoc241301(input string) (total int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	games := ReadGames(input)
	for _, game := range games {
		if tokens, ok := TokensSpent(game, 1); ok {
			total += tokens
		}
	}
	return
}

func Aoc241302(input string) (total int) {
	start := time.Now()
	defer func() { fmt.Println(time.Now().Sub(start)) }()
	games := ReadGames(input)
	for _, game := range games {
		if tokens, ok := TokensSpent(game, 2); ok {
			total += tokens
		}
	}
	return
}
