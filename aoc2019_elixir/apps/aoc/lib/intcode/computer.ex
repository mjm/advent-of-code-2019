defmodule Intcode.Computer do
  alias Intcode.{Computer, Instruction, Memory}

  @moduledoc """
  A process that runs a single Intcode program to completion.
  """

  @typedoc """
  The PID for an Intcode computer process.
  """
  @type t :: pid

  @typedoc """
  The PID for a process that is able to handle messages from the Intcode computer.

  ## Received messages

  A handler can be any process as long as it understands the following messages:

  *   `{:input, pid}`: The `t:Intcode.Computer.t/0` at `pid` is waiting for an input value.
      Execution of the Intcode program will pause until the handler sends an input value to
      `pid` using `Intcode.send_input/2`.
  *   `{:output, pid, value}`: The `t:Intcode.Computer.t/0` at `pid` produced an output
      `value`. The handler can deal with this as needed.
  *   `{:halt, pid}`: The `t:Intcode.Computer.t/0` at `pid` has halted. The handler can
      use this signal to perform post-processing and terminate if it hasn't already.
  """
  @type handler :: pid

  use Task
  use Bitwise
  require Logger

  defstruct [:memory, :handler, :input_dest, pc: 0, relative_base: 0]

  def start_link(data, handler) do
    Task.start_link(__MODULE__, :run, [data, handler])
  end

  @doc """
  Start an Intcode computer as an async `Task`.

  This allows waiting for the computer to halt using `Task.await/2`.
  """
  @spec async(Memory.data, handler) :: Task.t
  def async(data, handler) do
    Task.async(__MODULE__, :run, [data, handler])
  end

  @doc """
  Runs the Intcode computer.

  This should generally not be called directly, as it will block the process until the program
  in the memory data finishes. Instead, you can use `Intcode.Computer.async` to spawn a task
  to run the computer in another process.
  """
  @spec run(Memory.data, handler) :: {:ok, Memory.t}
  def run(data, handler) do
    {:ok, memory} = Memory.start_link(data)
    send(self(), :pop_inst)
    :ok = loop(%Computer{memory: memory, handler: handler})
    {:ok, memory}
  end

  @doc """
  Provides an input value to the computer.

  This should only be called by the handler when it receives a request for input from the
  computer. Sending input when the computer hasn't asked for it will result in undefined
  behavior.
  """
  @spec send_input(t, number) :: any
  def send_input(computer, value) do
    send(computer, {:input, value})
  end

  defp loop(computer) do
    receive do
      :pop_inst ->
        {:ok, inst, computer} = next_instruction(computer)
        send(self(), {:exec_inst, inst})
        loop(computer)

      {:exec_inst, inst} ->
        loop(execute_instruction(inst, computer))

      {:input, value} ->
        set_param(computer, computer.input_dest, value)
        send(self(), :pop_inst)
        loop(%{computer | input_dest: nil})

      :halt ->
        :ok
    end
  end

  defp next_instruction(computer) do
    %Computer{memory: memory, pc: pc} = computer
    value = Memory.get(memory, pc)
    {opcode, modes} = Instruction.decode(value)

    params =
      for {m, i} <- Enum.with_index(modes) do
        {Memory.get(memory, pc + i + 1), m}
      end

    {:ok, {opcode, List.to_tuple(params)}, %{computer | pc: pc + 1 + Enum.count(params)}}
  end

  defp execute_instruction(inst, computer) do
    %Computer{pc: pc} = computer

    Logger.metadata(inst: inspect(inst), pc: pc)
    Logger.debug("executing instruction")

    case inst do
      {:add, {x, y, z}} ->
        xx = get_param(computer, x)
        yy = get_param(computer, y)

        set_param(computer, z, xx + yy)
        send(self(), :pop_inst)
        computer

      {:mult, {x, y, z}} ->
        xx = get_param(computer, x)
        yy = get_param(computer, y)

        set_param(computer, z, xx * yy)
        send(self(), :pop_inst)
        computer

      {:input, {x}} ->
        send(computer.handler, {:input, self()})
        %{computer | input_dest: x}

      {:output, {x}} ->
        send(computer.handler, {:output, self(), get_param(computer, x)})
        send(self(), :pop_inst)
        computer

      {:jump_true, {p, addr}} ->
        send(self(), :pop_inst)

        case get_param(computer, p) do
          0 -> computer
          _ -> %{computer | pc: get_param(computer, addr)}
        end

      {:jump_false, {p, addr}} ->
        send(self(), :pop_inst)

        case get_param(computer, p) do
          0 -> %{computer | pc: get_param(computer, addr)}
          _ -> computer
        end

      {:less_than, {x, y, z}} ->
        cond do
          get_param(computer, x) < get_param(computer, y) ->
            set_param(computer, z, 1)

          true ->
            set_param(computer, z, 0)
        end

        send(self(), :pop_inst)
        computer

      {:equals, {x, y, z}} ->
        cond do
          get_param(computer, x) == get_param(computer, y) ->
            set_param(computer, z, 1)

          true ->
            set_param(computer, z, 0)
        end

        send(self(), :pop_inst)
        computer

      {:add_rb, {x}} ->
        send(self(), :pop_inst)

        Map.update!(computer, :relative_base, fn rb ->
          rb + get_param(computer, x)
        end)

      {:halt, _} ->
        send(self(), :halt)

        if computer.handler != nil do
          send(computer.handler, {:halt, self()})
        end

        computer
    end
  end

  defp get_param(%Computer{memory: memory}, {i, :abs}) do
    Memory.get(memory, i)
  end

  defp get_param(_computer, {i, :imm}) do
    i
  end

  defp get_param(%Computer{memory: memory, relative_base: base}, {i, :rel}) do
    Memory.get(memory, base + i)
  end

  defp set_param(%Computer{memory: memory}, {i, :abs}, value) do
    Memory.set(memory, i, value)
  end

  defp set_param(%Computer{memory: memory, relative_base: base}, {i, :rel}, value) do
    Memory.set(memory, base + i, value)
  end
end
