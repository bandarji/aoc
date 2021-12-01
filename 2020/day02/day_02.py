import re
import string


REGEX_PASSWORD_ENTRY = re.compile(r'(\d+)\-(\d+)\s+(\w)\:\s+(\w+)')


def valid_password_old(entry):
    match = re.match(REGEX_PASSWORD_ENTRY, entry)
    if match:
        floor = int(match.group(1))
        ceiling = int(match.group(2))
        char = match.group(3)
        password = match.group(4)
        count = 0
        count = sum(1 if c == char else 0 for c in password)
        print(f'char={char}, ceiling={ceiling}, floor={floor}, password={password}, count={count}')
        return 1 if count >= floor and count <= ceiling else 0
    return 0


def valid_password(entry):
    match = re.match(REGEX_PASSWORD_ENTRY, entry)
    total = 0
    if match:
        index_1 = int(match.group(1))
        index_2 = int(match.group(2))
        char = match.group(3)
        password = match.group(4)
        total = sum(
            [
                1 if password[index_1 - 1] == char else 0,
                1 if password[index_2 - 1] == char else 0,
            ]
        )
    return True if total == 1 else False



def read_passwords():
    with open('input.txt') as password_file:
        return sum(valid_password(entry) for entry in password_file)

print(read_passwords())

tests = [
    '1-3 a: abcde',
    '1-3 b: cdefg',
    '2-9 c: ccccccccc',
]
for test in tests:
    print(test, valid_password(test))
