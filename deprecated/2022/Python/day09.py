from lib import get_input


TEST = """R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2"""

TEST2 = """R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20"""


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day09.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    unique_positions = set()
    head = [0, 0]
    tail = [0, 0]
    entries = [
        (command.split()[0], command.split()[1])
        for command in input_.split("\n")
        if command
    ]
    for direction, count in entries:
        for _ in range(int(count)):
            delta_row = 0 if direction in "LR" else 1 if direction == "D" else -1
            delta_column = 0 if direction in "UD" else 1 if direction == "R" else -1
            head = (head[0] + delta_row, head[1] + delta_column)
            tail = move_tail(head, tail)
            unique_positions.add(tuple(tail))
    return len(unique_positions)



def _solve2(input_: str) -> int:
    unique_positions = set()
    head = [0, 0]
    tail_spots = [(0, 0) for _ in range(9)]
    entries = [
        (command.split()[0], command.split()[1])
        for command in input_.split("\n")
        if command
    ]
    for direction, count in entries:
        for _ in range(int(count)):
            delta_row = 0 if direction in "LR" else 1 if direction == "D" else -1
            delta_column = 0 if direction in "UD" else 1 if direction == "R" else -1
            head = (head[0] + delta_row, head[1] + delta_column)
            tail_spots[0] = move_tail(head, tail_spots[0])
            for m in range(1, 9):
                tail_spots[m] = move_tail(tail_spots[m - 1], tail_spots[m])
                unique_positions.add(tail_spots[8])
    return len(unique_positions)


def move_tail(head, tail):
    delta_row = abs(head[0] - tail[0])
    delta_col = abs(head[1] - tail[1])
    if delta_row > 1 and delta_col > 1:
        tail = (head[0] - 1 if tail[0] < head[0] else head[0] + 1, head[1] - 1 if tail[1] < head[1] else head[1] + 1)
    elif delta_row > 1:
        tail = (head[0] - 1 if tail[0] < head[0] else head[0] + 1, head[1])
    elif delta_col > 1:
        tail = (head[0], head[1] - 1 if tail[1] < head[1] else head[1] + 1)
    else:
        pass
    return tail


def main():
    assert _solve1(TEST) == 13
    assert _solve2(TEST2) == 36
    Day()


if __name__ == '__main__':
    main()
