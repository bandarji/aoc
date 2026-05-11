from lib import get_input


DISPLAY = """##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....""".replace("\n", "")

PRETEST = """noop
addx 3
addx -5"""

PART1_TEST = """addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop"""


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day10.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions[0]}")
        for index in range(0, 240, 40):
            print("".join(self.solutions[1][index : index + 40]))



def _solve1(input_: str) -> int:
    strength, cycle, reg_x = 0, 0, 1
    cycles = {20, 60, 100, 140, 180, 220}

    for instruction in input_.split("\n"):
        if not instruction:
            continue
        if instruction == "noop":
            delta_x = 0
        else:
            delta_x = int(instruction.split()[1])
            cycle += 1
            if cycle in cycles:
                strength += reg_x * cycle
        cycle += 1
        if cycle in cycles:
            strength += reg_x * cycle
        reg_x += delta_x
    return strength


def _solve2(input_: str) -> str:
    cycle, reg_x = 0, 1
    display = ["." for _ in range(40 * 6)]
    tolerance = _within_bounds(reg_x)

    for instruction in input_.split("\n"):
        if not instruction:
            continue
        if instruction == "noop":
            delta_x = 0
            delta_c = 1
        else:
            delta_x = int(instruction.split()[1])
            delta_c = 2
        for offset in range(delta_c):
            position = cycle + offset
            if position % 40 in tolerance:
                display[position] = "#"
        cycle += delta_c
        reg_x += delta_x
        tolerance = _within_bounds(reg_x)
    return "".join(display)


def _within_bounds(reg_x: int):
    return [p for p in (reg_x - 1, reg_x, reg_x + 1) if p >= 0]


def main():
    assert _solve1(PART1_TEST) == 13140
    assert _solve2(PART1_TEST) == DISPLAY
    Day()


if __name__ == '__main__':
    main()
