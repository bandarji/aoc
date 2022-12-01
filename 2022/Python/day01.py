TEST_INPUT = """1000
2000
3000

4000

5000
6000

7000
8000
9000

10000"""

PART1_TESTS = [
    (TEST_INPUT, 24000),
]

def heaviest_elf(input_: str) -> int:
    elves = []
    elf = 0
    for calories in input_.split("\n"):
        if not calories:
            elves.append(elf)
            elf = 0
        else:
            elf += int(calories)
    return max(elves)


def top_three(input_: str) -> int:
    elves = []
    elf = 0
    for calories in input_.split("\n"):
        if not calories:
            elves.append(elf)
            elf = 0
        else:
            elf += int(calories)
    elves.sort()
    return sum(elves[-3:])


def main():
    for input_, expected in PART1_TESTS:
        assert expected == heaviest_elf(input_)
    print(heaviest_elf(open("day01_input.txt").read()))
    print(top_three(open("day01_input.txt").read()))


if __name__ == '__main__':
    main()
