defmodule Day07 do
  @behaviour Aoc.Problem

  @impl Aoc.Problem
  def input do
    Intcode.Memory.from_string(File.read!("../day7/input.txt"))
  end

  @impl Aoc.Problem
  def part1(input) do
    Day07.Permutations.all(0..4) |> Enum.map(&part1_signal(input, &1)) |> Enum.max
  end

  defp part1_signal(input, settings) do
    signal(Day07.Connector.async_with_settings(settings), input)
  end

  @impl Aoc.Problem
  def part2(input) do
    Day07.Permutations.all(5..9) |> Enum.map(&part2_signal(input, &1)) |> Enum.max
  end

  defp part2_signal(input, settings) do
    signal(Day07.Connector.async_with_feedback(settings), input)
  end

  defp signal(conns, input) do
    # Queue up 0 for the first connector
    Day07.Connector.send_input(List.first(conns).pid, 0)

    # Start a Computer for each connector
    Enum.each(conns, fn conn ->
      Intcode.Computer.async(input, conn.pid)
    end)

    # Wait for the output from the last connector
    Task.await(List.last(conns))
  end
end
