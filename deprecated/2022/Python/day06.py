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
        self.solutions[0] = _solve(self.input_)
        self.solutions[1] = _solve(self.input_, window_size=14)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve(input_: str, window_size: int=4) -> int:
    for i, _ in enumerate(input_):
        if i < window_size:
            continue
        chars = input_[i - window_size: i]
        if len(chars) == window_size and len(set(chars)) == window_size:
            return i
    return 0


def main():
    for test, expected in PART1_TESTS:
        assert expected == _solve(test)
    for test, expected in PART2_TESTS:
        assert expected == _solve(test, window_size=14)
    Day()


if __name__ == '__main__':
    main()
