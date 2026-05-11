import collections


PAIRS = {o: c for o, c in ['()', '[]', '{}', '<>']}
POINTS = {c: p for c, p in zip(')]}>', (3, 57, 1197, 25137))}

TEST_DATA = """[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]"""


def read_input(filename: str) -> str:
    with open(filename) as file:
        return file.read()



def day1(entries: list) -> (int, list, list):
    incomplete = []
    invalid = []
    invalid_positions = []
    for row_index, entry in enumerate(entries):
        stack = []
        for char_index, char in enumerate(entry):
            if char in PAIRS.values():
                if char in PAIRS[stack[-1]]:
                    stack.pop()
                else:
                    invalid.append(char)
                    invalid_positions.append(row_index)
                    break
            else:
                stack.append(char)
    return sum([POINTS[i] for i in invalid]), invalid_positions, incomplete


def answer(test: bool=True) -> None:
    input_ = TEST_DATA if test else read_input('input.txt')
    entries = [e.strip() for e in input_.split('\n')]
    score, bad_positions, stacks = day1(entries)
    print(f'test={test}, score_1={score}')


def answer2(test: bool=True) -> int:
    input_ = TEST_DATA if test else read_input('input.txt')
    scores = []
    for entry in input_.split('\n'):
        stack = collections.deque()
        for char in entry.strip():
            if char in '([{<':
                stack.appendleft(PAIRS[char])
            elif char != stack.popleft():
                break
        else:
            score = 0
            for char in stack:
                score = score * 5 + ')]}>'.index(char) + 1
            scores.append(score)
    return sorted(scores)[len(scores) // 2]


def main():
    answer(test=True)
    answer(test=False)
    for test in (True, False):
        print(f'Day 2: test={test} score={answer2(test=test)}')


if __name__ == '__main__':
    main()
