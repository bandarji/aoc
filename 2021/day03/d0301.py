def read_diags():
    with open('input.txt') as file:
        readings = [s.strip() for s in file]
    return readings

def find_most_common_bit(diags, position):
    counts = {}
    for diag in diags:
        bit = diag[position]
        counts[bit] = counts.get(bit, 0) + 1
    return '1' if counts.get('1', 0) > counts.get('0', 0) else '0'


def main():
    gamma, epsilon = 0, 0
    diags = read_diags()
    bit_length = len(diags[0])
    most_common_bits = [
        find_most_common_bit(diags, p)
        for p in range(bit_length)
    ]
    gamma = int(''.join(most_common_bits), 2)
    epsilon = int(''.join('0' if b == '1' else '1' for b in most_common_bits), 2)
    product = gamma * epsilon
    print(f'gamma={gamma}, epsilon={epsilon}, product={product}')


if __name__ == '__main__':
    main()
