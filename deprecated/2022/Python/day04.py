from typing import Tuple

from lib import get_input


PART1_TEST = """2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8"""

PART2_TESTS = [
]

Values = Tuple[int]


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day04.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    pairs = 0
    for entry in input_.split("\n"):
        if not "-" in entry:
            continue
        e1_start, e1_end, e2_start, e2_end = _parse_line(entry)
        pairs += 1 if (e1_start - e2_start) * (e1_end - e2_end) <= 0 else 0
    return pairs


def _solve2(input_: str) -> int:
    pairs = 0
    for entry in input_.split("\n"):
        if not "-" in entry:
            continue
        e1_start, e1_end, e2_start, e2_end = _parse_line(entry)
        pairs += 1 if e2_start <= e1_start <= e2_end or e1_start <= e2_start <= e1_end else 0
    return pairs


def _parse_line(entry: str) -> Values:
    h1, h2 = entry.split(",")
    e1_start, e1_end = h1.split("-")
    e2_start, e2_end = h2.split("-")
    return int(e1_start), int(e1_end), int(e2_start), int(e2_end)


def main():
    assert _solve1(PART1_TEST) == 2
    assert _solve2(PART1_TEST) == 4
    for test, expected in PART2_TESTS:
        assert expected == _solve2(test)
    Day()


if __name__ == '__main__':
    main()
