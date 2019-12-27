defmodule Day02 do
  @behaviour Aoc.Problem

  @impl Aoc.Problem
  def input do
    Intcode.Memory.from_string(File.read!("../day2/input.txt"))
  end

  @impl Aoc.Problem
  def part1(input) do
    real_input = List.replace_at(List.replace_at(input, 1, 12), 2, 2)
    computer = Intcode.Computer.async(real_input, nil)
    {:ok, memory} = Task.await(computer)
    Intcode.Memory.get(memory, 0)
  end

  @impl Aoc.Problem
  def part2(input) do
    all_inputs = for a <- 0..99, b <- 0..99, do: {a, b}

    Enum.find(all_inputs, fn {a, b} ->
      real_input = List.replace_at(List.replace_at(input, 1, a), 2, b)
      computer = Intcode.Computer.async(real_input, nil)
      {:ok, memory} = Task.await(computer)
      Intcode.Memory.get(memory, 0) == 19_690_720
    end)
  end
end
