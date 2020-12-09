import re
from pprint import pprint as pretty


REGEX_CONTAINER_BAGS = re.compile(r'(.+)\ bags?\ contain\ (.*)\.')
REGEX_BAG_NUMBER = re.compile(r'\s*(\d+)\ (.*?)\ bags?')


TEST_RULES = """light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags."""


def read_file():
    with open('input.txt') as input_file:
        return input_file.read()


def important_entries(content):
    return re.findall(REGEX_CONTAINER_BAGS, content)


def bags_from_entries(entries):
    return {bag: {bag_: int(quantity) for quantity, bag_ in re.findall(REGEX_BAG_NUMBER, bag_info)}
            for bag, bag_info in entries}


def find_parents(bags):
    return {bag: {bag_ for bag_, qty in bags.items() if bag in qty} for bag in bags}


def search(bag_name):
    return parents[bag_name] | {k for b in parents[bag_name] for k in search(b)}


def count(bag_name):
    return sum(quantity * count(bag) for bag, quantity in bags.get(bag_name, {}).items()) + 1

entries = important_entries(TEST_RULES)
bags = bags_from_entries(entries)
parents = find_parents(bags)
print(len(search('shiny gold')))

entries = important_entries(read_file())
bags = bags_from_entries(entries)
parents = find_parents(bags)
print(len(search('shiny gold')))

print('Part Two:')
print(count('shiny gold') - 1)
