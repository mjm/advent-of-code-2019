defmodule Day14.Table do
  @moduledoc """
  Functions for creating and working with a table of chemical reactions.

  A chemical reaction table is represented as a digraph where each vertex is a
  material. The graph has edges from the output of a reaction to each input
  required for the reaction, with labels for the quantity of each component.
  This structure supports being able to determine the amount of ore needed to
  create different amounts of fuel (or other materials).
  """

  alias Day14.Reaction, as: Reaction

  @typedoc """
  A table of reactions.
  """
  @opaque t :: :digraph.graph

  @doc """
  Create a new reaction table from a string of lines describing each reaction.
  """
  @spec from_string(String.t) :: t
  def from_string(str) do
    String.split(str, "\n") |> Enum.map(&Reaction.from_string/1) |> new
  end

  @doc """
  Create a new reaction table from a list of reactions.
  """
  @spec new(list(Reaction.t)) :: t
  def new(reactions) do
    g = :digraph.new([:acyclic])
    :digraph.add_vertex(g, "ORE")
    Enum.each reactions, fn {_, {n, name}} ->
      :digraph.add_vertex(g, name, n)
    end
    Enum.each reactions, fn {ins, {_, name_out}} ->
      Enum.each ins, fn {n_in, name_in} ->
        [:"$e" | _] = :digraph.add_edge(g, name_out, name_in, n_in)
      end
    end
    g
  end

  defp sorted_materials(table) do
    :digraph_utils.topsort(table)
  end

  @doc ~S|
  Calculate the amount of ore required to produce a certain amount of a material.

  ## Examples

      iex> table = Day14.Table.from_string(String.trim("""
      ...> 10 ORE => 10 A
      ...> 1 ORE => 1 B
      ...> 7 A, 1 B => 1 C
      ...> 7 A, 1 C => 1 D
      ...> 7 A, 1 D => 1 E
      ...> 7 A, 1 E => 1 FUEL
      ...> """))
      iex> Day14.Table.required_ore(table, {1, "FUEL"})
      31

  |
  @spec required_ore(t, Reaction.component) :: number
  def required_ore(table, component)
  def required_ore(table, {amount, material}) do
    required_ore(table, sorted_materials(table), %{material => amount})
  end

  defp required_ore(_, ["ORE" | _], %{"ORE" => n}), do: n
  defp required_ore(table, [mat | rest], reqs) do
    case Map.get(reqs, mat, 0) do
      0 -> required_ore(table, rest, reqs)
      amount ->
        ratio = ceil(amount / output_amount(table, mat))
        required_ore(table, rest,
          Enum.reduce(reagents(table, mat), Map.delete(reqs, mat), fn {in_mat, n_in}, reqs ->
            new_amount = ratio * n_in
            Map.update(reqs, in_mat, new_amount, &(&1 + new_amount))
          end))
    end
  end
  
  defp output_amount(table, mat) do
    {_, amount} = :digraph.vertex(table, mat)
    amount
  end

  defp reagents(table, mat) do
    :digraph.out_edges(table, mat) |>
      Enum.map(fn e ->
        {_, _, input, amount} = :digraph.edge(table, e)
        {input, amount}
      end)
  end

  @doc """
  Determines the maximum amount of fuel that can be made by starting with the
  given amount of oil.
  """
  @spec fuel_possible(t, number) :: number
  def fuel_possible(table, ore) do
    fuel_possible(table, ore, floor(ore / required_ore(table, {1, "FUEL"})))
  end

  defp fuel_possible(table, max_ore, lower_bound) do
    ore = required_ore(table, {lower_bound, "FUEL"})
    cond do
      ore < max_ore -> fuel_possible(table, max_ore, lower_bound + 50000)
      true -> fuel_possible(table, max_ore, lower_bound - 50000, lower_bound)
    end
  end

  defp fuel_possible(table, max_ore, lower_bound, upper_bound) do
    ore = required_ore(table, {lower_bound, "FUEL"})
    cond do
      ore > max_ore -> lower_bound - 1
      true -> fuel_possible(table, max_ore, lower_bound + 1, upper_bound)
    end
  end
end
