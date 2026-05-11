import re

REGEX = re.compile(r'(\d+)\,(\d+)\s\-\>\s(\d+)\,(\d+)')

TEST_DATA = """0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2"""

def get_map_readings(filename):
    with open('input.txt') as file:
        return file.read()

def parse_readings(raw_readings):
    grid_size = 0
    readings = []
    for entry in raw_readings.split('\n'):
        try:
            line = [int(n) for n in re.match(REGEX, entry).groups()]
            grid_size = max([grid_size] + line)
            readings.append(line)
        except Exception as err:
            print(f'Error: ({err}), entry="{entry}"')
    return grid_size, readings

def init_grid(grid_size):
    return [[0 for i in range(grid_size + 1)] for _ in range(grid_size + 1)]

def map_lines(grid, data, day2=False):
    for s_x, s_y, e_x, e_y in data:
        if s_y == e_y:
            begin, end = min(s_x, e_x), max(s_x, e_x)
            for x in range(begin, end + 1):
                grid[x][s_y] += 1
        elif s_x == e_x:
            begin, end = min(s_y, e_y), max(s_y, e_y)
            for y in range(begin, end + 1):
                grid[s_x][y] += 1
        else:
            if day2:
                delta_x = 1 if s_x < e_x else -1
                delta_y = 1 if s_y < e_y else -1
                x = s_x
                y = s_y
                grid[s_x][s_y] += 1
                while x != e_x:
                    grid[x + delta_x][y + delta_y] += 1
                    x += delta_x
                    y += delta_y
    return grid

def count_overlaps(grid):
    total = 0
    for row in grid:
        for element in row:
            if element > 1:
                total += 1
    return total

def map_vents(test=False, day2=False):
    raw_readings = TEST_DATA if test else get_map_readings('input.txt')
    grid_size, data = parse_readings(raw_readings)
    grid = init_grid(grid_size)
    grid = map_lines(grid, data, day2=day2)
    overlaps = count_overlaps(grid)
    print(f'overlaps={overlaps}')

def main():
    map_vents(test=True)
    map_vents()
    map_vents(test=True, day2=True)
    map_vents(day2=True)

if __name__ == '__main__':
    main()
