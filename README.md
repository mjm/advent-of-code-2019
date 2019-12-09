# advent-of-code-2019

This repo has my [Advent of Code 2019][aoc2019] solutions.

[aoc2019]: https://adventofcode.com/2019

## Setup

My solutions are written in [Go][]. Each solution has library code and tests in the `day<n>` directory. The commands to produce the actual solutions are pretty small and live in `cmd/day<n>`.

The project uses [Bazel][] to build, run, and test. There's a Makefile as a shortcut for common actions.

[go]: https://golang.org
[bazel]: https://bazel.build

## Running a day's solution

For instance, to run the solution for day 6, run:

```
make day6
```

which will use `bazel run` to run that day's command with the proper input.

## Running tests

Each day's solution has unit tests that usually cover the examples given in the Advent of Code problem statement. To run the tests for every day, run:

```
make test
```
