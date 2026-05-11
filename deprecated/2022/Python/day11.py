import json
from typing import List
import logging
from math import lcm

from lib import get_input

Monkeys = List[dict]

PART1_INPUT = """Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
"""

# debug = logging.warning
debug = lambda x: ""

class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day11.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str, rounds: int=20, divide: bool=True) -> int:
    monkeys = _read_monkeys(input_)
    monkeys_lcm = lcm(*(m["divisible"] for m in monkeys))
    print(json.dumps(monkeys, indent=4))
    for round_, _ in enumerate(range(rounds), start=1):
        for monkey in monkeys:
            debug(f"Monkey {monkey['number']}:")
            monkey["inspections"] += len(monkey["worries"])
            for worry in monkey["worries"]:
                debug(f"  Monkey inspects an item with worry level of {worry}.")
                new_worry = _calc_new_worry(worry, monkey["operation"], monkeys_lcm, divide=divide)
                debug_msg = "is not" if new_worry % monkey["divisible"] else "is"
                debug(f"    Current worry level {new_worry} {debug_msg} divisible by {monkey['divisible']}.")
                destination = monkey["false"] if new_worry % monkey["divisible"] else monkey["true"]
                monkeys[destination]["worries"].append(new_worry)
                debug(f"    Item with worry level {new_worry} is thrown to monkey {destination} {monkeys[destination]['worries']}.")
                monkey["worries"] = []
            monkeys[monkey["number"]] = monkey
        if not divide:
            if round_ in {1, 20, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000}:
                print(f"\n== After round {round_} ==")
                for i, times in [(i, m["inspections"]) for i, m in enumerate(monkeys)]:
                    print(f"Monkey {i} inspected items {times} times.")
    monkeys.sort(key=lambda m: m["inspections"], reverse=True)
    return monkeys[0]["inspections"] * monkeys[1]["inspections"]


def _solve2(input_: str, rounds: int=10000, divide: bool=False) -> int:
    return _solve1(input_, rounds=10000, divide=False)


def _calc_new_worry(worry: int, operation: str, monkeys_lcm: int, divide: bool=True):
    ops = operation.split()
    left = worry if ops[0] == "old" else int(ops[0])
    right = worry if ops[2] == "old" else int(ops[2])
    value = left + right if ops[1] == "+" else left * right
    response = value // 3 if divide else value % monkeys_lcm
    debug(f"    Calculation: {left} {ops[1]} {right} = {value}.")
    debug(f"    Monkey gets bored with item. Worry level is divided by 3 to {response}.")
    return response


def _read_monkeys(input_: str) -> Monkeys:
    monkeys = []
    for entry in input_.split("\n"):
        if not entry:
            continue
        words = entry.split()
        if words[0] == "Monkey":
            monkey = {}
            monkey["inspections"] = 0
            monkey["number"] = int(words[1].replace(":", ""))
        elif words[0] == "Starting":
            monkey["worries"] = [int(w.replace(",", "")) for w in words[2:]]
        elif words[0] == "Operation:":
            monkey["operation"] = " ".join(words[3:])
        elif words[0] == "Test:":
            monkey["divisible"] = int(words[-1])
        elif words[0] == "If":
            if words[1] == "true:":
                monkey["true"] = int(words[-1])
            if words[1] == "false:":
                monkey["false"] = int(words[-1])
                monkeys.append(monkey)
        else:
            continue
    return monkeys



def main():
    assert _solve1(PART1_INPUT) == 10605
    print("Passed Test 1")
    assert _solve2(PART1_INPUT, rounds=10000, divide=False) == 2713310158
    print("Passed tests")
    # for test, expected in PART1_TESTS:
    #     assert expected == _solve1(test)
    # for test, expected in PART2_TESTS:
    #     assert expected == _solve2(test)
    Day()


if __name__ == '__main__':
    main()
