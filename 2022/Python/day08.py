from lib import get_input


TEST = """30373
25512
65332
33549
35390"""


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day08.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str) -> int:
    visible = 0
    grid = _build_grid(input_)
    for row, trees in enumerate(grid):
        if row == 0 or row == len(grid) - 1:
            continue
        visible += _count_visible(grid, row, trees)
    visible += 2 * len(grid[0]) + 2 * (len(grid) - 2)
    return visible


def _solve2(input_: str) -> int:
    high_score = 0
    grid = _build_grid(input_)
    for row, trees in enumerate(grid):
        for column, tree in enumerate(trees):
            if column == 0 or column == len(trees) - 1:
                continue
            if row == 0 or row == len(grid) - 1:
                continue
            scenic_score = _calc_scenic_score(grid, row, column, tree)
            if scenic_score > high_score:
                high_score = scenic_score
    return high_score


def _calc_scenic_score(grid, row, column, tree):
    left, right, above, below = 0, 0, 0, 0
    for c in range(column - 1, -1, -1):
        left += 1
        if tree <= grid[row][c]:
            break
    for c in range(column + 1, len(grid[0])):
        right += 1
        if tree <= grid[row][c]:
            break
    for r in range(row - 1, -1, -1):
        above += 1
        if tree <= grid[r][column]:
            break
    for r in range(row + 1, len(grid)):
        below += 1
        if tree <= grid[r][column]:
            break
    return left * right * above * below


def _count_visible(grid, row, trees):
    visible = 0
    for column, tree in enumerate(trees):
        if column == 0 or column == len(trees) - 1:
            continue
        visible += 1 if _tree_visible(grid, row, column, tree) else 0
    return visible


def _tree_visible(grid, row, column, tree) -> bool:
    left = max(grid[row][c] for c in range(column))
    right = max(grid[row][c] for c in range(column + 1, len(grid[0])))
    above = max(grid[r][column] for r in range(row))
    below = max(grid[r][column] for r in range(row + 1, len(grid)))
    return True if tree > left or tree > right or tree > above or tree > below else False


def _build_grid(input_: str):
    grid = []
    for row in input_.split("\n"):
        if row:
            grid.append([int(t) for t in row])
    return grid


def main():
    assert _solve1(TEST) == 21
    assert _solve2(TEST) == 8
    Day()


if __name__ == '__main__':
    main()
