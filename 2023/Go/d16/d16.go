package d16

import (
	"fmt"
	"strings"
)

const TEST1 string = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

type Tile [4]int

type Game struct {
	Board   [][]rune
	Beamed  map[string]bool
	Visited map[string]bool
	Vector  Tile
	Q       []Tile
}

func (g *Game) MapContraption(input string) {
	g.Board = [][]rune{}
	for _, row := range strings.Split(input, "\n") {
		g.Board = append(g.Board, []rune(row))
	}
}

func (g Game) AlreadyBeamed(t Tile) bool {
	key := LongKey(t)
	return g.Beamed[key]
}

func (g Game) EnergizedCount() int {
	return len(g.Visited)
}

func (g Game) ShowEnergized() {
	for y := 0; y < len(g.Board); y++ {
		for x := 0; x < len(g.Board[0]); x++ {
			t := Tile{y, x, 0, 0}
			if g.Visited[ShortKey(t)] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func (g *Game) ShootBeam(start Tile) {
	var y, x, dy, dx int
	var char rune
	g.Beamed = map[string]bool{}
	g.Visited = map[string]bool{}
	g.Q = []Tile{start}
	for len(g.Q) > 0 {
		y, x, dy, dx = g.Q[0][0], g.Q[0][1], g.Q[0][2], g.Q[0][3]
		g.Q = g.Q[1:]
		y += dy
		x += dx
		if y < 0 || x < 0 || y >= len(g.Board) || x >= len(g.Board[0]) {
			continue
		}
		char = g.Board[y][x]
		if char == '.' || (char == '-' && dx != 0) || (char == '|' && dy != 0) {
			t := Tile{y, x, dy, dx}
			if !g.AlreadyBeamed(t) {
				g.Beamed[LongKey(t)] = true
				g.Visited[ShortKey(t)] = true
				g.Q = append(g.Q, t)
			}
		} else if char == '/' {
			t := Tile{y, x, dx * -1, dy * -1}
			if !g.AlreadyBeamed(t) {
				g.Beamed[LongKey(t)] = true
				g.Visited[ShortKey(t)] = true
				g.Q = append(g.Q, t)
			}
		} else if char == '\\' {
			t := Tile{y, x, dx, dy}
			if !g.AlreadyBeamed(t) {
				g.Beamed[LongKey(t)] = true
				g.Visited[ShortKey(t)] = true
				g.Q = append(g.Q, t)
			}
		} else {
			if char == '|' {
				for _, dy := range [2]int{-1, 1} {
					t := Tile{y, x, dy, 0}
					if !g.AlreadyBeamed(t) {
						g.Beamed[LongKey(t)] = true
						g.Visited[ShortKey(t)] = true
						g.Q = append(g.Q, t)
					}
				}
			} else {
				for _, dx := range [2]int{-1, 1} {
					t := Tile{y, x, 0, dx}
					if !g.AlreadyBeamed(t) {
						g.Beamed[LongKey(t)] = true
						g.Visited[ShortKey(t)] = true
						g.Q = append(g.Q, t)
					}
				}
			}
		}
	}
}

func LongKey(v Tile) string {
	return fmt.Sprintf("y=%d,x=%d,dy=%d,dx=%d", v[0], v[1], v[2], v[3])
}

func ShortKey(v Tile) string {
	return fmt.Sprintf("y=%d,x=%d", v[0], v[1])
}

func Solve(input string, part int) (answer int) {
	game := &Game{}
	game.MapContraption(input)
	if part == 1 {
		game.ShootBeam(Tile{0, -1, 0, 1})
		answer = game.EnergizedCount()
	} else {
		most := 0
		board := game.Board
		height, width := len(board), len(board[0])
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				for _, dy := range [2]int{-1, 1} {
					game = &Game{Board: board}
					game.ShootBeam(Tile{y, x, dy, 0})
					if game.EnergizedCount() > most {
						most = game.EnergizedCount()
					}
				}
				for _, dx := range [2]int{-1, 1} {
					game = &Game{Board: board}
					game.ShootBeam(Tile{y, x, 0, dx})
					if game.EnergizedCount() > most {
						most = game.EnergizedCount()
					}
				}
			}
		}
		answer = most
	}
	return
}
