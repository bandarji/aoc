from lib import get_input


PART1_TESTS = [
    ("(())", 0),
    ("()()", 0),
    ("(((", 3),
    ("(()(()(", 3),
    ("))(((((", 3),
    ("())", -1),
    ("))(", -1),
    (")))", -3),
    (")())())", -3),
]

PART2_TESTS = [
    (")", 1),
    ("()())", 5),
]


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day01.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    floor = 0
    for move in input_:
        floor += 1 if move == "(" else -1 if move == ")" else 0
    return floor


def _solve2(input_: str) -> int:
    floor = 0
    for position, move in enumerate(input_, start=1):
        floor += 1 if move == "(" else -1 if move == ")" else 0
        if floor == -1:
            return position


def main():
    for test, expected in PART1_TESTS:
        assert expected == _solve1(test)
    for test, expected in PART2_TESTS:
        assert expected == _solve2(test)
    Day()


if __name__ == '__main__':
    main()
