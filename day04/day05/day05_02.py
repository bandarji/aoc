def bisect(side, front, back):
    half = (back - front) // 2
    if side == 'F' or side == 'L':
        back = front + half
    else:
        front = back - half
    return front, back


def find_my_seat(seats):
    for cursor, seat in enumerate(seats):
        if cursor == 0:
            continue  # skip first entry
        if seat - (seats[cursor - 1] + 1) != 0:
            return seat - 1
    return 0


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


seats = sorted(find_seat(e) for e in read_input())
seat = find_my_seat(seats)
print(seat)
