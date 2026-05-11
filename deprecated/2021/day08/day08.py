TEST_DATA = """be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |
fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |
fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef |
cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega |
efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga |
gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf |
gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf |
cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd |
ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg |
gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |
fgae cfgab fg bagce"""


def read_content(filename):
    with open(filename) as file:
        return file.read()


def appears(test=False, day=1):
    content = TEST_DATA if test else read_content('input.txt')
    signals = [
        [set(signals_) for signals_ in entry.split()]
        for entry in content.split('\n')
    ]
    answer = 0
    if day == 1:
        for signal in signals:
            answer += sum(len(d) in {2, 3, 4, 7} for d in signal[-4:])
    else:
        for signal in signals:
            patterns = sorted(signal[:10], key=len)
            if not patterns: continue
            patterns[3:6] = sorted(patterns[3:6], key=lambda d: (patterns[1].issubset(d), len(patterns[2] & d)))
            patterns[6:9] = sorted(patterns[6:9], key=lambda d: (patterns[2].issubset(d), patterns[1].issubset(d)))
            patterns = [
                patterns[index]
                for index in [7, 0, 3, 5, 2, 4, 6, 1, 9, 8]
            ]
            rep = ''
            for digit in signal[-4:]:
                for i, pattern in enumerate(patterns):
                    if pattern == digit:
                        rep += f'{i}'
            answer += int(rep)
    print(f'test={test}, day={day}, answer={answer}')


def main():
    appears(test=True, day=1)
    appears(test=False, day=1)
    appears(test=False, day=1)
    appears(test=False, day=2)


if __name__ == '__main__':
    main()
