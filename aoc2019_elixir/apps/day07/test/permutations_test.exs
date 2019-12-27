defmodule Day07.PermutationsTest do
  use ExUnit.Case, async: true

  test "generates permutations of a list of two elements" do
    perms = MapSet.new(Day07.Permutations.all(1..2))
    assert perms == MapSet.new([[1, 2], [2, 1]])
  end

  test "generates permutations of a list of three elements" do
    perms = MapSet.new(Day07.Permutations.all(3..5))

    assert perms ==
             MapSet.new([
               [3, 4, 5],
               [3, 5, 4],
               [4, 3, 5],
               [4, 5, 3],
               [5, 3, 4],
               [5, 4, 3]
             ])
  end
end
