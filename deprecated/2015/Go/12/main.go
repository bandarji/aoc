package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func JSONValues(input string) (total int) {
	num := 0
	s := strings.Builder{}
	for _, char := range input {
		if char == '\n' {
			continue
		}
		if char == '-' || (char >= '0' && char <= '9') {
			s.WriteRune(char)
		} else {
			if s.Len() > 0 {
				num, _ = strconv.Atoi(s.String())
				total += num
				s.Reset()
			}
		}
	}
	if s.Len() > 0 {
		num, _ := strconv.Atoi(s.String())
		total += num
	}
	return
}

func JSON2(input string) int {
	if !regexp.MustCompile("[{}]").MatchString(input) ||
		!regexp.MustCompile("red").MatchString(input) {
		return JSONValues(input)
	}

	var obj map[string]interface{}
	err := json.Unmarshal([]byte(input), &obj)
	if err != nil {
		var arr []interface{}
		err := json.Unmarshal([]byte(input), &arr)
		if err != nil {
			panic(err)
		}

		var arrayTotal int
		for _, v := range arr {
			// marshal each array element into a string, then pass it back into part2
			str, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			arrayTotal += JSON2(string(str))
		}

		return arrayTotal
	}
	for _, v := range obj {
		str, ok := v.(string)
		if ok && str == "red" {
			return 0
		}
	}

	var total int
	for _, v := range obj {
		str, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		total += JSON2(string(str))
	}

	return total
}

func RunTests(part int) bool {
	tests := map[string]int{}
	if part == 1 {
		tests = map[string]int{
			"[1,2,3]":              6,
			`{"a":2,"b":4}`:        6,
			"[[[3]]]":              3,
			`{"a":{"b":4},"c":-1}`: 3,
			`{"a":[-1,1]}`:         0,
			`[-1,{"a":1}]`:         0,
			"[]":                   0,
			"{}":                   0,
		}
	} else {
		tests = map[string]int{
			"[1,2,3]":                         6,
			`[1,{"c":"red","b":2},3]`:         4,
			`{"d":"red","e":[1,2,3,4],"f":5}`: 0,
			`[1,"red",5]`:                     6,
		}
	}
	for test, expected := range tests {
		result := 0
		if part == 1 {
			result = JSONValues(test)
		} else {
			result = JSON2(test)
		}
		if result != expected {
			return false
		}
	}
	return true
}

func main() {
	day, part := 12, 1
	if !RunTests(part) {
		log.Fatal("Tests failed, day one")
	}
	content, _ := os.ReadFile("input_day12.txt")
	fmt.Println("day", day, "part", part, "ans =", JSONValues(string(content)))
	part++
	if !RunTests(part) {
		log.Fatal("Tests failed, day two")
	}
	fmt.Println("day", day, "part", part, "ans =", JSON2(string(content)))
}
