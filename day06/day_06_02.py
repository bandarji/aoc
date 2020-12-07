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

def common_yes_counts(group):
    counts = {}
    people = len([g for g in group if g])
    for person in group:
        for vote in person:
            counts[vote] = counts.get(vote, 0) + 1
    votes = sum(1 for v in counts.values() if v == people)
    print(f'group={group}, people={people}, votes={votes}')
    return votes


def build_groups(text):
    return [t.split('\n') for t in text.split('\n\n')]

def read_input():
    with open('input.txt') as input_file:
        return input_file.read()

groups = build_groups(TEST_INPUT)
print(sum(common_yes_counts(group) for group in groups))

print(sum(common_yes_counts(group) for group in build_groups(read_input())))
