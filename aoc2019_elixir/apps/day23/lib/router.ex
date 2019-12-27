defmodule Day23.Router do
  @typedoc """
  A router PID.
  """
  @type t() :: pid()

  @typedoc """
  An address to a particular computer in the network.
  """
  @type addr() :: number()

  @typedoc """
  An addressed packet containing a point value.
  """
  @type packet() :: {addr(), number(), number()}

  defstruct queues: %{}, next_addr: 0, idle: MapSet.new(), nat: nil

  def async do
    Task.async(__MODULE__, :run, [])
  end

  def run do
    loop(%Day23.Router{})
  end

  @spec set_nat(t, Day23.NAT.t) :: any
  def set_nat(pid, nat) do
    send(pid, {:set_nat, nat})
  end

  @spec add_queue(t, Day23.PacketQueue.t) :: any
  def add_queue(pid, queue) do
    send(pid, {:add_queue, queue})
  end

  @spec route(t, packet) :: any
  def route(pid, packet) do
    send(pid, {:route_packet, packet})
  end

  @spec report_idle(t, addr) :: any
  def report_idle(pid, addr) do
    send(pid, {:report_idle, addr})
  end

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
