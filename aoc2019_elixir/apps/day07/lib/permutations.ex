defmodule Day07.Permutations do
  @moduledoc """
  Helper function for computing every permutation of an enumerable.
  """

  @doc """
  Returns a list of all possible permutations of the given enumerable.
  """
  @spec all(Enum.t) :: list(list(any))
  def all(nums) do
    {_, acc} = generate(Enum.count(nums), Enum.to_list(nums), [])
    acc
  end

  defp generate(k, nums, acc) do
    case k do
      1 ->
        {nums, [nums | acc]}

      _ ->
        Enum.reduce(0..(k - 2), generate(k - 1, nums, acc), fn i, {nums, acc} ->
          new_nums =
            cond do
              Integer.mod(k, 2) == 0 -> swap(nums, i, k - 1)
              true -> swap(nums, 0, k - 1)
            end

          generate(k - 1, new_nums, acc)
        end)
    end
  end

  defp swap(list, a, b) do
    List.replace_at(List.replace_at(list, a, Enum.at(list, b)), b, Enum.at(list, a))
  end
end
