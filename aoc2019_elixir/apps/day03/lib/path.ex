defmodule Day03.Path do
  @moduledoc """
  Functions for working with paths of wires.
  """

  @typedoc """
  A list of path segments that forms a complete wire path.
  """
  @type t :: list(segment)

  @typedoc """
  A single segment of a path, with a direction and distance.
  """
  @type segment :: {direction, integer}

  @typedoc """
  A cardinal direction for a segment of a path.
  """
  @type direction :: :up | :down | :left | :right

  @doc """
  Reads a wire's path from a description of its segments.

  ## Example

      iex> Day03.Path.from_string("R75,D30,R83,U83,L12")
      [
        {:right, 75},
        {:down, 30},
        {:right, 83},
        {:up, 83},
        {:left, 12},
      ]

  """
  @spec from_string(String.t) :: t
  def from_string(str) do
    String.split(str, ",") |> Enum.map(&segment_from_string/1)
  end

  @spec segment_from_string(String.t) :: segment
  defp segment_from_string(<<dir::utf8, dist::binary>>) do
    distance = String.to_integer(dist)

    case dir do
      ?U -> {:up, distance}
      ?D -> {:down, distance}
      ?L -> {:left, distance}
      ?R -> {:right, distance}
    end
  end

  @doc """
  Walk the path on the given map, marking the number of steps required to get to
  each point along the path. Returns the updated map.
  """
  @spec apply(Day03.Map.t, Day03.Map.id, t) :: Day03.Map.t
  def apply(m, id, path) do
    {m, _, _} =
      Enum.reduce(path, {m, {0, 0}, 0}, fn segment, {m, start, steps} ->
        apply_segment(m, id, segment, start, steps)
      end)

    m
  end

  defp apply_segment(m, id, segment, start, steps) do
    points = segment_points(segment, start)
    point_steps = Stream.iterate(steps + 1, &(&1 + 1))

    m =
      Enum.zip(points, point_steps)
      |> Enum.reduce(m, fn {point, steps}, m ->
        Day03.Map.set(m, point, id, steps)
      end)

    {m, List.last(points), steps + Enum.count(points)}
  end

  defp segment_points({dir, len}, {x0, y0}) do
    {dx, dy} = direction_offset(dir)
    for i <- 1..len, do: {x0 + i * dx, y0 + i * dy}
  end

  defp direction_offset(dir) do
    case dir do
      :up -> {0, -1}
      :down -> {0, 1}
      :left -> {-1, 0}
      :right -> {1, 0}
    end
  end
end
