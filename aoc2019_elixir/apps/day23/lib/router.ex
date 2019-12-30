defmodule Day23.Router do
  @moduledoc """
  A packet router for a network of Intcode computers.

  A router keeps track of a collection of `Day23.PacketQueue`s which are each
  assigned a unique address within the router. When the computers in the network
  produce output, their packet queue sends that output as a packet to the router,
  which then ensures it gets added to the packet queue it is addressed to.

  Packet queues also report their idle status to the router. When all queues are
  idle, the router sends a message to its NAT to indicate this. The NAT may
  respond by sending a message that will kickstart activity again.
  """

  @typedoc """
  A router PID.
  """
  @type t :: pid

  @typedoc """
  An address to a particular computer in the network.
  """
  @type addr :: number

  @typedoc """
  An addressed packet containing a point value.
  """
  @type packet :: {addr, number, number}

  defstruct queues: %{}, next_addr: 0, idle: MapSet.new(), nat: nil

  @doc """
  Starts a new router as an async `Task`.
  """
  @spec async :: Task.t()
  def async do
    Task.async(__MODULE__, :run, [])
  end

  @doc """
  Runs the router's message processing loop forever.
  """
  @spec run :: none
  def run do
    loop(%Day23.Router{})
  end

  @doc """
  Set the `Day23.NAT` for a particular router.
  """
  @spec set_nat(t, Day23.NAT.t()) :: any
  def set_nat(pid, nat) do
    send(pid, {:set_nat, nat})
  end

  @doc """
  Add a new packet queue to the router.

  The router will choose a unique address for the queue and use
  `Day23.PacketQueue.assign_addr/3` to tell the queue its address.
  """
  @spec add_queue(t, Day23.PacketQueue.t()) :: any
  def add_queue(pid, queue) do
    send(pid, {:add_queue, queue})
  end

  @doc """
  Asks the router to route a new packet.

  The router will find the queue that matches the address in the packet
  and add the packet to that queue.
  """
  @spec route(t, packet) :: any
  def route(pid, packet) do
    send(pid, {:route_packet, packet})
  end

  @doc """
  Report that the queue at an address is idle.
  """
  @spec report_idle(t, addr) :: any
  def report_idle(pid, addr) do
    send(pid, {:report_idle, addr})
  end

  @doc """
  Report that the queue at an address is active (not idle).
  """
  @spec report_active(t, addr) :: any
  def report_active(pid, addr) do
    send(pid, {:report_active, addr})
  end

  defp loop(router) do
    %Day23.Router{queues: queues, next_addr: next_addr} = router

    receive do
      {:set_nat, nat} ->
        loop(%{router | nat: nat})

      {:add_queue, queue} ->
        new_queues = Map.put_new(queues, next_addr, queue)
        Day23.PacketQueue.assign_addr(queue, next_addr, self())
        loop(%{router | next_addr: next_addr + 1, queues: new_queues})

      {:route_packet, {addr, x, y}} ->
        case addr do
          255 -> send(router.nat, {:packet, {x, y}})
          _ -> Day23.PacketQueue.enqueue(queues[addr], {x, y})
        end

        loop(router)

      {:report_idle, addr} ->
        idle = MapSet.put(router.idle, addr)

        if MapSet.size(idle) == map_size(queues) do
          send(router.nat, :all_idle)
        end

        loop(%{router | idle: idle})

      {:report_active, addr} ->
        loop(%{router | idle: MapSet.delete(router.idle, addr)})
    end
  end
end
