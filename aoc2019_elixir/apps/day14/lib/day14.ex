defmodule Day14 do
  @moduledoc """
  [Day 14: Space Stoichiometry](https://adventofcode.com/2019/day/14)
  """

  alias Day14.Table

  @behaviour Aoc.Problem

  @impl Aoc.Problem
  def input do
    Table.from_string(File.read!("../day14/input.txt"))
  end

  @impl Aoc.Problem
  def part1(input) do
    Table.required_ore(input, {1, "FUEL"})
  end

  @impl Aoc.Problem
  def part2(input) do
    Table.fuel_possible(input, 1_000_000_000_000)
  end
end
