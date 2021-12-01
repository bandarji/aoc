def get_depths(filename: str) -> list:
    with open(filename) as file:
        depths = [int(d) for d in file]
    return depths


def count_depth_increases(depths: list) -> int:
    """Consider three-reading sliding window"""
    windows = [
        r1 + r2 + r3
        for r1, r2, r3 in zip(depths, depths[1:], depths[2:])
    ]
    return sum(
        1 if current > prior else 0
        for prior, current in zip(windows, windows[1:])
    )


depths = get_depths('input.txt')
increases = count_depth_increases(depths)
print(f'increases = {increases}')
