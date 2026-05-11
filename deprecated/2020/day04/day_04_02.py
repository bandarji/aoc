import re
import string


REGEX_HEIGHT = re.compile(r'\d+(cm|in)$')
REGEX_HAIR_COLOR = re.compile(r'\#[0-9a-f]{6}$')
REGEX_PASSPORT = re.compile(r'\d{9}$')


INVALIDS = """eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007""".split('\n')

VALIDS = """pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719""".split('\n')


def str_to_num(info):
    if not info:
        return -1
    try:
        value = int(info)
    except ValueError:
        return -1
    return value


def check_byr(info):
    byr = str_to_num(info)
    return True if byr >= 1920 and byr <= 2002 else False

def check_iyr(info):
    iyr = str_to_num(info)
    return True if iyr >= 2010 and iyr <= 2020 else False

def check_eyr(info):
    eyr = str_to_num(info)
    return True if eyr >= 2020 and eyr <= 2030 else False

def check_hgt(info):
    match = re.match(REGEX_HEIGHT, info)
    if not match:
        return False
    digits = ''
    for char in info:
        if char in string.digits:
            digits += char
        else:
            break
    height = str_to_num(digits)
    if 'in' in info and height >= 59 and height <= 76:
        return True
    elif 'cm' in info and height >= 150 and height <= 193:
        return True
    else:
        return False

def check_hcl(info):
    match = re.match(REGEX_HAIR_COLOR, info)
    return True if match else False

def check_ecl(info):
    valid_colors = 'amb blu brn gry grn hzl oth'.split()
    return info and info in valid_colors

def check_pid(info):
    match = re.match(REGEX_PASSPORT, info)
    return True if match else False

def get_input():
    with open('input.txt') as input_file:
        return input_file.read().split('\n')

def valid_passport(person):
    required_fields = 'byr iyr eyr hgt hcl ecl pid'.split()
    person = {e.split(':')[0]: e.split(':')[1] for e in person.strip().split()}
    return all(
        [
            check_byr(person.get('byr', '')),
            check_iyr(person.get('iyr', '')),
            check_eyr(person.get('eyr', '')),
            check_hgt(person.get('hgt', '')),
            check_ecl(person.get('ecl', '')),
            check_hcl(person.get('hcl', '')),
            check_pid(person.get('pid', '')),
        ]
    )

def count_valid_passports(info):
    count = 0
    person = ''
    for text in info:
        if not text:
            count += 1 if valid_passport(person) else 0
            person = ''
        else:
            person += f' {text}'
    count += 1 if valid_passport(person) else 0
    return count

print(f'INVALIDS: valid={count_valid_passports(INVALIDS)}')
print(f'VALIDS: valid={count_valid_passports(VALIDS)}')
print(f'REAL: valid={count_valid_passports(get_input())}')

# 146 too high
