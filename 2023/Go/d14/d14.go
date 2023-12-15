package d14

import (
	"strings"
)

const TEST1 string = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

const CYCLES int = 1_000_000_000 // 1x10^9 / one billion

var Cache = make(map[string]int)

func ReadGrid(input string) (platform [][]string) {
	platform = [][]string{}
	for _, row := range strings.Split(input, "\n") {
		platform = append(platform, strings.Split(row, ""))
	}
	return
}

func Dimensions(platform [][]string) (height, width int) {
	height, width = len(platform), len(platform[0])
	return
}

func TiltPlatform(platform [][]string) ([][]string, int) {
	answer := 0
	height, width := Dimensions(platform)
	for col := 0; col < width; col++ {
		block, going := -1, 0
		for row := 0; row < height; row++ {
			if platform[row][col] == "#" {
				answer += going * (height - block - 1)
				answer -= going * (going - 1) / 2
				block, going = row, 0
			}
			if platform[row][col] == "O" {
				going++
			}
		}
		if going > 0 {
			answer += going * (height - block - 1)
			answer -= going * (going - 1) / 2
		}
	}
	return platform, answer
}

// func Rotate(platform [][]string) [][]string {
// 	height, width := Dimensions(platform)
//     rotated := make([][]string, height)
//     for i := 0; i < width; i++ {
//         rotated[i] = make([]string, height)
//     }
//     for i := 0; i < height; i++ {
//         for j := 0; j < width; j++ {
//             rotated[j][height-1-i] = platform[i][j]
//         }
//     }
//     return rotated
// }

// func TiltAllDirections(platform [][]string) (answer int) {
// 	height, width := Dimensions(platform)
// 	for i := 0; i < CYCLES; i++ {
// 		platform := PerformOneCycle(platform)
// 	}
// 	return answer
// }

// func HashPlatform(platform [][]string) string {
// 	hash := md5.Sum([]byte(fmt.Sprintf("%+v", platform)))
// 	return hex.EncodeToString(hash[:])
// }

// func CacheSet(platform [][]string, answer int) {
// 	Cache[HashPlatform(platform)] = answer
// }

// func CacheRead(platform [][]string) int {
// 	key := HashPlatform(platform)
// 	if answer, exists := Cache[key]; exists {
// 		return answer
// 	}
// 	return -1
// }

// func GetMD5Hash(text string) string {
// 	hash := md5.Sum([]byte(text))
// 	return hex.EncodeToString(hash[:])
//  }

// func PerformOneCycle(platform [][]string) (answer int) {
// 	for i := 0; i < 4; i++ {
// 		platform, answer := TiltPlatform(platform)
// 	}
// 	if answer, exists :=
// 	return
// }

func Solve(input string, part int) int {
	platform := ReadGrid(input)
	platform, answer := TiltPlatform(platform)
	_ = platform
	return answer
}
