defmodule Intcode.ListHandler do
  @moduledoc """
  A basic `t:Intcode.Computer.handler/0` implementation with a predetermined list of input values
  that collects a list of output values.

  This handler can be used for simple Intcode programs where the input is already known before
  the program runs and the output needs to be inspected after the computer halts.
  """

  def async(inputs) do
    Task.async(__MODULE__, :run, [inputs])
  end

  def run(inputs) do
    loop(inputs, [])
  end

  defp loop(inputs, outputs) do
    receive do
      {:input, pid} ->
        [hd | tl] = inputs
        Intcode.send_input(pid, hd)
        loop(tl, outputs)

      {:output, _, value} ->
        loop(inputs, [value | outputs])

      {:halt, _} ->
        Enum.reverse(outputs)
    end
  end
end
