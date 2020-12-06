TEST = 'FBFBBFFRLR'

def bisect(side, front, back):
    half = (back - front) // 2
    if side == 'F' or side == 'L':
        back = front + half
    else:
        front = back - half
    return front, back


def find_seat(seat):
    front, back = 0, 127
    for char in seat[:7]:
        front, back = bisect(char, front, back)
    row = front
    front, back = 0, 7
    for char in seat[7:]:
        front, back = bisect(char, front, back)
    seat_id = row * 8 + front
    return seat_id

def read_input():
    with open('input.txt') as input_file:
        content = [e.strip() for e in input_file]
    return content

tests = [
    ('FBFBBFFRLR', 357),
    ('BFFFBBFRRR', 567),
    ('FFFBBBFRRR', 119),
    ('BBFFBBFRLL', 820),
]

for input_, expected in tests:
    print(f'Sent: {input_}, Expected: {expected}, Received: {find_seat(input_)}')

high_seat = max(find_seat(e) for e in read_input())
print(high_seat)
