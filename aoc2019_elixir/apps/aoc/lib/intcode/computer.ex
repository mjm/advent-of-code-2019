defmodule Intcode.Computer do
  use Task
  require Logger

  def start_link(data, handler) do
    Task.start_link(__MODULE__, :run, [data, handler])
  end

  def async(data, handler) do
    Task.async(__MODULE__, :run, [data, handler])
  end

  def run(data, _handler) do
    {:ok, memory} = Intcode.Memory.start_link(data)
    send(self(), :pop_inst)
    :ok = loop({memory, 0})
    {:ok, memory}
  end

  defp loop(state) do
    {memory, pc} = state
    receive do
      :pop_inst ->
        {:ok, inst, new_pc} = next_instruction(memory, pc)
        send(self(), {:exec_inst, inst})
        loop({memory, new_pc})
      {:exec_inst, inst} ->
        new_pc = execute_instruction(inst, memory, pc)
        loop({memory, new_pc})
      :halt -> :ok
    end
  end

  defp next_instruction(memory, pc) do
    value = Intcode.Memory.get(memory, pc)
    opcode = Intcode.Instruction.opcode(value)
    num_params = Intcode.Instruction.param_count(opcode)

    params = case num_params do
      0 -> []
      _ -> for i <- 1..num_params do
        Intcode.Memory.get(memory, pc+i)
      end
    end

    {:ok, {opcode, List.to_tuple(params)}, pc + 1 + num_params}
  end

  defp execute_instruction(inst, memory, pc) do
    Logger.metadata(inst: inspect(inst), pc: pc)
    Logger.debug("executing instruction")

    case inst do
      {:add, {x, y, z}} ->
        xx = Intcode.Memory.get(memory, x)
        yy = Intcode.Memory.get(memory, y)

        Intcode.Memory.set(memory, z, xx + yy)
        send(self(), :pop_inst)
        pc
      {:mult, {x, y, z}} ->
        xx = Intcode.Memory.get(memory, x)
        yy = Intcode.Memory.get(memory, y)

        Intcode.Memory.set(memory, z, xx * yy)
        send(self(), :pop_inst)
        pc
      {:halt, _} ->
        send(self(), :halt)
        pc
    end
  end
end
