defmodule Day03.Map do
  def set(m, point, id, steps) do
    Map.update(m, point, %{id => steps}, &Map.put_new(&1, id, steps))
  end

  def count_at(m, point) do
    Enum.count(Map.get(m, point, %{}))
  end

  def nearest_intersection(m) do
    Enum.filter(m, fn {_, cells} -> Enum.count(cells) >= 2 end) |>
      Enum.map(fn {point, _} -> {point, point_distance(point)} end) |>
      Enum.min_by(fn {_, distance} -> distance end)
  end

  def shortest_intersection(m) do
    Enum.filter(m, fn {_, cells} -> Enum.count(cells) >= 2 end) |>
      Enum.map(fn {point, cells} -> {point, Enum.reduce(cells, 0, fn {_, steps}, acc -> acc + steps end)} end) |>
      Enum.min_by(fn {_, steps} -> steps end)
  end

  defp point_distance({x, y}) do
    abs(x) + abs(y)
  end
end
