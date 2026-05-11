import itertools


TEST_INPUT = """35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576"""


def find_group(numbers, primer):
    run = len(numbers)
    cursor = 0  # tracking position
    advance = 0  # range indicator ahead of the cursor
    while cursor < run:
        while cursor + advance < run:
            span = numbers[cursor:cursor + advance]
            total = sum(span)
            if total < primer:
                advance += 1
                continue
            if total == primer:
                return min(span) + max(span)
            if total > primer:
                advance = 0
                break
        cursor += 1

def find_xmas_primer(numbers, preamble):
    for i, number in enumerate(numbers):
        if i < preamble:
            continue
        pass
        if adds_to(numbers[i - preamble: i], number):
            return number
    return -1


def adds_to(numbers, total):
    """Find any two numbers which add to the total."""
    return not any(sum(nums) == total for nums in list(itertools.combinations(numbers, 2)))


def read_input_file():
    with open('input.txt') as input_file:
        return [int(n) for n in input_file if n]


numbers = [int(n) for n in TEST_INPUT.split('\n') if n]
primer = find_xmas_primer(numbers, 5)
print(f'Part One: {primer}')
code = find_group(numbers, primer)
print(f'Part Two: {code}')

numbers = read_input_file()
primer = find_xmas_primer(numbers, 25)
print(f'Part One: {primer}')
code = find_group(numbers, primer)
print(f'Part Two: {code}')
