defmodule Day09 do
  @moduledoc """
  [Day 9: Sensor Boost](https://adventofcode.com/2019/day/9)
  """

  @behaviour Aoc.Problem

  @impl Aoc.Problem
  def input do
    Intcode.Memory.from_string(File.read!("../day9/input.txt"))
  end

  @impl Aoc.Problem
  def part1(input) do
    handler = Intcode.ListHandler.async([1])
    Intcode.Computer.async(input, handler.pid)
    Task.await(handler)
  end

  @impl Aoc.Problem
  def part2(input) do
    handler = Intcode.ListHandler.async([2])
    Intcode.Computer.async(input, handler.pid)
    Task.await(handler, 60000)
  end
end
