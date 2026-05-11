from lib import get_input


PART1_TESTS = [
]

PART2_TESTS = [
]


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day0X.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    return 0


def _solve2(input_: str) -> int:
    return 0


def main():
    for test, expected in PART1_TESTS:
        assert expected == _solve1(test)
    for test, expected in PART2_TESTS:
        assert expected == _solve2(test)
    Day()


if __name__ == '__main__':
    main()
