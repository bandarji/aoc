import sys


TEST_INPUT = """2199943210
3987894921
9856789892
8767896789
9899965678"""


def read_input(filename: str) -> str:
    with open(filename) as file:
        return file.read()


def assemble_heightmap(heights: list) -> (dict, int, int):
    heightmap = {}
    for r, row in enumerate(heights.split('\n')):
        for c, height in enumerate(row.strip()):
            heightmap[(r, c)] = int(height)
    return heightmap, r, c


def find_low_points(heightmap: dict, max_x: int, max_y: int, coords: bool=False) -> list:
    points = []
    for x in range(max_x + 1):
        for y in range(max_y + 1):
            neighbors = [
                heightmap.get((x - 1, y), 10),
                heightmap.get((x + 1, y), 10),
                heightmap.get((x, y - 1), 10),
                heightmap.get((x, y + 1), 10),
            ]
            try:
                point = heightmap[(x, y)]
            except KeyError:
                continue
            if point < min(neighbors):
                points.append(point if not coords else (x, y))
    return points


def find_basins(coord, heightmap, denylist) -> (list, set):
    basins = []
    x, y = coord
    c_left = (x - 1, y)
    c_right = (x + 1, y)
    c_up = (x, y - 1)
    c_down = (x, y + 1)
    denylist.add(coord)

    conditions = [
        c_left in denylist,
        c_right in denylist,
        c_up in denylist,
        c_down in denylist,
    ]
    if all(conditions):
        return [coord], denylist
    if heightmap.get(coord) == 9:
        denylist.add(coord)
        return [coord], denylist
    conditions = [
        heightmap.get(c_left, 10) in {9, 10},
        heightmap.get(c_right, 10) in {9, 10},
        heightmap.get(c_up, 10) in {9, 10},
        heightmap.get(c_down, 10) in {9, 10},
    ]
    if any(conditions):
        denylist.add(coord)
    if c_left not in denylist:
        basins_left, denylist_left = find_basins(c_left, heightmap, denylist)
        basins += basins_left
        denylist = set.union(denylist, denylist_left)
    if c_right not in denylist:
        basins_right, denylist_right = find_basins(c_right, heightmap, denylist)
        basins += basins_right
        denylist = set.union(denylist, denylist_right)
    if c_up not in denylist:
        basins_up, denylist_up = find_basins(c_up, heightmap, denylist)
        basins += basins_up
        denylist = set.union(denylist, denylist_up)
    if c_down not in denylist:
        basins_down, denylist_down = find_basins(c_down, heightmap, denylist)
        basins += basins_down
        denylist = set.union(denylist, denylist_down)
    return basins, denylist


def answer(test: bool=True, day: int=1) -> None:
    input_ = TEST_INPUT if test else read_input('input.txt')
    heightmap, max_x, max_y = assemble_heightmap(input_)
    if day == 1:
        points = find_low_points(heightmap, max_x, max_y, coords=False)
        risk = sum(n + 1 for n in points)
        print(points, risk)
    else:
        coords = find_low_points(heightmap, max_x, max_y, coords=True)
        print(coords)
        raise SystemExit()
        denylist = set()
        basins = [
            find_basins(coord, heightmap, denylist)
            for coord in coords
        ]
        basin_sizes = sorted([len(basin[0]) for basin in basins])[-3:]
        print(basin_sizes)


def main():
    answer(test=True, day=1)
    answer(test=False, day=1)
    answer(test=True, day=2)


if __name__ == '__main__':
    sys.setrecursionlimit(2**15)
    main()
