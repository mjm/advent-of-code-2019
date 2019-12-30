defmodule Intcode.Instruction do
  @moduledoc false

  @doc ~S"""
  Decodes an instruction opcode and its parameter modes.

  ## Examples

      iex> Intcode.Instruction.decode(1002)
      {:mult, [:abs, :imm, :abs]}
      iex> Intcode.Instruction.decode(3)
      {:input, [:abs]}
      iex> Intcode.Instruction.decode(21001)
      {:add, [:abs, :imm, :rel]}

  """
  def decode(value) do
    op = Integer.mod(value, 100)
    modes = Integer.floor_div(value, 100)

    code = opcode(op)
    {code, param_modes(param_count(code), modes, [])}
  end

  defp param_modes(n, modes, acc) do
    case n do
      0 ->
        acc

      _ ->
        param_modes(
          n - 1,
          Integer.floor_div(modes, 10),
          acc ++ [param_mode(Integer.mod(modes, 10))]
        )
    end
  end

  def opcode(value) do
    case value do
      1 -> :add
      2 -> :mult
      3 -> :input
      4 -> :output
      5 -> :jump_true
      6 -> :jump_false
      7 -> :less_than
      8 -> :equals
      9 -> :add_rb
      99 -> :halt
    end
  end

  def param_count(code) do
    case code do
      :add -> 3
      :mult -> 3
      :input -> 1
      :output -> 1
      :jump_true -> 2
      :jump_false -> 2
      :less_than -> 3
      :equals -> 3
      :add_rb -> 1
      :halt -> 0
    end
  end

  def param_mode(mode) do
    case mode do
      0 -> :abs
      1 -> :imm
      2 -> :rel
    end
  end
end
