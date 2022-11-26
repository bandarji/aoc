from lib import get_input
from typing import List


Dimensions = List[int]


PART1_TESTS = [
    ("2x3x4", 58),
    ("1x1x10", 43),
]

PART2_TESTS = [
    ("2x3x4", 34),
    ("1x1x10", 14),
]


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day02.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    area = 0
    for entry in input_.split("\n"):
        l, w, h = sorted(_parse_entry(entry))
        area +=  l * w + 2 * l * w + 2 * w * h + 2 * h * l
    return area


def _solve2(input_: str) -> int:
    feet = 0
    for entry in input_.split("\n"):
        l, w, h = sorted(_parse_entry(entry))
        feet += 2 * (l + w) + l * w * h
    return feet


def _parse_entry(entry: str) -> Dimensions:
    try:
        dims = [int(e) for e in entry.split("x")]
    except Exception:
        dims = [0, 0, 0]
    return dims


def main():
    for test, expected in PART1_TESTS:
        assert expected == _solve1(test)
    for test, expected in PART2_TESTS:
        assert expected == _solve2(test)
    Day()


if __name__ == '__main__':
    main()
