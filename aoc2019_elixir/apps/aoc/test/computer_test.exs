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

  test "computer can add with immediate mode" do
    data = Intcode.memory_from_string("1101,100,-1,4,0")
    task = Intcode.Computer.async(data, nil)
    {:ok, memory} = Task.await(task)
    assert Intcode.Memory.get(memory, 4) == 99
  end

  test "computer can multiply with immediate mode" do
    data = Intcode.memory_from_string("1002,4,3,4,33")
    task = Intcode.Computer.async(data, nil)
    {:ok, memory} = Task.await(task)
    assert Intcode.Memory.get(memory, 4) == 99
  end

  test "computer can handle single input and single output" do
    data = Intcode.memory_from_string("3,9,8,9,10,9,4,9,99,-1,8")

    loop = fn loop ->
      receive do
        {:input, pid} ->
          Intcode.send_input(pid, 8)
          loop.(loop)

        {:output, _, val} ->
          val
      end
    end

    handler = Task.async(fn -> loop.(loop) end)
    Intcode.Computer.async(data, handler.pid)
    assert Task.await(handler) == 1
  end

  test "computer can handle relative base" do
    data = Intcode.memory_from_string("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")

    loop = fn loop, values ->
      receive do
        {:output, _, val} ->
          loop.(loop, [val | values])

        {:halt, _} ->
          Enum.reverse(values)
      end
    end

    handler = Task.async(fn -> loop.(loop, []) end)
    Intcode.Computer.async(data, handler.pid)

    assert Task.await(handler) == [
             109,
             1,
             204,
             -1,
             1001,
             100,
             1,
             100,
             1008,
             100,
             16,
             101,
             1006,
             101,
             0,
             99
           ]
  end
end
