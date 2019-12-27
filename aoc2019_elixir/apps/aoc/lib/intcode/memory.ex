defmodule Intcode.Memory do
  use Agent

  def start_link(data) do
    Agent.start_link(fn -> data_to_map(data) end)
  end

  def get(memory, index) do
    Agent.get(memory, &Map.get(&1, index))
  end

  def set(memory, index, value) do
    Agent.update(memory, &Map.put(&1, index, value))
  end

  def inspect(memory) do
    Agent.get(memory, &IO.inspect/1)
  end

  def from_string(str) do
    for num <- String.split(str, ","), do: String.to_integer(num)
  end

  defp data_to_map(data) do
    for {value, index} <- Enum.with_index(data), into: %{}, do: {index, value}
  end
end
