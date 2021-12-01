import functools


TEST_MAP = """..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#""".split('\n')


def read_map():
    rows = []
    with open('input.txt') as map_file:
        rows = [e.strip() for e in map_file]
    return rows


def navigate_map(area_map, dx=1, dy=1):
    x, y = 0, 0  # starting coordinates
    distance = len(area_map)  # mountain run length
    width = len(area_map[0])
    encounters = []
    while y < distance - dy:
        x = (x + dx) % width
        y += dy
        print(f'x={x}, y={y}, thing={area_map[y][x]}, distance={distance}, width={width}')
        encounters.append(area_map[y][x])
    return sum(1 if c == '#' else 0 for c in encounters)


print(TEST_MAP)
print(navigate_map(TEST_MAP, dx=3, dy=1))

area_map = read_map()
print(navigate_map(area_map, dx=3, dy=1))

movements = [
    {'dx': 1, 'dy': 1},
    {'dx': 3, 'dy': 1},
    {'dx': 5, 'dy': 1},
    {'dx': 7, 'dy': 1},
    {'dx': 1, 'dy': 2},
]

product = functools.reduce(lambda x, y: x * y, [navigate_map(area_map, **m) for m in movements])
print(product)
