defmodule Aoc.Problem do
  @moduledoc """
  A behaviour for defining a module that solves a single day's problem for Advent of Code.

  Advent of Code problems have a single input that is used for a two-part problem. This
  behaviour defines a common structure for loading the input and providing solutions to
  each part of the problem for one day.
  """

  @doc "Loads the input for the problem."
  @callback input() :: any()

  @doc "Performs part 1 of the problem and returns the result."
  @callback part1(any()) :: any()

  @doc "Performs part 2 of the problem and returns the result."
  @callback part2(any()) :: any()
end
