defmodule Intcode.ListHandler do
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
