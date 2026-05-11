class Submarine:

    def __init__(self, position: int=0, depth: int=0, aim: int=0):
        self.position = position
        self.depth = depth
        self.aim = aim

    def move(self, direction, units):
        if direction == 'down':
            self.aim += units
        elif direction == 'up':
            self.aim -= units
        elif direction == 'forward':
            self.position += units
            self.depth += units * self.aim
        else:
            raise SystemExit('Bad direction')


def retrieve_commands() -> list:
    commands = []
    with open('input.txt') as file:
        for command in file:
            direction, units = command.split()
            units = int(units)
            commands.append([direction, units])
    return commands


def position_product(position, depth):
    return position * depth


def main():
    commands = retrieve_commands()
    sub = Submarine()
    for direction, units in commands:
        sub.move(direction, units)
    print(f'Position product: {sub.position} * {sub.depth} = {position_product(sub.position, sub.depth)}')


if __name__ == '__main__':
    main()
