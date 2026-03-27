package adventofcode

import (
	"fmt"
	"strings"
)

type y16d22FS struct {
	x, y, size, used, avail, percent int
}

func y16d22(input string, part int) int {
	filesystems := y16d22ParseInput(input)
	if part == 1 {
		return y16d22PartOne(filesystems)
	}
	return y16d22PartTwo(filesystems)
}

func y16d22PartTwo(filesystems []y16d22FS) int {
	return 0
}

func y16d22PartOne(filesystems []y16d22FS) int {
	viable := 0
	for _, fs1 := range filesystems {
		for _, fs2 := range filesystems {
			if y16d22SameFileSystems(fs1, fs2) {
				continue
			}
			if fs1.used > 0 && fs1.used <= fs2.avail {
				viable++
			}
		}
	}
	return viable
}

func y16d22ParseInput(input string) []y16d22FS {
	filesystems := []y16d22FS{}
	for _, line := range strings.Split(input, "\n") {
		if !strings.HasPrefix(line, "/dev/grid/node-x") {
			continue
		}
		var fs y16d22FS
		line = strings.Join(strings.Fields(line), " ")
		fmt.Sscanf(
			line,
			"/dev/grid/node-x%d-y%d %dT %dT %dT %d%%",
			&fs.x, &fs.y, &fs.size, &fs.used, &fs.avail, &fs.percent,
		)
		filesystems = append(filesystems, fs)
	}
	return filesystems
}

func y16d22SameFileSystems(fs1, fs2 y16d22FS) bool {
	return fs1.x == fs2.x && fs1.y == fs2.y
}
