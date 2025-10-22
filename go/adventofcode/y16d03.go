package adventofcode

import "strings"

func y16d03(input string, part int) (answer int) {
	triangles := y16d03ParseInput(input, part)
	for _, triangle := range triangles {
		if triangle[0]+triangle[1] > triangle[2] && triangle[0]+triangle[2] > triangle[1] && triangle[1]+triangle[2] > triangle[0] {
			answer++
		}
	}
	return
}

func y16d03ParseInput(input string, part int) (triangles [][3]int) {
	for line := range strings.SplitSeq(input, "\n") {
		fields := strings.Fields(line)
		triangles = append(triangles, [3]int{strToInt(fields[0]), strToInt(fields[1]), strToInt(fields[2])})
	}
	if part == 2 {
		triangles = y16d03GetColumnedSides(triangles)
	}
	return
}

func y16d03GetColumnedSides(triangles [][3]int) (columnedTriangles [][3]int) {
	for i := 0; i < len(triangles); i += 3 {
		for c := 0; c < 3; c++ {
			columnedTriangles = append(columnedTriangles, [3]int{triangles[i][c], triangles[i+1][c], triangles[i+2][c]})
		}
	}
	return
}
