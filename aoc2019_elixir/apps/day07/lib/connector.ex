defmodule Day07.Connector do
  @moduledoc """
  A `t:Intcode.Computer.handler/0` that connects the input of one computer to the output
  of another.

  This module has functions for creating collections of handlers with different phase
  settings and wiring those handlers together correctly. A Connector will usually have
  a `next` connector, to which it will send any outputs it sees. If it doesn't have a
  `next` connector (because it's not set up for feedback and is the last in the chain),
  then when it receives an output from the computer, it will return it and finish its
  task.

  When a Connector receives a value from another connector, it adds it to a queue to be
  sent the next time the computer requests input.
  """

  @typedoc """
  A PID for a connector process.
  """
  @type t :: pid

  @typedoc """
  A list of phase settings for the amplifiers in an array.

  Each setting should be unique: the list should contain no duplicates.
  """
  @type phase_settings :: list(number)

  @doc """
  Create a list of connectors with the given phase settings.

  This will create as many connectors as there are settings in the list.
  Each connector will be set up to forward its outputs to the next connector
  in the list. The final connector will have no `next` connector and will
  simply return the output it receives from its `Intcode.Computer`.
  """
  @spec async_with_settings(phase_settings) :: list(Task.t())
  def async_with_settings(settings) do
    Enum.reverse(settings)
    |> Enum.reduce([], fn setting, amps ->
      case amps do
        [] -> [async(nil, setting)]
        [hd | tl] -> [async(hd.pid, setting) | [hd | tl]]
      end
    end)
  end

  @doc """
  Create a feedback loop of connectors with the given phase settings.

  Unlike with `async_with_settings/1`, the final connector in the list will
  send its output back to the first connector, creating a feedback loop. Only
  when the final `Intcode.Computer` halts will the output be returned.
  """
  @spec async_with_feedback(phase_settings) :: list(Task.t())
  def async_with_feedback(settings) do
    conns = async_with_settings(settings)
    set_next(List.last(conns).pid, List.first(conns).pid)
    conns
  end

  @spec async(t | nil, number) :: Task.t()
  defp async(next, setting) do
    Task.async(__MODULE__, :run, [next, setting])
  end

  @doc false
  @spec run(t | nil, number) :: number
  def run(next, setting) do
    loop(next, nil, [setting], nil)
  end

  @doc """
  Send an input value to a particular connector.

  This will add the value to that connector's queue so it can be sent to the computer
  the next time it requests input.

  Generally, the connectors use this to send information to each other, but it's also
  available to send the initial value to the first connector to start the process.
  """
  @spec send_input(t, number) :: any
  def send_input(connector, value) do
    send(connector, {:chain, value})
  end

  @spec set_next(t, t) :: any
  defp set_next(connector, next) do
    send(connector, {:set_next, next})
  end

  defp loop(next, waiting_pid, queue, last_output) do
    receive do
      {:input, pid} ->
        case queue do
          [] ->
            loop(next, pid, [], last_output)

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
          nil ->
            value

          _ ->
            send_input(next, value)
            loop(next, waiting_pid, queue, value)
        end

      {:set_next, new_next} ->
        loop(new_next, waiting_pid, queue, last_output)

      {:halt, _} ->
        last_output
    end
  end
end
