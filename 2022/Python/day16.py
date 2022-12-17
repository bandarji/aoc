from pprint import pprint as pretty

from lib import get_input


TEST = """Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II"""



class Day:

    def __init__(self, input_: str | None=None):
        self.input_ = get_input("input_day0X.txt") if not input_ else input_
        self.solutions = [0, 0]
        self.solve()
        self.display_solutions()

    def solve(self):
        self.solutions[0] = _solve1(self.input_)
        self.solutions[1] = _solve2(self.input_)

    def display_solutions(self):
        print(f"{self.solutions}")

class Volcano:

    def __init__(self, input_: str):
        self.input_ = input_
        self.parse_input()

    def parse_input(self):
        self.flows = {}
        self.paths = {}
        for entry in self.input_.split("\n"):
            if not entry:
                continue
            words = entry.split()
            valve = words[1]
            flow = int(words[4].split("=")[1][:-1])
            to_valves = [v.replace(",", "") for v in words[9:]]
            self.flows[valve] = flow
            self.paths[valve] = to_valves

    def run1(self, minutes: int):
        states = [(1, "AA", 0, ("", ))]
        visited = {}
        optimal = 0

        while states:
            minute, valve, cost, opened = states.pop()
            opened = {v for v in opened}
            if visited.get((minute, valve), -1) >= cost:
                continue
            visited[(minute, valve)] = cost
            if minute == minutes:
                optimal = max(optimal, cost)
                continue
            updated_cost = cost + sum(self.flows.get(valve, 0) for valve in opened)
            if self.flows[valve] > 0 and valve not in opened:
                opened.add(valve)
                updated_state = minute + 1, valve, updated_cost, tuple(opened)
                states.append(updated_state)
                opened = {v for v in opened if v != valve}
            for path in self.paths[valve]:
                updated_state = (minute + 1, path, updated_cost, tuple(opened))
                states.append(updated_state)
        return optimal

    def run2(self, minutes: int):
        states = [(1, "AA", "AA", 0, ("", ))]
        visited = {}
        optimal = 0

        max_flow = sum(self.flows.values())

        while states:
            minute, valve, elephant, cost, opened = states.pop()
            opened = {v for v in opened}
            if visited.get((minute, valve, elephant), -1) >= cost:
                continue
            visited[(minute, valve, elephant)] = cost
            if minute == minutes:
                optimal = max(optimal, cost)
                continue
            current_flow = sum(self.flows.get(valve, 0) for valve in opened)
            if current_flow >= max_flow:
                updated_cost = cost + current_flow
                while minute < minutes - 1:
                    minute += 1
                    updated_cost = cost + current_flow
                updated_state = (minute + 1, valve, elephant, updated_cost, tuple(opened))
                states.append(updated_state)
                continue
            if self.flows[valve] > 0 and valve not in opened:
                opened.add(valve)
                if self.flows[elephant] > 0 and elephant not in opened:
                    opened.add(elephant)
                    updated_cost = cost + sum(self.flows.get(valve, 0) for valve in opened)
                    states.append((minute + 1, valve, elephant, updated_cost, tuple(opened)))
                    opened = {v for v in opened if v != elephant}
                updated_cost = cost + sum(self.flows.get(valve, 0) for valve in opened)
                for path in self.paths[elephant]:
                    updated_state = (minute + 1, valve, path, updated_cost, tuple(opened))
                    states.append(updated_state)
                opened = {v for v in opened if v != valve}
            for path in self.paths[valve]:
                if self.flows[elephant] > 0 and elephant not in opened:
                    opened.add(elephant)
                    updated_cost = cost + sum(self.flows.get(valve, 0) for valve in opened)
                    updated_state = (minute + 1, path, elephant, updated_cost, tuple(opened))
                    states.append(updated_state)
                    opened = {v for v in opened if v != elephant}
                updated_cost = cost + sum(self.flows.get(valve, 0) for valve in opened)
                for elephant_path in self.paths[elephant]:
                    states.append((minute + 1, path, elephant_path, updated_cost, tuple(opened)))
        return optimal


def _solve1(input_: str, minutes: int) -> int:
    volcano = Volcano(input_)
    return volcano.run1(minutes + 1)

def _solve2(input_: str, minutes: int) -> int:
    volcano = Volcano(input_)
    a = volcano.run2(minutes)
    print(a)

def main():
    assert _solve1(TEST, 30) == 1651
    assert _solve2(TEST, 26) == 1707
    print(f"Part 1:", _solve1(get_input("input_day16.txt"), 30))
    # Day()


if __name__ == '__main__':
    main()
