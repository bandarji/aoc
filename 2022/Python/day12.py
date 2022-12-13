from lib import get_input


ORD_A = ord("a")

TEST = """Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi"""


class OSPF:

    def __init__(self, input_: str):
        self.grid = [list(r) for r in input_.split("\n") if r]
        self.fix_grid()
        self.assemble_travels()

    def visit(self, current):
        for spot in self.travels[current]:
            if spot not in self.visited:
                if self.costs[current] + 1 < self.costs[spot]:
                    self.costs[spot] = self.costs[current] + 1
        self.notvisited.remove(current)
        self.visited.append(current)

    def play(self, start=None):
        self.costs = {k: 999 for k in self.notvisited}
        if not start:
            self.costs[self.start] = 0
        else:
            self.costs[start] = 0
        while self.end not in self.visited:
            current = min(self.notvisited, key=self.costs.get)
            self.visit(current)
        return self.costs[self.end]

    def starting_points(self):
        points = []
        for ri, row in enumerate(self.grid):
            for ci, cost in enumerate(row):
                if cost == 0:
                    spot = (ri, ci)
                    points.append(spot)
        return points

    def assemble_travels(self):
        self.travels = {}
        height, width = len(self.grid), len(self.grid[0])
        for ri, row in enumerate(self.grid):
            for ci, cost in enumerate(row):
                possible = []
                for d1 in (-1, 0, 1):
                    for d2 in (-1, 0, 1):
                        conditions = [
                            ri + d1 < height,
                            ci + d2 < width,
                            ri + d1 >= 0,
                            ci + d2 >= 0,
                            abs(d1) + abs(d2) == 1,
                        ]
                        if all(conditions) and self.grid[ri + d1][ci + d2] - self.grid[ri][ci] <= 1:
                            spot = (ri + d1, ci + d2)
                            possible.append(spot)
                self.travels[(ri, ci)] = possible

    def fix_grid(self):
        self.notvisited = []
        self.visited = []
        for ri, row in enumerate(self.grid):
            for ci, char in enumerate(row):
                spot = (ri, ci)
                if char == "S":
                    self.start = spot
                    char = "a"
                if char == "E":
                    self.end = spot
                    char = "z"
                self.grid[ri][ci] = ord(char) - ORD_A
                self.notvisited.append(spot)

class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day12.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    ospf = OSPF(input_)
    return ospf.play()


def _solve2(input_: str) -> int:
    costs = []
    ospf = OSPF(input_)
    starts = ospf.starting_points()
    for start in starts:
        ospf = OSPF(input_)
        cost = ospf.play(start=start)
        costs.append(cost)
    return min(costs)


def main():
    assert _solve1(TEST) == 31
    assert _solve2(TEST) == 29
    Day()


if __name__ == '__main__':
    main()
