package adventofcode

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// const y15d07InterestingWire string = "a"

type Y15D12 struct{}

func (d *Y15D12) GetInput(year, day int) string {
	return readContent(formatFilename(year, day))
}

func (d *Y15D12) Part1(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 1: %d", year, day, y15d12(d.GetInput(year, day), 1))
}

func (d *Y15D12) Part2(year, day int) string {
	return fmt.Sprintf("Year=%d Day=%02d Part 2: %d", year, day, y15d12(d.GetInput(year, day), 2))
}

func y15d12(input string, part int) int {
	if part == 1 {
		return y15d12p1(input)
	}
	return y15d12p2(input)
}

func y15d12p1(input string) (total int) {
	sb := strings.Builder{}
	for _, char := range input {
		if regexp.MustCompile("[-0-9]").MatchString(string(char)) {
			sb.WriteRune(char)
		} else if sb.Len() > 0 {
			total += strToInt(sb.String())
			sb.Reset()
		}
	}
	if sb.Len() > 0 {
		total += strToInt(sb.String())
	}
	return
}

func y15d12p2(input string) int {
	listTotal, grandTotal := 0, 0
	if !regexp.MustCompile("[{}]").MatchString(input) || !regexp.MustCompile("red").MatchString(input) {
		return y15d12p1(input)
	}
	var j map[string]any
	err := json.Unmarshal([]byte(input), &j)
	if err != nil {
		list := []any{}
		json.Unmarshal([]byte(input), &list)
		for _, value := range list {
			s, _ := json.Marshal(value)
			listTotal += y15d12p2(string(s))
		}
		return listTotal
	}
	for _, value := range j {
		if s, ok := value.(string); ok && s == "red" {
			return 0
		}
	}
	for _, value := range j {
		s, _ := json.Marshal(value)
		grandTotal += y15d12p2(string(s))
	}
	return grandTotal
}
