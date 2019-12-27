defmodule Days do
  use Application

  @impl true
  def start(_type, _args) do
    Days.Supervisor.start_link(name: Days.Supervisor)
  end

  def solve(day, part) do
    input = day.input()

    case part do
      :part1 -> day.part1(input)
      :part2 -> day.part2(input)
    end
  end
end
