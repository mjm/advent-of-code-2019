defmodule Day03.Map do
  @moduledoc """
  Functions for working with wiring maps.
  """

  @typedoc """
  A map of points that have been visited by a path of wires.
  """
  @type t :: map

  @typedoc """
  A point on the map in `{x, y}` coordinates.
  """
  @type point :: {integer, integer}

  @typedoc """
  An identifier for a particular wire path.

  This is used to identify which points on the map have the paths crossing.
  """
  @type id :: atom

  @doc """
  Sets the number of steps required to reach a point on the map along a given path.
  """
  @spec set(t, point, id, integer) :: t
  def set(m, point, id, steps) do
    Map.update(m, point, %{id => steps}, &Map.put_new(&1, id, steps))
  end

  @doc """
  Returns the point of the nearest intersection of two wires.

  The nearest point is determined by Manhattan distance from the origin `{0, 0}`.
  """
  @spec nearest_intersection(t) :: {point, integer}
  def nearest_intersection(m) do
    Enum.filter(m, fn {_, cells} -> Enum.count(cells) >= 2 end)
    |> Enum.map(fn {point, _} -> {point, point_distance(point)} end)
    |> Enum.min_by(fn {_, distance} -> distance end)
  end

  @doc """
  Returns the intersection of two wires that requires the fewest steps.

  This requires following the actual paths to know how many steps along each path
  it takes to get to the point.
  """
  @spec shortest_intersection(t) :: {point, integer}
  def shortest_intersection(m) do
    Enum.filter(m, fn {_, cells} -> Enum.count(cells) >= 2 end)
    |> Enum.map(fn {point, cells} ->
      {point, Enum.reduce(cells, 0, fn {_, steps}, acc -> acc + steps end)}
    end)
    |> Enum.min_by(fn {_, steps} -> steps end)
  end

  defp point_distance({x, y}) do
    abs(x) + abs(y)
  end
end
