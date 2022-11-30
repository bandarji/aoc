import hashlib


PUZZLE_INPUT = "yzbqklnj"

PART1_TESTS = [
    ("abcdef", 609043),
    ("pqrstuv", 1048970),
]

PART2_TESTS = [
    ("^v", 3),
    ("^>v<", 3),
    ("^v^v^v^v^v", 11),
]


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = PUZZLE_INPUT if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    return first_hash(input_, 5)


def _solve2(input_: str) -> int:
    return first_hash(input_, 6)


def first_hash(string: str, zeroes: int) -> int:
    number = 0
    while True:
        test = f"{string}{number}".encode()
        if hashlib.md5(test).hexdigest().startswith("0" * zeroes):
            return number
        number += 1


def main():
    for test, expected in PART1_TESTS:
        assert expected == _solve1(test)
    # for test, expected in PART2_TESTS:
    #     assert expected == _solve2(test)
    Day()


if __name__ == '__main__':
    main()
