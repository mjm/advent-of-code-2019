defmodule Day03 do
  @behaviour Aoc.Problem

  @impl Aoc.Problem
  def input do
    String.split(File.read!("../day3/input.txt"), "\n")
  end

  @impl Aoc.Problem
  def part1(input) do
    paths = Enum.map(input, &Day03.Path.from_string(&1))
    m = Enum.zip(paths, [:a, :b]) |> Enum.reduce(%{}, fn {path, id}, m -> Day03.Path.apply(m, id, path) end)
    Day03.Map.nearest_intersection(m)
  end

  @impl Aoc.Problem
  def part2(input) do
    paths = Enum.map(input, &Day03.Path.from_string(&1))
    m = Enum.zip(paths, [:a, :b]) |> Enum.reduce(%{}, fn {path, id}, m -> Day03.Path.apply(m, id, path) end)
    Day03.Map.shortest_intersection(m)
  end
end
