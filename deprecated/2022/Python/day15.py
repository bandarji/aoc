from pprint import pprint as pretty
from typing import List, Tuple

from lib import get_input

Device = Tuple[int, int]
Devices = List[Device]
Numbers = List[int]

TEST = """Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3"""


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day15.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


class Day15a:

    def __init__(self, input_: str, row: int):
        self.row = row
        self.sensors = []
        self.beacons = set()
        self.x_low = 0
        self.x_high = 0
        self.max_distance = 0
        self.parse_input(input_)
        self.run()

    def parse_input(self, input_: str) -> None:
        for entry in input_.split("\n"):
            if not entry:
                continue
            words = entry.split()
            sx = int(words[2][2:-1])
            sy = int(words[3][2:-1])
            bx = int(words[8][2:-1])
            by = int(words[9][2:])
            self.x_low = min(self.x_low, sx, bx)
            self.x_high = max(self.x_high, sx, bx)
            distance = taxi_distance(sx, sy, bx, by)
            self.max_distance = max(self.max_distance, distance)
            self.sensors.append((sx, sy, distance))
            self.beacons.add((bx, by))

    def run(self):
        total = 0
        y = self.row
        for x in range(self.x_low - self.max_distance, self.x_high + self.max_distance + 1):
            if (x, y) in self.beacons:
                continue
            match = False
            for (sx, sy, sr) in self.sensors:
                dist = taxi_distance(x, y, sx, sy)
                if dist <= sr:
                    total += 1
                    match = True
                    break
        self.answer = total

class Day15b:

    def __init__(self, input_: str, xy_max: int):
        self.xy_max = xy_max
        self.sensors = []
        self.beacons = set()
        self.x_low = 0
        self.x_high = 0
        self.max_distance = 0
        self.parse_input(input_)

    def parse_input(self, input_: str) -> None:
        for entry in input_.split("\n"):
            if not entry:
                continue
            words = entry.split()
            sx = int(words[2][2:-1])
            sy = int(words[3][2:-1])
            bx = int(words[8][2:-1])
            by = int(words[9][2:])
            self.x_low = min(self.x_low, sx, bx)
            self.x_high = max(self.x_high, sx, bx)
            distance = taxi_distance(sx, sy, bx, by)
            self.max_distance = max(self.max_distance, distance)
            self.sensors.append((sx, sy, distance))
            self.beacons.add((bx, by))

    def calc_coverage(self, y: int):
        coverage = []
        for sx, sy, sr in self.sensors:
            dist = abs(sy - y)
            if dist > sr:
                continue
            coverage.append((sx - (sr - dist), sx + (sr - dist)))
        dupe_checks = [False for _ in coverage]
        for i in range(len(coverage) - 1):
            for j in range(i + 1, len(coverage)):
                if i == j:
                    continue
                a, b = coverage[i]
                c, d = coverage[j]
                inspected = j
                if (a > c) or (a == c and d > b):
                    a, b, c, d = c, d, a, b
                    inspected = i
                if a <= c and b >= d:
                    dupe_checks[inspected] = True
        response = []
        for cursor in range(len(coverage)):
            if not dupe_checks[cursor]:
                response.append(coverage[cursor])
        return response


    def run(self):
        for y in range(self.xy_max + 1):
            coverage = self.calc_coverage(y)
            inspect = set()
            for cx, cy in coverage:
                if cx - 1 >= 0 and cx - 1 <= self.xy_max:
                    inspect.add(cx - 1)
                if cy + 1 >= 0 and cy + 1 <= self.xy_max:
                    inspect.add(cy + 1)
        for x in inspect:
            covered = False
            for cx, cy in coverage:
                if x >= cx and x <= cy:
                    covered = True
            if not covered:
                return x * xy_max + y
        return 0, 0



def taxi_distance(p1x, p1y, p2x, p2y) -> int:
    return abs(p2x - p1x) + abs(p2y - p1y)

def _solve1(input_: str, row: int) -> int:
    day15a = Day15a(input_, row)
    return day15a.answer

def _solve2(input_: str, xy_max) -> int:
    day15b = Day15b(input_, xy_max)
    return day15b.run()


def main():
    assert _solve1(TEST, 10) == 26
    # assert _solve2(TEST, 20) == 56000011
    print("Tests pass")
    print("Part 1:", _solve1(get_input("input_day15.txt"), 2_000_000))
    print("Part 2:", _solve2(get_input("input_day15.txt"), 4_000_000))

    # for test, expected in PART1_TESTS:
    #     assert expected == _solve1(test)
    # for test, expected in PART2_TESTS:
    #     assert expected == _solve2(test)
    # Day()


if __name__ == '__main__':
    main()
