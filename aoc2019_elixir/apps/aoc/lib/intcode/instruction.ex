defmodule Intcode.Instruction do
  def opcode(value) do
    case value do
      1 -> :add
      2 -> :mult
      99 -> :halt
    end
  end

  def param_count(code) do
    case code do
      :add -> 3
      :mult -> 3
      :halt -> 0
    end
  end
end
