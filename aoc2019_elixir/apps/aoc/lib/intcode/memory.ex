defmodule Intcode.Memory do
  @moduledoc """
  A simple server for storing the memory state of an Intcode computer.
  """

  use Agent

  @typedoc """
  A memory process that can be read from and updated as an Intcode program runs.
  """
  @type t :: pid

  @typedoc """
  An address to a location in memory.

  Memory addresses are any positive integer. It is possible to address memory that didn't have
  a value when the memory was created.
  """
  @type addr :: number

  @typedoc """
  A value stored at an address in the memory.

  This memory only stores integers.
  """
  @type value :: number

  @typedoc """
  A list of initial values for some Intcode memory.
  """
  @type data :: list(value)

  @spec start_link(data) :: Agent.on_start()
  def start_link(data) do
    Agent.start_link(fn -> data_to_map(data) end)
  end

  @doc """
  Retrieves a single value from an address in memory.

  If no value has been stored at the address, returns 0.
  """
  @spec get(t, addr) :: value
  def get(memory, addr) do
    Agent.get(memory, &Map.get(&1, addr, 0))
  end

  @doc """
  Sets an address in the memory to a given value.
  """
  @spec set(t, addr, value) :: value
  def set(memory, addr, value) do
    Agent.update(memory, &Map.put(&1, addr, value))
  end

  @doc """
  Dumps the internal state of the memory using `IO.inspect/1`.
  """
  @spec inspect(t) :: any
  def inspect(memory) do
    Agent.get(memory, &IO.inspect/1)
  end

  @doc """
  Reads memory data from a string of comma-separated integers.

  Produces memory data suitable for `start_link/1` or `Intcode.Computer.async/2`.
  """
  @spec from_string(String.t()) :: data
  def from_string(str) do
    for num <- String.split(str, ","), do: String.to_integer(num)
  end

  @spec data_to_map(data) :: %{optional(addr) => value}
  defp data_to_map(data) do
    for {value, index} <- Enum.with_index(data), into: %{}, do: {index, value}
  end
end
