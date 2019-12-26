defmodule Intcode.ComputerTest do
  use ExUnit.Case, async: true

  test "computer can add numbers" do
    data = Intcode.memory_from_string("1,0,0,0,99")
    task = Intcode.Computer.async(data, nil)
    {:ok, memory} = Task.await(task)
    assert Intcode.Memory.get(memory, 0) == 2
  end

  test "computer can multiply numbers" do
    data = Intcode.memory_from_string("2,3,0,3,99")
    task = Intcode.Computer.async(data, nil)
    {:ok, memory} = Task.await(task)
    assert Intcode.Memory.get(memory, 3) == 6
  end

  test "computer can run multiple instructions" do
    data = Intcode.memory_from_string("1,1,1,4,99,5,6,0,99")
    task = Intcode.Computer.async(data, nil)
    {:ok, memory} = Task.await(task)
    assert Intcode.Memory.get(memory, 0) == 30
    assert Intcode.Memory.get(memory, 4) == 2
  end
end
