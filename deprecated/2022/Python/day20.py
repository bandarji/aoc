from collections import deque
from typing import List

from lib import get_input


Numbers = List[int]


TEST = """1
2
-3
3
-2
0
4"""


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day20.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    numbers = _parse_input(input_)
    count = len(numbers)
    positions = deque(
        [(number, index) for index, number in enumerate(numbers)]
    )
    for index, number in enumerate(numbers):
        cursor = positions.index((number, index))
        positions.remove((number, index))
        positions.rotate(0 - number)
        positions.insert(cursor, (number, index))
    repositioned = [p[0] for p in positions]
    initial = repositioned.index(0)
    answer = sum(repositioned[(i + initial) % count] for i in (1000, 2000, 3000))
    return answer


def _solve2(input_: str) -> int:
    numbers = _parse_input(input_)
    count = len(numbers)
    positions = deque(
        [(811589153 * number, index) for index, number in enumerate(numbers)]
    )
    for _ in range(10):
        for index, number in enumerate(numbers):
            number *= 811589153
            cursor = positions.index((number, index))
            positions.remove((number, index))
            positions.rotate(0 - number)
            positions.insert(cursor, (number, index))
    repositioned = [p[0] for p in positions]
    initial = repositioned.index(0)
    answer = sum(repositioned[(i + initial) % count] for i in (1000, 2000, 3000))
    return answer


def _parse_input(input_: str) -> Numbers:
    return [int(n) for n in input_.split("\n") if n]


def main():
    assert _solve1(TEST) == 3
    assert _solve2(TEST) == 1623178306
    Day()


if __name__ == '__main__':
    main()
