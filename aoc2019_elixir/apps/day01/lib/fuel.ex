defmodule Day01.Fuel do
  def base_required(mass) do
    floor(mass / 3 - 2)
  end

  def total_required(mass) do
    fuel = base_required(mass)
    cond do
      fuel < 0 -> 0
      fuel > 0 -> fuel + total_required(fuel)
      true -> fuel
    end
  end
end
