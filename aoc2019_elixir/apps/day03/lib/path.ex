defmodule Day03.Path do
  def from_string(str) do
    String.split(str, ",") |> Enum.map(&segment_from_string/1)
  end

  defp segment_from_string(<<dir::utf8, dist::binary>>) do
    distance = String.to_integer(dist)

    case dir do
      ?U -> {:up, distance}
      ?D -> {:down, distance}
      ?L -> {:left, distance}
      ?R -> {:right, distance}
    end
  end

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
