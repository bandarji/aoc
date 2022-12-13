from ast import literal_eval
from functools import cmp_to_key

from lib import get_input


TEST = """[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]"""


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day13.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_):
    indices_sum = 0
    pairs = _assemble_pairs(input_)
    for index, (first, second) in enumerate(pairs, start=1):
        result = _compare(first, second)
        if result > 0:
            indices_sum += index
    return indices_sum


def _compare(first, second):
    if isinstance(first, int):
        if isinstance(second, int):
            if first < second:
                return 1
            if second < first:
                return -1
            return 0
        return _compare([first], second)
    if isinstance(second, int):
        return _compare(first, [second])
    for cursor in range(min([len(first), len(second)])):
        result = _compare(first[cursor], second[cursor])
        if result < 0:
            return -1
        if result == 1:
            return 1
    if len(first) < len(second):
        return 1
    if len(second) < len(first):
        return -1
    return 0


def _assemble_pairs(input_):
    first = ""
    second = ""
    pairs = []
    for index, entry in enumerate([e for e in input_.split("\n") if e]):
        if index % 2:
            second = entry
            pairs.append((literal_eval(first), literal_eval(second)))
            first = ""
            second = ""
        else:
            first = entry
    return pairs


def _solve2(input_):
    indices_product = 1
    packets = _assemble_packets(input_)
    packets.append([[2]])
    packets.append([[6]])
    packets.sort(key=cmp_to_key(lambda first, second: _compare(first, second) * -1))
    for index, packet in enumerate(packets, start=1):
        if packet == [[2]] or packet == [[6]]:
            indices_product *= index
    return indices_product


def _assemble_packets(input_):
    packets = []
    for pairs in input_.split("\n\n"):
        first, second = pairs.rstrip().split("\n")
        packets.append(literal_eval(first))
        packets.append(literal_eval(second))
    return packets


def main():
    assert _solve1(TEST) == 13
    assert _solve2(TEST) == 140
    Day()


if __name__ == '__main__':
    main()
