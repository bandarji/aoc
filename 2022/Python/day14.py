from pprint import pprint as pretty
from typing import List, Tuple

from lib import get_input

Point = Tuple[int, int]
Points = List[Point]

TEST = """498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9"""


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day14.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")

class Cave:

    def __init__(self, input_: str, start: Point, day: int):
        self.day = day
        self.input_ = input_
        self.floor = 0
        self.mapping = {}
        self.draw_cave()
        if day == 1:
            self.pour_sand1(start)
        else:
            self.pour_sand2(start)
        self.count_sand()

    def pour_sand1(self, start):
        for _ in range(self.floor * (_max_x_coord(self.mapping) - _min_x_coord(self.mapping)) + 1):
            x, y = start
            while True:
                if y >= self.floor:
                    break
                if (x, y + 1) not in self.mapping:
                    y += 1
                elif (x - 1, y + 1) not in self.mapping:
                    x -= 1
                    y += 1
                elif (x + 1, y + 1) not in self.mapping:
                    x += 1
                    y += 1
                else:
                    self.mapping[(x, y)] = "o"
                    break
            # self.render()

    def pour_sand2(self, start):
        while True:
            x, y = start
            while True:
                if y > self.floor:
                    break
                if (x, y + 1) not in self.mapping:
                    y += 1
                elif (x - 1, y + 1) not in self.mapping:
                    x -= 1
                    y += 1
                elif (x + 1, y + 1) not in self.mapping:
                    x += 1
                    y += 1
                else:
                    self.mapping[(x, y)] = "o"
                    break
            if self.mapping.get(start) == "o":
                break
            # self.render()

    def count_sand(self):
        self.grains = sum(1 if spot == "o" else 0 for spot in self.mapping.values())

    def render(self):
        for y in range(self.floor + 1):
            print(
                "".join(
                    self.mapping.get((x, y), ".")
                    for x in range(_min_x_coord(self.mapping), _max_x_coord(self.mapping) + 1)
                )
            )
        print()

    def draw_cave(self):
        for wall_info in self.input_.split("\n"):
            if not wall_info:
                continue
            points = _get_rockwall_points(wall_info)
            x, y = points.pop(0)
            self.floor = max(self.floor, y)
            self.mapping[(x, y)] = "#"
            for next_x, next_y in points:
                while x != next_x or y != next_y:
                    delta_x = (1 if next_x > x else 0) - (1 if x > next_x else 0)
                    delta_y = (1 if next_y > y else 0) - (1 if y > next_y else 0)
                    x += delta_x
                    y += delta_y
                    self.mapping[(x, y)] = "#"
                    self.floor = max(self.floor, y)
        if self.day == 2:
            self.floor += 2
            for x in range(_min_x_coord(self.mapping) - 10_000, _max_x_coord(self.mapping) + 10_000):
                self.mapping[(x, self.floor)] = "#"

def _min_x_coord(mapping: Points) -> int:
    return min(c[0] for c in mapping)

def _max_x_coord(mapping: Points) -> int:
    return max(c[0] for c in mapping)


def _solve1(input_: str) -> int:
    cave = Cave(input_, (500, 0), 1)
    return cave.grains


def _solve2(input_: str) -> int:
    cave = Cave(input_, (500, 0), 2)
    return cave.grains


def _get_rockwall_points(wall_info: str) -> Points:
    points = wall_info.split(" -> ")
    points = [tuple([int(n) for n in point.split(",")]) for point in points]
    return points


def main():
    assert _solve1(TEST) == 24
    assert _solve2(TEST) == 93
    print("Tests passed")
    Day()


if __name__ == '__main__':
    main()
