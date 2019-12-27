defmodule Day07.Connector do
  def async_with_settings(settings) do
    Enum.reverse(settings) |> Enum.reduce([], fn setting, amps ->
      case amps do
        [] -> [async(nil, setting)]
        [hd | tl] -> [async(hd.pid, setting) | [hd | tl]]
      end
    end)
  end

  def async_with_feedback(settings) do
    conns = async_with_settings(settings)
    set_next(List.last(conns).pid, List.first(conns).pid)
    conns
  end

  def async(next, setting) do
    Task.async(__MODULE__, :run, [next, setting])
  end

  def run(next, setting) do
    loop(next, nil, [setting], nil)
  end

  def send_input(connector, value) do
    send(connector, {:chain, value})
  end

  def set_next(connector, next) do
    send(connector, {:set_next, next})
  end

  defp loop(next, waiting_pid, queue, last_output) do
    receive do
      {:input, pid} ->
        case queue do
          [] -> loop(next, pid, [], last_output)
          [hd | tl] ->
            Intcode.send_input(pid, hd)
            loop(next, nil, tl, last_output)
        end

      {:chain, value} ->
        case waiting_pid do
          nil ->
            loop(next, nil, queue ++ [value], last_output)
          pid ->
            Intcode.send_input(pid, value)
            loop(next, nil, [], last_output)
        end

      {:output, _, value} ->
        case next do
          nil -> value
          _ ->
            send_input(next, value)
            loop(next, waiting_pid, queue, value)
        end
      
      {:set_next, new_next} ->
        loop(new_next, waiting_pid, queue, last_output)

      {:halt, _} -> last_output
    end
  end
end
