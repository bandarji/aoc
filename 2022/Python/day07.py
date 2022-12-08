from typing import Dict

from lib import get_input


FileSystem = Dict[str, int]

PART1_TEST = """$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k"""


class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day07.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")


def _solve1(input_: str, recover: int=100_000) -> int:
    total_space_used = 0
    fsystem = _assemble_tree(input_)
    for directory, space in sorted(fsystem.items(), key=lambda e: e[1]):
        if space > 100_000:
            break
        total_space_used += space
    return total_space_used


def _solve2(input_: str, capacity: int=70_000_000, min_free: int=30_000_000) -> int:
    fsystem = _assemble_tree(input_)
    used = fsystem.get("/", 0)
    to_free = used - (capacity - min_free)
    candidates = []
    for directory, space in sorted(fsystem.items(), key=lambda e: e[1]):
        if space >= to_free:
            candidates.append(space)
    return min(candidates)


def _assemble_tree(input_: str) -> FileSystem:
    fsystem = {}
    path = []
    for replay_line in input_.split("\n"):
        if not replay_line:
            continue
        if replay_line.startswith("$ cd "):
            elements = replay_line.split()
            if elements[2] == "..":
                path.pop()
            else:
                path.append(elements[2])
        elif replay_line.startswith(("$ ls", "dir")):
            continue
        else:
            fsize = int(replay_line.split()[0])
            for d in range(len(path)):
                fpath = "/".join(path[:d + 1])
                fsystem[fpath] = fsystem.get(fpath, 0) + fsize
    return fsystem


def main():
    assert _solve1(PART1_TEST) == 95437
    assert _solve2(PART1_TEST) == 24933642
    Day()


if __name__ == '__main__':
    main()
