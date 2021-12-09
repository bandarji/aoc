from pprint import pprint as pretty
from math import ceil

TEST_DATA = '16,1,2,0,4,2,7,1,2,14'


def get_input(filename):
    with open(filename) as file:
        return file.read()


def fuel_used(positions, end_position, day=1):
    if day == 1:
        return sum(abs(end_position - p) for p in positions)
    return sum(burn_fuel1(abs(end_position - p)) for p in positions)


def position_lowest_fuel(day=1, test=False):
    print(f'day={day}, test={test}')
    content = TEST_DATA if test else get_input('input.txt').strip()
    data = [int(n) for n in content.split(',')]
    if day == 1:
        positions = {p: fuel_used(data, p, day=day) for p in set(data)}
        fuel = min(positions.values())
        print(fuel)
        # for position, fuel in sorted(positions.items(), key=lambda r: 0 - r[1]):
        #     print(f'  position={position}, fuel={fuel}')
    else:
        find_minimum_fuel(data)


def find_minimum_fuel(data):
    minimum_fuel = [0, 0]
    min_pos, max_pos = min(data), max(data)
    for position in range(min_pos, max_pos):
        fuel = 0
        for crab in data:
            delta = abs(crab - position)
            fuel += ((delta * (delta + 1)) // 2)
        if minimum_fuel[1] == 0:
            minimum_fuel = [position, fuel]
        if fuel < minimum_fuel[1]:
            minimum_fuel = [position, fuel]
    print(minimum_fuel)


def burn_fuel1(moves):
    return ceil((moves * (moves + 1)) / 2)

def burn_fuel(remaining_moves):
    total = 0
    while remaining_moves > 0:
        total += remaining_moves
        remaining_moves -= 1
    return total

def test():
    tests = [16, 1, 2, 0, 4, 2, 7, 1, 2, 14, 5]
    position = 5
    fuel = 0
    for test in tests:
        fuel_ = burn_fuel(abs(test - position))
        fuel += fuel_
        print(f'Move from {test} to {position}: {fuel_}')
    print(f'Total: {fuel}\n')


def position_lowest_fuel1(day=2, test=False):
    content = TEST_DATA if test else get_input('input.txt')
    data = [int(n) for n in content.split(',')]
    fuel = 0
    minimum_fuel = None
    for i in range(max(data)):
        for u in data:
            for y, _ in enumerate(range(max(u, i) - min(u, i)), start=1):
                fuel += y
        if not minimum_fuel:
            minimum_fuel = fuel
        else:
            if minimum_fuel < fuel:
                minimum_fuel = fuel
        fuel = 0
    print(f'min fuel: {minimum_fuel}')
    return fuel

def main():
    test()
    position_lowest_fuel(day=1, test=True)
    position_lowest_fuel(day=1, test=False)
    position_lowest_fuel(day=2, test=True)
    position_lowest_fuel1(day=2, test=False)

if __name__ == '__main__':
    main()

# wrong: 99541469
# wrong: 4294967296
# wrong: 1146237523
# wrong: 1146237523
