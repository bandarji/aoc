import json
from typing import Dict, List

from lib import get_input


Monkeys = Dict[str, int]
Strings = List[str]


TEST = """root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32"""


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day21.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    entries = _parse_input(input_)
    monkeys = _process_monkeys(entries)
    return monkeys.get("root", 0)


def _solve2(input_: str) -> int:
    entries = _parse_input(input_)
    monkeys = _process_monkeys_2(entries)
    return monkeys.get("root", 0)

def _process_monkeys_2(entries: Strings) -> Monkeys:
    known_monkeys = {}
    unknown_monkeys = {}
    for entry in entries:
        words = entry.split()
        name = entry[:4]
        if len(words) == 2 and name != "humn":
            known_monkeys[name] = int(words[1])
        else:
            unknown_monkeys[name] = words[1:]
    cycle = 0
    while len(unknown_monkeys) > 1:
        cycle += 1
        print(f"Cycle: {cycle}\nUnknown Monkeys: {len(unknown_monkeys)}\n")
        for monkey, right_side in unknown_monkeys.items():
            if monkey == "humn" or "humn" in right_side:
                continue
            left, operand, right = right_side
            if left in known_monkeys:
                left = known_monkeys.get(left)
            if right in known_monkeys:
                right = known_monkeys.get(right)
            if isinstance(left, int) and isinstance(right, int):
                known_monkeys[monkey] = _calc_yell(left, operand, right)
                unknown_monkeys.pop(monkey)
                break
            # print(json.dumps(known_monkeys, indent=4, sort_keys=True))
            # print(json.dumps(unknown_monkeys, indent=4, sort_keys=True))
        print(json.dumps(unknown_monkeys, indent=4, sort_keys=True))
    return known_monkeys

def _process_monkeys(entries: Strings) -> Monkeys:
    known_monkeys = {}
    unknown_monkeys = {}
    for entry in entries:
        words = entry.split()
        name = entry[:4]
        if len(words) == 2:
            known_monkeys[name] = int(words[1])
        else:
            unknown_monkeys[name] = words[1:]
    cycle = 0
    while unknown_monkeys:
        cycle += 1
        print(f"Cycle: {cycle}\nUnknown Monkeys: {len(unknown_monkeys)}\n")
        for monkey, (left, operand, right) in unknown_monkeys.items():
            if left in known_monkeys:
                left = known_monkeys.get(left)
            if right in known_monkeys:
                right = known_monkeys.get(right)
            if isinstance(left, int) and isinstance(right, int):
                known_monkeys[monkey] = _calc_yell(left, operand, right)
                print(f"Removing {monkey} from unknown_monkeys...")
                unknown_monkeys.pop(monkey)
                break
            # print(json.dumps(known_monkeys, indent=4, sort_keys=True))
            # print(json.dumps(unknown_monkeys, indent=4, sort_keys=True))
    return known_monkeys

def _calc_yell(left: int, operand: str, right: int) -> int:
    operations = {
        "+": lambda l, r: l + r,
        "-": lambda l, r: l - r,
        "*": lambda l, r: l * r,
        "/": lambda l, r: l // r,
    }
    return operations.get(operand)(left, right)

def _parse_input(input_: str) -> Strings:
    return [e for e in input_.split("\n") if e]

def main():
    assert _solve1(TEST) == 152
    Day()

if __name__ == '__main__':
    main()
