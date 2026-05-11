package d11

import (
	"log"
	"strings"
)

const TEST1 string = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func ReadImage(input string) (image [][]string) {
	image = [][]string{}
	for _, line := range strings.Split(input, "\n") {
		image = append(image, strings.Split(line, ""))
	}
	return
}

func Display(image [][]string) {
	for _, row := range image {
		log.Println(strings.Join(row, ""))
	}
}

func EmptySpace(strip []string) bool {
	for _, char := range strip {
		if char != "." {
			return false
		}
	}
	return true
}

func EmptyRows(image [][]string) (rows []int) {
	rows = []int{}
	for r, row := range image {
		if EmptySpace(row) {
			rows = append(rows, r)
		}
	}
	return
}

func EmptyCols(image [][]string) (cols []int) {
	cols = []int{}
	height := len(image)
	width := len(image[0])
	for c := 0; c < width; c++ {
		strip := []string{}
		for r := 0; r < height; r++ {
			strip = append(strip, image[r][c])
		}
		if EmptySpace(strip) {
			cols = append(cols, c)
		}
	}
	return
}

func PinpointGalaxies(image [][]string) (galaxies [][2]int) {
	galaxies = [][2]int{}
	for r, row := range image {
		for c, char := range row {
			if char == "#" {
				galaxies = append(galaxies, [2]int{r, c})
			}
		}
	}
	return
}

func DistanceToAllGalaxies(expansion int, galaxies [][2]int, eRows, eCols []int) (total int) {
	emptyRows := map[int]bool{}
	emptyCols := map[int]bool{}
	for _, r := range eRows {
		emptyRows[r] = true
	}
	for _, c := range eCols {
		emptyCols[c] = true
	}
	for i, galaxy := range galaxies {
		for _, galaxy2 := range galaxies[i:] {
			r1, c1, r2, c2 := galaxy[0], galaxy[1], galaxy2[0], galaxy2[1]
			for r := min(r1, r2); r < max(r1, r2); r++ {
				total++
				if emptyRows[r] {
					total += expansion - 1
				}
			}
			for c := min(c1, c2); c < max(c1, c2); c++ {
				total++
				if emptyCols[c] {
					total += expansion - 1
				}
			}
		}
	}
	return
}

func Solve(input string, part int, expansion int) (answer int) {
	image := ReadImage(input)
	emptyRows := EmptyRows(image)
	emptyCols := EmptyCols(image)
	galaxies := PinpointGalaxies(image)
	answer = DistanceToAllGalaxies(expansion, galaxies, emptyRows, emptyCols)

	return
}
