TEST_DATA = """00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010""".split('\n')


def read_diags():
    with open('input.txt') as file:
        readings = [s.strip() for s in file]
    return readings


def find_most_common_bit(diags, position, tie='0'):
    counts = {}
    for diag in diags:
        bit = diag[position]
        counts[bit] = counts.get(bit, 0) + 1
    if counts.get('0', 0) == counts.get('1', 0):
        return tie
    else:
        return '1' if counts.get('1', 0) > counts.get('0', 0) else '0'


def o2_rating(diags):
    for position in range(len(diags[0]) + 1):
        print(f'position={position}, diags={diags}')
        if len(diags) == 1:
            break
        bit = find_most_common_bit(diags, position, tie='1')
        diags = [p for p in diags if p[position] == bit]
    return int(''.join(diags[0]), 2)


def co2_rating(diags):
    for position in range(len(diags[0]) + 1):
        print(f'position={position}, diags={diags}')
        if len(diags) == 1:
            break
        bit = '0' if find_most_common_bit(diags, position, tie='1') == '1' else '1'
        diags = [p for p in diags if p[position] == bit]
    return int(''.join(diags[0]), 2)


def main():
    diags = read_diags()
    o2_gen_rating = o2_rating(TEST_DATA)
    print()
    co2_scrubber_rating = co2_rating(TEST_DATA)
    print(f'o2_gen_rating={o2_gen_rating}, co2_scrubber_rating={co2_scrubber_rating}, product={o2_gen_rating * co2_scrubber_rating}')
    print()
    o2_gen_rating = o2_rating(diags)
    print()
    co2_scrubber_rating = co2_rating(diags)
    print(f'o2_gen_rating={o2_gen_rating}, co2_scrubber_rating={co2_scrubber_rating}, product={o2_gen_rating * co2_scrubber_rating}')



if __name__ == '__main__':
    main()
