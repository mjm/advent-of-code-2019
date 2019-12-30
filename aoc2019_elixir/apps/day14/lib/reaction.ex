defmodule Day14.Reaction do
  @typedoc """
  A reaction that consumes one or more input components and produces some output
  component.
  """
  @type t :: {list(component), component}

  @typedoc """
  An input or output to a reaction. Includes the name of the material and the amount
  required or produced by the reaction.
  """
  @type component :: {number, String.t}

  @doc """
  Parse a reaction from a string describing the inputs and outputs.

  ## Examples

      iex> Day14.Reaction.from_string("157 ORE => 5 NZVS")
      {[{157, "ORE"}], {5, "NZVS"}}
      iex> Day14.Reaction.from_string("22 VJHF, 37 MNCFX => 5 FWMGM")
      {[{22, "VJHF"}, {37, "MNCFX"}], {5, "FWMGM"}}

  """
  @spec from_string(String.t) :: t
  def from_string(str) do
    [ins, out] = String.split(str, " => ")
    {
      Enum.map(String.split(ins, ", "), &component_from_string/1),
      component_from_string(out)
    }
  end

  @spec component_from_string(String.t) :: component
  defp component_from_string(str) do
    [amount, label] = String.split(str, " ")
    {String.to_integer(amount), label}
  end
end
