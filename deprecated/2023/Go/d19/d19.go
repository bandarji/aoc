package d19

import (
	"fmt"
	"sort"
	"strings"
)

const TEST1 string = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

func ProcessWorkflows(input string) (workflows map[string]Workflow) {
	workflows = map[string]Workflow{}
	for _, line := range strings.Split(input, "\n") {
		ruleStr := line[strings.Index(line, "{")+1 : len(line)-1]
		rules := strings.Split(ruleStr, ",")
		// log.Printf("name=%s, ruleStr=%s, rules=%s", line[0:strings.Index(line, "{")], ruleStr, rules)
		wf := Workflow{Name: line[0:strings.Index(line, "{")]}
		for _, ruleStr := range rules {
			var rule Rule
			if strings.ContainsAny(ruleStr, "<>") {
				fmt.Sscanf(ruleStr[2:], "%d:%s", &rule.Value, &rule.Next)
				rule.Part = ruleStr[0]
				rule.Comp = ruleStr[1]
			} else {
				rule.Next = ruleStr
			}
			wf.Rules = append(wf.Rules, rule)
		}
		workflows[wf.Name] = wf
	}
	return
}

func ProcessRatings(input string) (ratings []map[byte]int) {
	ratings = []map[byte]int{}
	var x, m, a, s int
	for _, line := range strings.Split(input, "\n") {
		fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
		ratings = append(ratings, map[byte]int{'x': x, 'm': m, 'a': a, 's': s})
	}
	return
}

func ProcessInput(input string) (workflows map[string]Workflow, ratings []map[byte]int) {
	sections := strings.Split(input, "\n\n")
	workflows = ProcessWorkflows(sections[0])
	ratings = ProcessRatings(sections[1])
	return
}

type Rule struct {
	Part, Comp byte
	Value      int
	Next       string
}

func (r Rule) Run(parts map[byte]int) (bool, string) {
	if r.Comp == '<' {
		return parts[r.Part] < r.Value, r.Next
	} else if r.Comp == '>' {
		return parts[r.Part] > r.Value, r.Next
	} else {
		return true, r.Next
	}
}

type Workflow struct {
	Name  string
	Rules []Rule
}

func (w Workflow) Run(parts map[byte]int) string {
	for _, rule := range w.Rules {
		if yes, next := rule.Run(parts); yes {
			return next
		}
	}
	return w.Rules[len(w.Rules)-1].Next
}

func Processes(workflows map[string]Workflow, parts map[byte]int) bool {
	next := "in"
	for {
		next = workflows[next].Run(parts)
		if next == "R" {
			return false
		}
		if next == "A" {
			break
		}
	}
	return true
}

func Part2Run(workflows map[string]Workflow) (answer int) {
	lo := map[byte][]int{}
	hi := map[byte][]int{}
	for _, wf := range workflows {
		for _, rule := range wf.Rules {
			if rule.Comp == '<' {
				lo[rule.Part] = append(lo[rule.Part], rule.Value)
				hi[rule.Part] = append(hi[rule.Part], rule.Value-1)
			} else if rule.Comp == '>' {
				lo[rule.Part] = append(lo[rule.Part], rule.Value+1)
				hi[rule.Part] = append(hi[rule.Part], rule.Value)
			}
		}
	}
	for key := range lo {
		lo[key] = append(lo[key], 1)
		sort.Slice(lo[key], func(i, j int) bool { return lo[key][i] < lo[key][j] })
	}
	for key := range hi {
		hi[key] = append(hi[key], 4000)
		sort.Slice(hi[key], func(i, j int) bool { return hi[key][i] < hi[key][j] })
	}
	for xi, x := range lo['x'] {
		for mi, m := range lo['m'] {
			for ai, a := range lo['a'] {
				for si, s := range lo['s'] {
					parts := map[byte]int{'x': x, 'm': m, 'a': a, 's': s}
					if Processes(workflows, parts) {
						answer += (hi['x'][xi] - x + 1) *
							(hi['m'][mi] - m + 1) *
							(hi['a'][ai] - a + 1) *
							(hi['s'][si] - s + 1)
					}
				}
			}
		}
	}
	return
}

func Solve(input string, part int) (answer int) {
	workflows, ratings := ProcessInput(input)
	if part == 1 {
		for _, parts := range ratings {
			if Processes(workflows, parts) {
				for _, xmas := range parts {
					answer += xmas
				}
			}
		}
	} else {
		answer = Part2Run(workflows)
	}
	return
}
