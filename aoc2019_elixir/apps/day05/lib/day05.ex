defmodule Day05 do
  @behaviour Aoc.Problem

  @impl Aoc.Problem
  def input do
    Intcode.Memory.from_string(File.read!("../day5/input.txt"))
  end

  @impl Aoc.Problem
  def part1(input) do
    handler = Task.async(fn -> part1_loop([]) end)
    Intcode.Computer.async(input, handler.pid)
    Task.await(handler)
  end

  defp part1_loop(codes) do
    receive do
      {:input, pid} ->
        Intcode.send_input(pid, 1)
        part1_loop(codes)

      {:output, _, value} ->
        part1_loop([value | codes])

      {:halt, _} ->
        Enum.reverse(codes)
    end
  end

  @impl Aoc.Problem
  def part2(input) do
    handler = Task.async(&part2_loop/0)
    Intcode.Computer.async(input, handler.pid)
    Task.await(handler)
  end

  defp part2_loop do
    receive do
      {:input, pid} ->
        Intcode.send_input(pid, 5)
        part2_loop()

      {:output, _, value} ->
        value
    end
  end
end
