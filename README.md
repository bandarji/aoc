# Advent of Code

[Advent of Code Site](https://adventofcode.com/)

I **love** Advent of Code, but I have never gotten an opportunity to code the
solutions in real time (daily). While I write code for work, and much of that
ends up performance-optimized, I rarely think about computer science
algorithms. Instead, much of my programming ends up working with libraries
and interfacing with APIs. The Advent of Code puzzles provide an opportunity
to go back in time, to think about lower-level approaches to problems. Some
work friends and I have private leaderboards or vow to work with a specific
language to solve puzzles, but I cannot compete with speedcoders. I mean,
I have seen times posted on the global leaderboard before I have fully
grokked the problem to solve.

In 2025, I decided to earn all the stars, completing puzzles in order,
starting with the first ones from 2015. And, I decided to solve them all
using the Go programming language. These live in the `/go` directory, which
contains a `Makefile` for build, execution and testing.

## Build and Execution

```
> make build && ./aoc --help
Advent of Code puzzles solved with Go. Use `make` to see build and execution options.

Usage:
  make YEAR=2015 DAY=1 run [flags]

Flags:
  -d, --day int    Day of the puzzle
  -h, --help       help for make
  -y, --year int   Year of the puzzle
```

Parsing the command line arguments comes from the Cobra package (see `/cmd`
for more information).

You do not need to build the executable in order to solve a specific day.
Instead, use `make run`.

```
> make YEAR=2016 DAY=7 run
Year=2016 Day=07 Part 1: 105 (13.804417ms)
Year=2016 Day=07 Part 2: 258 (10.719625ms)
```

## Tests

Using `make test` will provide test and coverage output. All functions adhere
to a naming standard to easily display test coverage for a specific day.

```
> make test | grep y16d07.go | column -t
github.com/bandarji/aoc/adventofcode/y16d07.go:16:  y16d07              100.0%
github.com/bandarji/aoc/adventofcode/y16d07.go:35:  y16d07ParseAddress  100.0%
github.com/bandarji/aoc/adventofcode/y16d07.go:54:  y16d07IsTLS         100.0%
github.com/bandarji/aoc/adventofcode/y16d07.go:64:  y16d07HasABBA       100.0%
github.com/bandarji/aoc/adventofcode/y16d07.go:75:  y16d07IsSSL         100.0%
github.com/bandarji/aoc/adventofcode/y16d07.go:87:  y16d07FindABAs      100.0%
```

I typically cannot solve these in real time because by the end of the first
week my wife transitions from faces to foot tapping to stern words to making
me fear for my life. :kissing_heart:

# Progress

| Year | Stars | Notes |
| :--- | :--- | :--- |
| 2024 | 33 | Go |
| 2023 | 47 | Go |
| 2022 | 25 | Go, Rust and Python |
| 2021 | 23 | Go and Python |
| 2020 | 19 | Python - first year to join |
| 2019 |  2 | C - wrote solution in 2020 |
| 2018 |  0 | Did not participate |
| 2017 |  0 | Did not participate |
| 2016 | 23 | Go |
| 2015 | 50 | Go, Rust and Python |

## 2015 Complete!

![AOC 2015 Complete](assets/blog-aoc-2015-complete.png)
