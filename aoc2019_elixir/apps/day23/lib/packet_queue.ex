defmodule Day23.PacketQueue do
  @moduledoc """
  A packet queue for a single Intcode computer in a network.

  The packet queue receives packet messages from a `Day23.Router` and queues them up
  so they can be provided to an Intcode computer upon request.

  The packet queue is the `t:Intcode.Computer.handler/0` for a computer in a network.
  When the computer requests input, the packet queue will provide the `x` and `y`
  values of the next packet in the queue if there is one. Otherwise, it will provide
  the computer with `-1` to indicate there is no packet.

  The packet queue also watches the output of the computer and sends the packets the
  computer outputs to the `Day23.Router` so they can be sent to the appropriate queue.
  """

  @typedoc """
  A packet queue PID.
  """
  @type t() :: pid()

  defstruct router: nil,
            addr: nil,
            queue: [],
            awaiting_addr: nil,
            partial_packet: [],
            idle_count: 0

  @doc """
  Starts a new task running a packet queue.

  Use the task's PID as the handler for an Intcode computer in the network.
  """
  @spec async :: Task.t()
  def async do
    Task.async(__MODULE__, :run, [])
  end

  @doc """
  Run the loop for the packet queue to process messages.
  """
  @spec run :: none
  def run do
    loop(%Day23.PacketQueue{})
  end

  @doc """
  Assign a network address to the queue.

  This informs the queue which router the computer it manages belongs to.
  It also causes the packet queue to inform the Intcode computer of its
  address, which is required before the computer can begin sending traffic.
  """
  @spec assign_addr(t, Day23.Router.addr(), Day23.Router.t()) :: any
  def assign_addr(pid, addr, router) do
    send(pid, {:assign_addr, addr, router})
  end

  @doc """
  Add a new packet to the queue.

  The packet will be delivered to the Intcode computer the next time it asks
  for input.
  """
  @spec enqueue(t, {number, number}) :: any
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
