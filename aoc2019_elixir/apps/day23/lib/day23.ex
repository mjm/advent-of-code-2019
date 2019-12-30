defmodule Day23 do
  @moduledoc """
  [Day 23: Category Six](https://adventofcode.com/2019/day/23)
  """

  @behaviour Aoc.Problem

  @impl Aoc.Problem
  def input do
    Intcode.Memory.from_string(File.read!("../day23/input.txt"))
  end

  @impl Aoc.Problem
  def part1(input) do
    router = Day23.Router.async()
    nat = Day23.NAT.async(router.pid, :once)
    Day23.Router.set_nat(router.pid, nat.pid)

    queues = for _ <- 1..50 do
      q = Day23.PacketQueue.async()
      Day23.Router.add_queue(router.pid, q.pid)
      q
    end

    for q <- queues do
      Intcode.Computer.async(input, q.pid)
    end

    Task.await(nat)
  end

  @impl Aoc.Problem
  def part2(input) do
    router = Day23.Router.async()
    nat = Day23.NAT.async(router.pid, :ongoing)
    Day23.Router.set_nat(router.pid, nat.pid)

    queues = for _ <- 1..50 do
      q = Day23.PacketQueue.async()
      Day23.Router.add_queue(router.pid, q.pid)
      q
    end

    for q <- queues do
      Intcode.Computer.async(input, q.pid)
    end

    Task.await(nat, 60000)
  end
end
