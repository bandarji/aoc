TEST_DATA = '3,4,3,1,2'


def read_data(filename):
    with open(filename) as file:
        return file.read()


def total_fish(counts):
    return sum(counts)


def count_fish(days=1, test=False):
    content = TEST_DATA if test else read_data('input.txt').strip()
    counts = [0 for _ in range(9)]
    for ttl in content.split(','):
        ttl = int(ttl)
        counts[ttl] += 1
    print(f'Initial: {counts}, total: {total_fish(counts)}')
    for day in range(days):
        zeros = counts[0]
        counts = counts[1:] + [zeros]
        counts[6] += zeros
        if (day + 1) % 50 == 0:
            print(f'  day={day + 1}, fish={total_fish(counts)}')
    print(f'TOTAL: {total_fish(counts)}')

def main():
    count_fish(days=18, test=True)
    count_fish(days=80, test=True)
    count_fish(days=80, test=False)
    count_fish(days=256, test=True)
    count_fish(days=256, test=False)

if __name__ == '__main__':
    main()
