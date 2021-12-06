TEST_DATA = """7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7""".split('\n')

class Game:

    def __init__(self, input_):
        self.numbers = input_[0].split(',')
        self.boards = self.build_boards(input_)
        self.current_number = 0

    def build_boards(self, input_):
        boards = []
        rows = []
        for i, line in enumerate(input_):
            row = line.split()
            if i == 0 or len(row) < 5:
                continue
            rows.append([[n, False] for n in row])
            if len(rows) == 5:
                boards.append(Board(rows))
                rows = []
        return boards

    def play_bingo(self):
        for number in self.numbers:
            self.current_number = number
            for board in self.boards:
                board.mark_number(number)
                if board.complete:
                    return board

class Board:

    def __init__(self, rows):
        self.rows = rows
        self.complete = False

    def display(self):
        for row in self.rows:
            for spot in row:
                number, flag = spot
                print(f' *{number:>2}' if flag else f'  {number:>2}', end='')
            print()
        print()

    def mark_number(self, picked_number):
        for i, row in enumerate(self.rows):
            for j, spot in enumerate(row):
                number, flag = spot
                if number == picked_number:
                    self.rows[i][j][1] = True
                    break
        self.won_game()

    def sum_unmarked(self):
        total = sum(
            int(self.rows[i][j][0]) if not self.rows[i][j][1] else 0
            for i in range(5) for j in range(5)
        )
        return total

    def won_game(self):
        for column in range(5):
            if all(self.rows[column][row][1] for row in range(5)):
                self.complete = True
        for row in range(5):
            if all(self.rows[column][row][1] for column in range(5)):
                self.complete = True


def read_input():
    with open('input.txt') as file:
        data = file.read().split('\n')
    return data


def game_on(test=True):
    data = read_input()
    game = Game(TEST_DATA if test else data)
    winner = game.play_bingo()
    winner.display()
    print(f'test={test}\nCurrent number: {game.current_number}')
    print(f'Unmarked sum: {winner.sum_unmarked()}')
    print(f'Product: {int(game.current_number) * winner.sum_unmarked()}\n')


def main():
    game_on()
    game_on(test=False)


if __name__ == '__main__':
    main()
