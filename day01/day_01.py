import functools


TARGET = 2020


def parse_input():
    with open('input.txt') as input_file:
        return [int(n) for n in input_file]


def adds_up(*numbers, target=2020):
    return sum(numbers) == target


def find_two(numbers):
    for cursor, number_1 in enumerate(numbers):
        for number_2 in numbers[:cursor] + numbers[cursor:]:
            if adds_up(*(number_1, number_2)):
                return number_1, number_2
    return -1, -1

def find_three(numbers):
    for cursor_1, number_1 in enumerate(numbers):
        for cursor_2, number_2 in enumerate(numbers[:cursor_1] + numbers[cursor_1:]):
            for cursor_3, number_3 in enumerate(n for n in numbers if n not in (number_1, number_2)):
                if adds_up(*(number_1, number_2, number_3), target=TARGET):
                    return number_1, number_2, number_3
    return -1, -1, -1

def display_answer(answers):
    print(f"{' * '.join(str(n) for n in answers)} = {functools.reduce(lambda x, y: x * y, answers)}")

numbers = parse_input()
answers = find_two(numbers)
if any(n == -1 for n in answers):
    print(f'Error: did not find two numbers adding up to {TARGET}')
else:
    display_answer(answers)

answers = find_three(numbers)
if any(n == -1 for n in answers):
    print(f'Error: did not find two numbers adding up to {TARGET}')
else:
    display_answer(answers)
