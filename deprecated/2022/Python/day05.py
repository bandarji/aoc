from typing import List, Tuple

from lib import get_input

Instruction = Tuple[int]
Instructions = List[Instruction]
Stack = List[str]
Stacks = List[Stack]

PART1_TEST = """    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2"""


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day05.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> str:
    stacks = _parse_stacks(input_)
    for from_stack, to_stack in _parse_instructions(input_):
        stacks = _move_crate(stacks, from_stack, to_stack)
        # _render_stacks(stacks)
    return "".join(s[-1] for s in stacks)


def _solve2(input_: str) -> int:
    stacks = _parse_stacks(input_)
    for count, from_stack, to_stack in _parse_instructions_2(input_):
        stacks = _move_crates(stacks, count, from_stack, to_stack)
    return "".join(s[-1] for s in stacks)


def _move_crates(stacks: Stacks, count: int, from_stack: int, to_stack: int):
    tmp = [stacks[from_stack - 1].pop() for _ in range(count)]
    tmp.reverse()
    stacks[to_stack - 1] += tmp
    return stacks


def _parse_instructions_2(input_: str) -> Instructions:
    instructions = []
    for entry in input_.split("\n"):
        if entry.startswith("move "):
            _, count, _, from_stack, _, to_stack = entry.split()
            instructions.append((int(count), int(from_stack), int(to_stack)))
    return instructions


def _move_crate(stacks: Stacks, from_stack: int, to_stack: int) -> Stacks:
    stacks[to_stack - 1].append(stacks[from_stack - 1].pop())
    return stacks


def _parse_stacks(input_: str) -> Stacks:
    for entry in input_.split("\n"):
        if entry.startswith(" 1 "):
            stacks = [[] for _ in range(max(int(e) for e in entry.split()))]
            break
    for entry in input_.split("\n"):
        if "]" in entry:
            stacks = _parse_cargo(stacks, entry)
    return stacks

def _render_stacks(stacks: Stacks):
    print()
    for i, stack in enumerate(stacks, start=1):
        print(f"{i}: {' '.join(stack)}")


def _parse_cargo(stacks: Stacks, entry: str) -> Stacks:
    for index, char in enumerate(entry):
        if char == "[":
            stacks[index // 4].insert(0, entry[index + 1])
    return stacks


def _parse_instructions(input_: str) -> Instructions:
    instructions = []
    for entry in input_.split("\n"):
        if entry.startswith("move "):
            _, count, _, from_stack, _, to_stack = entry.split()
            for _ in range(int(count)):
                instructions.append((int(from_stack), int(to_stack)))
    return instructions


def main():
    assert _solve1(PART1_TEST) == "CMZ"
    assert _solve2(PART1_TEST) == "MCD"
    Day()


if __name__ == '__main__':
    main()
