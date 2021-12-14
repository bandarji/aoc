import json


REAL_DATA = """FK-gc
gc-start
gc-dw
sp-FN
dw-end
FK-start
dw-gn
AN-gn
yh-gn
yh-start
sp-AN
ik-dw
FK-dw
end-sp
yh-FK
gc-gn
AN-end
dw-AN
gn-sp
gn-FK
sp-FK
yh-gc"""

TEST_DATA = """start-A
start-b
A-c
A-b
b-d
A-end
b-end"""


def parse_content(content: str) -> dict:
    connections = {}
    for entry in content.split('\n'):
        pairing = entry.strip().split('-')
        for p_1, p_2 in zip(pairing, reversed(pairing)):
            if p_2 != 'start':
                if p_1 not in connections:
                    connections[p_1] = [p_2]
                else:
                    connections[p_1].append(p_2)
    connections.pop('end')
    return connections


def count_paths1(data: dict, test: bool=True, path=None):
    if not path:
        path = ['start']
    paths = 0
    for point in data[path[-1]]:
        if point.isupper() or not point in path:
            paths += 1 if point == 'end' else count_paths1(data, path=path + [point])
    return paths

def count_paths2(data: dict, test: bool=True, path=None):
    if not path:
        path = ['start']
    paths = 0
    for point in data[path[-1]]:
        if point == 'end':
            paths += 1
        else:
            paths += (count_paths2, count_paths1)[point.islower() and point in path](data, path=path + [point])
    return paths


def answer(test: bool=True, day: int=1):
    data = parse_content(TEST_DATA if test else REAL_DATA)
    if day == 1:
        print(f'test={test}, day={day}, paths={count_paths1(data, test=test)}')
        return
    print(f'test={test}, day={day}, paths={count_paths2(data, test=test)}')

def main():
    answer(test=True, day=1)
    answer(test=False, day=1)
    answer(test=True, day=2)
    answer(test=False, day=2)


if __name__ == '__main__':
    main()
