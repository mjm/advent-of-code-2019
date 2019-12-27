defmodule Intcode do
  def send_input(computer, value) do
    Intcode.Computer.send_input(computer, value)
  end

  def memory_from_string(str) do
    Intcode.Memory.from_string(str)
  end
end
