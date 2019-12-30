defmodule Day14 do
  alias Day14.Table, as: Table

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
    Table.fuel_possible(input, 1000000000000)
  end
end
