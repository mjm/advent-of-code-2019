defmodule Aoc.Problem do
  @doc "Loads the input for the problem"
  @callback input() :: any()

  @doc "Performs part 1 of the problem and returns the result"
  @callback part1(any()) :: any()

  @doc "Performs part 2 of the problem and returns the result"
  @callback part2(any()) :: any()
end
