defmodule Day01 do
  @behaviour Aoc.Problem
  use GenServer

  @impl Aoc.Problem
  def input do
    Enum.map(String.split(File.read!("../day1/input.txt"), "\n"), &String.to_integer(&1))
  end

  @impl Aoc.Problem
  def part1(input) do
    GenServer.call(Day01, {:part1, input})
  end

  @impl Aoc.Problem
  def part2(input) do
    GenServer.call(Day01, {:part2, input})
  end

  def start_link(opts) do
    GenServer.start_link(__MODULE__, :ok, opts)
  end

  @impl true
  def init(:ok) do
    {:ok, %{}}
  end

  @impl true
  def handle_call({:part1, input}, _from, state) do
    {:reply, {:ok, Enum.sum(Enum.map(input, fn m -> Day01.Fuel.base_required(m) end))}, state}
  end

  @impl true
  def handle_call({:part2, input}, _from, state) do
    {:reply, {:ok, Enum.sum(Enum.map(input, fn m -> Day01.Fuel.total_required(m) end))}, state}
  end
end
