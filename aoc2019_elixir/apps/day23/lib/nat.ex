require Logger

defmodule Day23.NAT do
  @moduledoc """
  A process that is responsible for keeping the network active.

  The NAT receives packets addressed to address `255` during the course of normal
  network traffic. When the router informs the NAT that all addresses are idling,
  the NAT may rebroadcast the last packet it received to the computer at address
  `0` to restart traffic.

  If the NAT would send computer `0` two packets in a row with the same `y` value,
  it will instead return the `{x, y}` pair from its task.
  """

  @typedoc """
  A NAT pid.
  """
  @type t() :: pid()

  @typedoc """
  A mode that the NAT can run in.
  """
  @type mode() :: :once | :ongoing

  @doc """
  Start a new NAT for a router.

  The NAT should be created before any computers on the network start sending
  traffic. The NAT will not receive any packets from the router until it is added
  to the router with `Day23.Router.set_nat/2`.
  """
  @spec async(Day23.Router.t(), mode) :: Task.t()
  def async(router, mode) do
    Task.async(__MODULE__, :run, [router, mode])
  end

  @doc """
  Runs the NAT in the given mode.

  If `mode` is `:once`, then the NAT will simply return the first packet it
  receives from the router. No restarting of network traffic will occur.

  If `mode` is `:ongoing`, then the NAT will perform its normal function of
  restarting the network traffic when it idles.
  """
  @spec run(Day23.Router.t(), mode) :: {number, number}
  def run(router, mode) do
    Logger.metadata(mode: mode)

    case mode do
      :once -> run_once()
      :ongoing -> loop(router, nil, nil)
    end
  end

  defp run_once do
    receive do
      {:packet, {x, y}} ->
        Logger.debug("nat received packet", x: x, y: y)
        {x, y}
    end
  end

  defp loop(router, last_received, last_broadcast) do
    receive do
      {:packet, {x, y}} ->
        Logger.debug("nat received packet", x: x, y: y)
        loop(router, {x, y}, last_broadcast)

      :all_idle ->
        {x, y} = last_received
        {_, y0} = last_broadcast || {nil, nil}
        Logger.debug("all queues idle", x: x, y: y, prev_y: y0)

        cond do
          y == y0 ->
            {x, y}

          true ->
            Day23.Router.route(router, {0, x, y})
            loop(router, last_received, last_received)
        end
    end
  end
end
