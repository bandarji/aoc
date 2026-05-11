package d10

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day(input string, part int) (bot int) {
	bots, instructions := ParseInstructions(input)
	bot = ProcessInstructions(bots, instructions, part)
	return
}

func ProcessInstructions(bots map[string][]string, instructions []Instruction, part int) int {
	outputs := map[string]string{}
	for {
		for _, ins := range instructions {
			if len(bots[ins.Bot]) > 2 {
				os.Exit(1)
			}
			if len(bots[ins.Bot]) == 2 {
				sort.Strings(bots[ins.Bot])
				low, high := bots[ins.Bot][0], bots[ins.Bot][1]
				bots[ins.Bot] = []string{}
				if part == 1 && low == "00017" && high == "00061" {
					botNumber, _ := strconv.Atoi(strings.TrimLeft(ins.Bot, "0"))
					return botNumber
				}
				var outBucket, receiveBot string
				if strings.HasPrefix(ins.RuleLo, "output ") {
					fmt.Sscanf(ins.RuleLo, "output %s", &outBucket)
					outputs[PadNumber(outBucket)] = low
				} else {
					fmt.Sscanf(ins.RuleLo, "bot %s", &receiveBot)
					bots[PadNumber(receiveBot)] = append(bots[PadNumber(receiveBot)], low)
				}
				if strings.HasPrefix(ins.RuleHi, "output ") {
					fmt.Sscanf(ins.RuleHi, "output %s", &outBucket)
					outputs[PadNumber(outBucket)] = high
				} else {
					fmt.Sscanf(ins.RuleHi, "bot %s", &receiveBot)
					bots[PadNumber(receiveBot)] = append(bots[PadNumber(receiveBot)], high)
				}
			}
		}
		if part == 2 && len(outputs["00000"]) > 1 && len(outputs["00001"]) > 1 && len(outputs["00002"]) > 1 {
			break
		}
	}
	var n0, n1, n2 int
	n0, _ = strconv.Atoi(strings.TrimLeft(outputs["00000"], "0"))
	n1, _ = strconv.Atoi(strings.TrimLeft(outputs["00001"], "0"))
	n2, _ = strconv.Atoi(strings.TrimLeft(outputs["00002"], "0"))
	return n0 * n1 * n2
}

type Instruction struct {
	Bot, RuleLo, RuleHi string
}

func ParseInstructions(input string) (bots map[string][]string, instructions []Instruction) {
	bots = map[string][]string{}
	instructions = []Instruction{}
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)
		_ = fields
		if strings.HasPrefix(line, "value ") {
			var chip, botID string
			chip = PadNumber(fields[1])
			botID = PadNumber(fields[5])
			bots[botID] = append(bots[botID], chip)
		} else if strings.HasPrefix(line, "bot ") {
			var bot, lo, hi string
			bot = PadNumber(fields[1])
			lo = strings.Join(fields[5:7], " ")
			hi = strings.Join(fields[len(fields)-2:], " ")
			instructions = append(instructions, Instruction{bot, lo, hi})
		}
	}
	return
}

func PadNumber(s string) (out string) {
	out = fmt.Sprintf("%05s", s)
	return
}
