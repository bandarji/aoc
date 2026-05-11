from typing import List

from lib import get_input


Strings = List[str]

SNAFU_MAP = {
    "2": 2,
    "1": 1,
    "0": 0,
    "-": -1,
    "=": -2,
}

SNAFU_TESTS = [
    ("20", 10),
    ("2=", 8),
    ("2=-01", 976),
]

TEST_SNAFU_CONVERSION = """1=-0-2     1747
 12111      906
  2=0=      198
    21       11
  2=01      201
   111       31
 20012     1257
   112       32
 1=-1=      353
  1-12      107
    12        7
    1=        3
   122       37"""

TEST_SNAFUS = "\n".join(e.split()[0].strip() for e in TEST_SNAFU_CONVERSION.split("\n"))

TEST_INT_CONVERSION = """        1              1
        2              2
        3             1=
        4             1-
        5             10
        6             11
        7             12
        8             2=
        9             2-
       10             20
       15            1=0
       20            1-0
     2022         1=11-2
    12345        1-0---0
314159265  1121-1110-1=0"""


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day25.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> str:
    snafus = _parse_input(input_)
    return int_to_snafu(sum(snafu_to_int(snafu) for snafu in snafus))


def _solve2(input_: str) -> int:
    return 0


def _parse_input(input_: str) -> Strings:
    return [s.strip() for s in input_.split("\n") if s]


def snafu_to_int(snafu: str) -> int:
    total = 0
    for power, digit in enumerate(snafu[::-1]):
        total += SNAFU_MAP.get(digit, 0) * 5 ** power
    return total


def int_to_snafu(number: int) -> str:
    snafu = []
    while number > 0:
        snafu.append("012=-"[number % 5])
        number = (number + 2) // 5
    return "".join(reversed(snafu))


def main():
    for snafu, expected in SNAFU_TESTS:
        assert snafu_to_int(snafu) == expected
    for entry in TEST_SNAFU_CONVERSION.split("\n"):
        snafu, expected = [e.strip() for e in entry.split()]
        expected = int(expected)
        assert snafu_to_int(snafu) == expected
    for entry in TEST_INT_CONVERSION.split("\n"):
        number, expected = [e.strip() for e in entry.split()]
        number = int(number)
        assert int_to_snafu(number) == expected
    assert _solve1(TEST_SNAFUS) == "2=-1=0"
    Day()


if __name__ == '__main__':
    main()
