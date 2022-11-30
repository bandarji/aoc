from lib import get_input

PART1_TESTS = [
    ("ugknbfddgicrmopn", 1),
    ("aaa", 1),
    ("jchzalrnumimnmhp", 0),
    ("haegwjzuvuyypxyu", 0),
    ("dvszwmarrgswjxmb", 0),
]

PART2_TESTS = [
    ("qjhvhtzxzqqjkmpb", 1),
    ("xxyxx", 1),
    ("uurcxstgmygtbstg", 0),
    ("ieodomkazucvgmuy", 0),
]

BAD_STRINGS = "ab cd pq xy".split()
VOWELS = "aeiou"
ALPHABET = "abcdefghijklmnopqrstuvwxyz"


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day05.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    nice = 0
    for string in input_.split("\n"):
        if _is_nice(string):
            nice += 1
    return nice


def _solve2(input_: str) -> int:
    nice = 0
    for string in input_.split("\n"):
        if _is_new_nice(string):
            nice += 1
    return nice


def _three_vowels(string: str) -> bool:
    count = sum(1 if c in VOWELS else 0 for c in string)
    return True if count >= 3 else False


def _adjacent_letters(string: str) -> bool:
    return any(f"{c}{c}" in string for c in ALPHABET)


def _is_nice(string: str) -> bool:
    if any(bs in string for bs in BAD_STRINGS):
        return False
    if all([_three_vowels(string), _adjacent_letters(string)]):
        return True
    return False


def _is_new_nice(string: str) -> bool:
    return all([_two_twice(string), _encapsulates(string)])


def _encapsulates(string: str) -> bool:
    for index, char in enumerate(string):
        try:
            if char == string[index + 2]:
                return True
        except IndexError:
            break
    return False


def _two_twice(string: str) -> bool:
    for index, char in enumerate(string):
        if index == 0:
            continue
        search = f"{string[index - 1]}{char}"
        if search in string[index + 1:]:
            return True
    return False


def main():
    for test, expected in PART1_TESTS:
        assert expected == _solve1(test)
    for test, expected in PART2_TESTS:
        assert expected == _solve2(test)
    Day()


if __name__ == '__main__':
    main()
