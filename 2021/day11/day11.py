from typing import List, Tuple


TEST_DATA = """5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526"""

REAL_DATA = """2566885432
3857414357
6761543247
5477332114
3731585385
1716783173
1277321612
3371176148
1162578285
6144726367"""


def get_neighbors(row: int, column: int) -> List[Tuple]:
    coord_diffs = [(i, j) for i in range(-1, 2) for j in range(-1, 2)]
    return [
        (row + dy, column + dx)
        for dy, dx in coord_diffs
        if all(
            [
                row + dy >= 0,
                row + dy <= 9,
                column + dx >= 0,
                column + dx <= 9,
            ]
        )
    ]


def answer(test: bool=True):
    entries = TEST_DATA.strip().split('\n') if test else REAL_DATA.strip().split('\n')
    grid = [
        [int(d) for d in entry]
        for entry in entries
    ]
    flashes, step, step2 = 0, 0, 0
    while True:
        flashed = []
        neighbors = []
        for r in range(10):
            for c in range(10):
                if grid[r][c] == 9:
                    grid[r][c] == 0
                    flashes += 1
                    flashed.append((r, c))
                    neighbors.extend(get_neighbors(r, c))
                else:
                    grid[r][c] += 1
        while neighbors:
            r, c = neighbors.pop(0)
            if (r, c) not in flashed:
                if grid[r][c] == 9:
                    grid[r][c] = 0
                    flashes += 1
                    flashed.append((r, c))
                    neighbors.extend(get_neighbors(r, c))
                else:
                    grid[r][c] += 1
        if step + 1 == 100:
            print(f'test={test}, day1 flashes={flashes}')
        if len(flashed) == 100:
            step2 = step + 1
            print(f'test={test}, day2 falshes={step2}')
            break
        step += 1


def main():
    answer(test=True)
    answer(test=False)

if __name__ == '__main__':
    main()
