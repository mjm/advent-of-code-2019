defmodule Day23.PacketQueue do
  @typedoc """
  A packet queue PID.
  """
  @type t() :: pid()

  defstruct router: nil, addr: nil, queue: [], awaiting_addr: nil, partial_packet: [], idle_count: 0

  def async do
    Task.async(__MODULE__, :run, [])
  end

  def run do
    loop(%Day23.PacketQueue{})
  end

  @spec assign_addr(t, Day23.Router.addr, Day23.Router.t) :: any()
  def assign_addr(pid, addr, router) do
    send(pid, {:assign_addr, addr, router})
  end

  @spec enqueue(t, {number, number}) :: any()
  def enqueue(pid, point) do
    send(pid, {:enqueue, point})
  end

  defp loop(pq) do
    receive do
      {:assign_addr, addr, router} ->
        case pq.awaiting_addr do
          nil -> 
            loop(%{pq | router: router, addr: addr, queue: [addr]})
          pid ->
            Intcode.send_input(pid, addr)
            loop(%{pq | router: router, addr: addr, awaiting_addr: nil})
        end

      {:input, pid} ->
        case pq do
          %Day23.PacketQueue{addr: nil} ->
            loop(%{pq | awaiting_addr: pid})

          %Day23.PacketQueue{queue: []} ->
            Intcode.send_input(pid, -1)

            if pq.idle_count == 3 do
              Day23.Router.report_idle(pq.router, pq.addr)
            end

            loop(%{pq | idle_count: pq.idle_count + 1})

          %Day23.PacketQueue{queue: [hd | tl]} ->
            Intcode.send_input(pid, hd)
            loop(%{pq | queue: tl})
        end

      {:enqueue, {x, y}} ->
        Day23.Router.report_active(pq.router, pq.addr)
        loop(%{pq | queue: pq.queue ++ [x, y], idle_count: 0})

      {:output, _, value} ->
        Day23.Router.report_active(pq.router, pq.addr)
        packet = pq.partial_packet ++ [value]

        case Enum.count(packet) do
          3 ->
            Day23.Router.route(pq.router, List.to_tuple(packet))
            loop(%{pq | partial_packet: [], idle_count: 0})

          _ ->
            loop(%{pq | partial_packet: packet, idle_count: 0})
        end
    end
  end
end
