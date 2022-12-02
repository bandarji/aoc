from lib import get_input


PART1_TESTS = [
    ("A Y\nB X\nC Z", 15),
]

PART2_TESTS = [
    ("A Y\nB X\nC Z", 12),
]

POINTS = {
    "X": 1,
    "Y": 2,
    "Z": 3,
}

PLAYS = {
    "AX": 3,
    "AY": 6,
    "AZ": 0,
    "BX": 0,
    "BY": 3,
    "BZ": 6,
    "CX": 6,
    "CY": 0,
    "CZ": 3,
}

STRATEGIES = {
    "AX": 3,
    "AY": 1,
    "AZ": 2,
    "BX": 1,
    "BY": 2,
    "BZ": 3,
    "CX": 2,
    "CY": 3,
    "CZ": 1,
}


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day02.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        try:
            self.solutions[0] = _solve1(self.input_)
        except KeyError:
            raise SystemExit(f"KeyError discovered with {self.input}")
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    score = 0
    for game in input_.split("\n"):
        try:
            opponent, myself = game.split()
            score += POINTS[myself] + PLAYS[f"{opponent}{myself}"]
        except ValueError:
            continue
    return score


def _solve2(input_: str) -> int:
    score = 0
    for game in input_.split("\n"):
        try:
            opponent, result = game.split()
            play = 6 if result == "Z" else 3 if result == "Y" else 0
            score += play + STRATEGIES[f"{opponent}{result}"]
        except ValueError:
            continue
    return score


def main():
    for test, expected in PART1_TESTS:
        assert expected == _solve1(test)
    for test, expected in PART2_TESTS:
        assert expected == _solve2(test)
    Day()


if __name__ == '__main__':
    main()
