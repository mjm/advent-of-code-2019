require Logger

defmodule Day23.NAT do
  @typedoc """
  A NAT pid.
  """
  @type t() :: pid()

  @typedoc """
  A mode that the NAT can run in.
  """
  @type mode() :: :once | :ongoing

  @spec async(Day23.Router.t, mode) :: Task.t
  def async(router, mode) do
    Task.async(__MODULE__, :run, [router, mode])
  end

  @spec run(Day23.Router.t, mode) :: {number, number}
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
        Logger.debug("nat received packet", [x: x, y: y])
        {x, y}
    end
  end

  defp loop(router, last_received, last_broadcast) do
    receive do
      {:packet, {x, y}} ->
        Logger.debug("nat received packet", [x: x, y: y])
        loop(router, {x, y}, last_broadcast)

      :all_idle ->
        {x, y} = last_received
        {_, y0} = last_broadcast || {nil, nil}
        Logger.debug("all queues idle", [x: x, y: y, prev_y: y0])

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
