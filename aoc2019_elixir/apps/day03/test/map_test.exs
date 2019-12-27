defmodule Day03.MapTest do
  use ExUnit.Case, async: true

  test "can set initial values for points" do
    m = %{} |>
      Day03.Map.set({1,0}, :a, 1) |>
      Day03.Map.set({1,1}, :a, 2) |>
      Day03.Map.set({2,1}, :a, 3)

    assert Enum.count(m) == 3
  end

  test "can set values for different paths on a cell" do
    m = %{} |>
      Day03.Map.set({1,0}, :a, 1) |>
      Day03.Map.set({1,0}, :b, 1) |>
      Day03.Map.set({1,0}, :b, 5)

    assert Enum.count(m) == 1
    assert Enum.count(m[{1,0}]) == 2
  end

  test "can apply a path" do
    path = [{:up, 3}, {:right, 2}]
    m = Day03.Path.apply(%{}, :a, path)
    
    assert Enum.count(m) == 5
  end
end
