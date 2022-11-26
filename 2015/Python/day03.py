from lib import get_input
from typing import List


Position = List[int]


PART1_TESTS = [
    (">", 2),
    ("^>v<", 4),
    ("^v^v^v^v^v", 2),
]

PART2_TESTS = [
    ("^v", 3),
    ("^>v<", 3),
    ("^v^v^v^v^v", 11),
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
    positions = {(0, 0)}
    position = [0, 0]
    for move in input_:
        position = _new_position(position, move)
        positions.add(tuple(position))
    return len(positions)


def _solve2(input_: str) -> int:
    houses = {(0, 0)}
    robot = [0, 0]
    santa = [0, 0]
    for turn, move in enumerate(input_):
        if turn % 2:
            robot = _new_position(robot, move)
            houses.add(tuple(robot))
        else:
            santa = _new_position(santa, move)
            houses.add(tuple(santa))
    return len(houses)


def _new_position(prior: Position, move: str) -> Position:
    current = [n for n in prior]
    if move == "^":
        current[1] += 1
    elif move == "v":
        current[1] -= 1
    elif move == ">":
        current[0] += 1
    elif move == "<":
        current[0] -= 1
    else:
        pass
    return current


def main():
    for test, expected in PART1_TESTS:
        assert expected == _solve1(test)
    for test, expected in PART2_TESTS:
        assert expected == _solve2(test)
    Day()


if __name__ == '__main__':
    main()
