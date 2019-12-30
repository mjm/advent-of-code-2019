defmodule Days do
  @moduledoc """
  A runner for getting the solutions to any day's solutions.
  """
  use Application

  @impl true
  def start(_type, _args) do
    Days.Supervisor.start_link(name: Days.Supervisor)
  end

  @doc """
  Gets the solution to a single part of a day's problem.
  """
  @spec solve(module, :part1 | :part2) :: any
  def solve(day, part) do
    input = day.input()

    case part do
      :part1 -> day.part1(input)
      :part2 -> day.part2(input)
    end
  end
end
