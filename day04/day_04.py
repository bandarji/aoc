TEST_INPUT = """ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in""".split('\n')


def get_input():
    with open('input.txt') as input_file:
        return input_file.read().split('\n')

def valid_passport(person):
    required_fields = 'byr iyr eyr hgt hcl ecl pid'.split()
    person = {e.split(':')[0]: e.split(':')[1] for e in person.strip().split()}
    return all(f in person for f in required_fields)

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

print(count_valid_passports(TEST_INPUT))
print(count_valid_passports(get_input()))
