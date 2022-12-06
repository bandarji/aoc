from lib import get_input


PART1_TESTS = [
    ("bvwbjplbgvbhsrlpgdmjqwftvncz", 5),
    ("nppdvjthqldpwncqszvftbrmjlhg", 6),
    ("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10),
    ("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11),
]

PART2_TESTS = [
    ("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19),
    ("bvwbjplbgvbhsrlpgdmjqwftvncz", 23),
    ("nppdvjthqldpwncqszvftbrmjlhg", 23),
    ("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29),
    ("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26),
]


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day06.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    for i, _ in enumerate(input_):
        if i < 4:
            continue
        chars = input_[i - 4: i]
        if len(chars) == 4 and len(set(chars)) == 4:
            return i
    return 0


def _solve2(input_: str) -> int:
    for i, _ in enumerate(input_):
        if i < 4:
            continue
        chars = input_[i - 14: i]
        if len(chars) == 14 and len(set(chars)) == 14:
            return i
    return 0


def main():
    for test, expected in PART1_TESTS:
        assert expected == _solve1(test)
    for test, expected in PART2_TESTS:
        assert expected == _solve2(test)
    Day()


if __name__ == '__main__':
    main()
