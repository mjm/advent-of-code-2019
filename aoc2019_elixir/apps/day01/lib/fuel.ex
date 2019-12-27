defmodule Day01.Fuel do
  @doc ~S"""
  Calculates the base amount of fuel required to carry the given mass.

  ## Examples

    iex> Day01.Fuel.base_required(12)
    2
    iex> Day01.Fuel.base_required(14)
    2
    iex> Day01.Fuel.base_required(1969)
    654
    iex> Day01.Fuel.base_required(100756)
    33583

  """
  def base_required(mass) do
    floor(mass / 3 - 2)
  end

  @doc ~S"""
  Calculates the total amount of fuel required to carry the given mass.
  The total amount includes the mass for all of the fuel as well.

  ## Examples

    iex> Day01.Fuel.total_required(12)
    2
    iex> Day01.Fuel.total_required(14)
    2
    iex> Day01.Fuel.total_required(1969)
    966
    iex> Day01.Fuel.total_required(100756)
    50346

  """
  def total_required(mass) do
    fuel = base_required(mass)

    cond do
      fuel < 0 -> 0
      fuel > 0 -> fuel + total_required(fuel)
      true -> fuel
    end
  end
end
