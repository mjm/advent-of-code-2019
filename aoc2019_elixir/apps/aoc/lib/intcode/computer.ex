defmodule Intcode.Computer do
  use Task
  use Bitwise
  require Logger

  defstruct [:memory, :handler, :input_dest, pc: 0]

  def start_link(data, handler) do
    Task.start_link(__MODULE__, :run, [data, handler])
  end

  def async(data, handler) do
    Task.async(__MODULE__, :run, [data, handler])
  end

  def run(data, handler) do
    {:ok, memory} = Intcode.Memory.start_link(data)
    send(self(), :pop_inst)
    :ok = loop(%Intcode.Computer{memory: memory, handler: handler})
    {:ok, memory}
  end

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
        set_param(computer.memory, computer.input_dest, value)
        send(self(), :pop_inst)
        loop(%{computer | input_dest: nil})
      :halt -> :ok
    end
  end

  defp next_instruction(computer) do
    %Intcode.Computer{memory: memory, pc: pc} = computer
    value = Intcode.Memory.get(memory, pc)
    {opcode, modes} = Intcode.Instruction.decode(value)

    params = for {m, i} <- Enum.with_index(modes) do
      {Intcode.Memory.get(memory, pc+i+1), m}
    end

    {:ok, {opcode, List.to_tuple(params)}, %{computer | pc: pc + 1 + Enum.count(params)}}
  end

  defp execute_instruction(inst, computer) do
    %Intcode.Computer{memory: memory, pc: pc} = computer

    Logger.metadata(inst: inspect(inst), pc: pc)
    Logger.debug("executing instruction")

    case inst do
      {:add, {x, y, z}} ->
        xx = get_param(memory, x)
        yy = get_param(memory, y)

        set_param(memory, z, xx + yy)
        send(self(), :pop_inst)
        computer
      {:mult, {x, y, z}} ->
        xx = get_param(memory, x)
        yy = get_param(memory, y)

        set_param(memory, z, xx * yy)
        send(self(), :pop_inst)
        computer
      {:input, {x}} ->
        send(computer.handler, {:input, self()})
        %{computer | input_dest: x}
      {:output, {x}} ->
        send(computer.handler, {:output, self(), get_param(memory, x)})
        send(self(), :pop_inst)
        computer
      {:jump_true, {p, addr}} ->
        send(self(), :pop_inst)
        case get_param(memory, p) do
          0 -> computer
          _ -> %{computer | pc: get_param(memory, addr)}
        end
      {:jump_false, {p, addr}} ->
        send(self(), :pop_inst)
        case get_param(memory, p) do
          0 -> %{computer | pc: get_param(memory, addr)}
          _ -> computer
        end
      {:less_than, {x, y, z}} ->
        cond do
          get_param(memory, x) < get_param(memory, y) ->
            set_param(memory, z, 1)
          true -> set_param(memory, z, 0)
        end
        send(self(), :pop_inst)
        computer
      {:equals, {x, y, z}} ->
        cond do
          get_param(memory, x) == get_param(memory, y) ->
            set_param(memory, z, 1)
          true -> set_param(memory, z, 0)
        end
        send(self(), :pop_inst)
        computer
      {:halt, _} ->
        send(self(), :halt)
        if computer.handler != nil do
          send(computer.handler, {:halt, self()})
        end
        computer
    end
  end

  defp get_param(memory, {i, :abs}) do
    Intcode.Memory.get(memory, i)
  end

  defp get_param(_memory, {i, :imm}) do
    i
  end

  defp set_param(memory, {i, :abs}, value) do
    Intcode.Memory.set(memory, i, value)
  end
end
