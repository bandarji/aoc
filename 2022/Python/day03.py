from typing import List
import string

from lib import get_input


Strings = List[str]


PART1_TESTS = [
    ("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw", 157),
]

PRIORITIES = {
    k: v
    for v, k in enumerate(string.ascii_lowercase + string.ascii_uppercase, start=1)
}

PART2_TESTS = [
    ("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw", 70),
]


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day03.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    priority_total = 0
    for entry in input_.split("\n"):
        compartments = _split_rucksack(entry)
        duped_item = _find_duped_item(compartments)
        priority_total += PRIORITIES.get(duped_item, 0)
    return priority_total

def _solve2(input_: str) -> int:
    priority_total = 0
    content = input_.split("\n")
    for chunk in [content[i : i + 3] for i in range(0, len(content), 3)]:
        if not len(chunk) == 3:
            continue
        common_item = _find_common_item(chunk)
        priority_total += PRIORITIES.get(common_item, 0)
    return priority_total

def _find_common_item(rucksacks: Strings) -> str:
    rucksack_1, rucksack_2, rucksack_3 = rucksacks
    for item in rucksack_1:
        if item in rucksack_2 and item in rucksack_3:
            return item
    return "∅"


def _split_rucksack(rucksack: str) -> Strings:
    midpoint = len(rucksack) // 2
    return [rucksack[:midpoint], rucksack[midpoint:]]

def _find_duped_item(compartments: Strings) -> str:
    compartment_1, compartment_2 = compartments
    for item in compartment_1:
        if item in compartment_2:
            return item
    return "∅"

def main():
    for test, expected in PART1_TESTS:
        assert expected == _solve1(test)
    for test, expected in PART2_TESTS:
        assert expected == _solve2(test)
    Day()


if __name__ == '__main__':
    main()
