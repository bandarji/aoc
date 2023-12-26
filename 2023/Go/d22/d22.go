package d22

import (
	"aoc2023/common"
	"sort"
	"strings"
)

const TEST1 string = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

type Brick []int
type ByBricks []Brick

func (b ByBricks) Len() int           { return len(b) }
func (b ByBricks) Less(i, j int) bool { return b[i][2] < b[j][2] }
func (b ByBricks) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func ReadBricks(input string) (bricks []Brick) {
	bricks = []Brick{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.Replace(line, "~", ",", -1)
		values := Brick{}
		for _, value := range strings.Split(line, ",") {
			values = append(values, common.StrToInt(value))
		}
		bricks = append(bricks, values)
	}
	return
}

func Coincides(b1, b2 Brick) bool {
	if max(b1[0], b2[0]) <= min(b1[3], b2[3]) && max(b1[1], b2[1]) <= min(b1[4], b2[4]) {
		return true
	}
	return false
}

func DropBricks(bricks []Brick, part int) (answer int) {
	kSupV := map[int][]int{}
	vSupK := map[int][]int{}
	sort.Sort(ByBricks(bricks))
	maxZ := 1
	for i, brick := range bricks {
		maxZ = 1
		for _, check := range bricks[0:i] {
			if Coincides(brick, check) {
				maxZ = max(maxZ, check[5]+1)
			}
		}
		brick[5] -= brick[2] - maxZ
		brick[2] = maxZ
	}
	sort.Sort(ByBricks(bricks))
	for i, upper := range bricks {
		for j, lower := range bricks[0:i] {
			if Coincides(lower, upper) && upper[2] == lower[5]+1 {
				if _, exists := kSupV[i]; !exists {
					kSupV[i] = []int{}
				}
				kSupV[i] = append(kSupV[i], j)
				if _, exists := vSupK[j]; !exists {
					vSupK[j] = []int{}
				}
				vSupK[j] = append(vSupK[j], i)
			}
		}
	}
	if part == 1 {
		for i := 0; i < len(bricks); i++ {
			dis := true
			for _, j := range vSupK[i] {
				if len(kSupV[j]) < 2 {
					dis = false
				}
			}
			if dis {
				answer += 1
			}
		}
	} else {
		for i := 0; i < len(bricks); i++ {
			q := []int{}
			falling := map[int]bool{}
			falling[i] = true
			for _, j := range vSupK[i] {
				if len(kSupV[j]) == 1 {
					q = append(q, j)
				}
			}
			for _, v := range q {
				falling[v] = true
			}
			for len(q) > 0 {
				j := q[0]
				q = q[1:]
				for _, k := range vSupK[j] {
					if falling[k] {
						continue
					}
					all := true
					for _, n := range kSupV[k] {
						if !falling[n] {
							all = false
							break
						}
					}
					if all {
						q = append(q, k)
						falling[k] = true
					}
				}
			}
			answer += len(falling) - 1
		}
	}
	return
}

func Solve(input string, part int) (answer int) {
	bricks := ReadBricks(input)
	answer = DropBricks(bricks, part)
	return
}
