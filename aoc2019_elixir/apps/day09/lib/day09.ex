defmodule Day09 do
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
