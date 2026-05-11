def get_depths(filename: str) -> list:
    with open(filename) as file:
        depths = [int(d) for d in file]
    return depths


def count_depth_increases(depths: list) -> int:
    return sum(
        1 if current > prior else 0
        for prior, current in zip(depths, depths[1:])
    )


depths = get_depths('input.txt')
increases = count_depth_increases(depths)
print(f'increases = {increases}')
