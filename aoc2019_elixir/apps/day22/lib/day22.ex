defmodule Day22 do
  alias Day22.Move, as: Move

  @behaviour Aoc.Problem

  @impl Aoc.Problem
  def input do
    File.read!("../day22/input.txt")
      |> String.split("\n")
      |> Enum.map(&Move.from_string/1)
  end

  @impl Aoc.Problem
  def part1(input) do
    size = 10007
    {m, b} = Move.perform_list(input, size)
    Integer.mod(2019 * m + b, size)
  end

  @impl Aoc.Problem
  def part2(input) do
    size = 119315717514047
    {m, b} = Move.undo_list(input, size) |> Move.repeat(size, 101741582076661)
    Integer.mod(2020 * m + b, size)
  end
end
