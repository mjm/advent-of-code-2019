defmodule DaysTest do
  use ExUnit.Case
  doctest Days

  test "greets the world" do
    assert Days.hello() == :world
  end
end
