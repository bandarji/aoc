package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Direction 	string
	Units 		int
}

type Location struct {
	Position, Depth, Aim	int
}


func assembleCommand(entry string) Command {
	words := strings.Fields(entry)
	units, _ := strconv.Atoi(words[1])
	command := Command{words[0], units}
	return command
}

func moveSubOne(s *Location, c *Command) {
	switch c.Direction {
	case "forward":
		s.Position += c.Units
	case "down":
		s.Depth += c.Units
	case "up":
		s.Depth -= c.Units
	}
}


func moveSubTwo(s *Location, c *Command) {
	switch c.Direction {
	case "down": s.Aim += c.Units
	case "up": s.Aim -= c.Units
	case "forward":
		{
			s.Position += c.Units
			s.Depth += c.Units * s.Aim
		}
	}
}


func processCommands(filename string) (int, int) {
	product1 := 0
	product2 := 0
	sub1 := Location{0, 0, 0}
	sub2 := Location{0, 0, 0}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content := bufio.NewScanner(f)
	for content.Scan() {
		command := assembleCommand(content.Text())
		moveSubOne(&sub1, &command)
		moveSubTwo(&sub2, &command)
	}
	fmt.Println(sub1, sub2)
	product1 = sub1.Position * sub1.Depth
	product2 = sub2.Position * sub2.Depth
	return product1, product2
}


func main() {
	answer1, answer2 := processCommands("input.txt")
	fmt.Println("answer 1:", answer1)
	fmt.Println("answer 2:", answer2)
}
