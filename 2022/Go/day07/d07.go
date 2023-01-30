package main

import (
	"aoc/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type fileSystem map[string]int

const INPUT_FILENAME string = "input_day07.txt"
const TEST_INPUT string = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func d0701(input string) int {
	totalUsed := 0
	fs := assembleTree(input)
	for _, dirUsed := range fs {
		if dirUsed <= 100000 {
			totalUsed += dirUsed
		}
	}
	return totalUsed
}

func d0702(input string) int {
	var delete int = 0
	sizes := make([]int, 0)
	fs := assembleTree(input)
	used, _ := fs["/"]
	delete = 30000000 - (70000000 - used)
	for _, size := range fs {
		if size >= delete {
			sizes = append(sizes, size)
		}
	}
	minSize := findSmallest(&sizes)
	return minSize
}

func findSmallest(sizes *[]int) int {
	min := 0
	for index, size := range *sizes {
		if index == 0 {
			min = size
			continue
		}
		if size < min {
			min = size
		}
	}
	return min
}

func assembleTree(input string) fileSystem {
	fs := make(map[string]int)
	path := make([]string, 0)
	for _, entry := range strings.Split(input, "\n") {
		if len(entry) < 2 {
			continue
		}
		fields := strings.Fields(entry)
		if fields[0] == "$" && fields[1] != "ls" {
			if fields[1] == "cd" {
				if fields[2] == ".." {
					path = path[:len(path)-1]
				} else {
					path = append(path, fields[2])
				}
			}
		} else {
			if fields[0] != "dir" {
				fSize, _ := strconv.Atoi(fields[0])
				for depth, _ := range path {
					fPath := strings.Join(path[:depth+1], "/")
					currSize, exists := fs[fPath]
					if exists {
						fs[fPath] = currSize + fSize
					} else {
						fs[fPath] = fSize
					}
				}
			}
		}
	}
	return fs
}

func main() {
	var result int = 0
	input := common.ReadFile(INPUT_FILENAME)
	result = d0701(TEST_INPUT)
	if result != 95437 {
		errMsg := fmt.Sprintf("Failed tests: day 07 01, Received %d instead of %d", result, 95437)
		log.Fatal(errMsg)
	}
	result = d0702(TEST_INPUT)
	if result != 24933642 {
		errMsg := fmt.Sprintf("Failed tests: day 07 01, Received %d instead of %d", result, 95437)
		log.Fatal(errMsg)
	}
	log.Println("Day 07 01:", d0701(input))
	log.Println("Day 07 02:", d0702(input))
}
