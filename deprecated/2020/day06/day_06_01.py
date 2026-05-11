import string


TEST_INPUT = """abc

a
b
c

ab
ac

a
a
a
a

b"""


def count_yes_answers(text):
    return len(set(c for c in text if c in string.ascii_lowercase))

def count_all_yes_answers(text):
    return sum(count_yes_answers(t) for t in text.split('\n\n'))

def read_input():
    with open('input.txt') as input_file:
        return input_file.read()

print(count_all_yes_answers(TEST_INPUT))
print(count_all_yes_answers(read_input()))
