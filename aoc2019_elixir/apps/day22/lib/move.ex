defmodule Day22.Move do
  @moduledoc """
  Functions and types for working with shuffling moves.
  """

  import Day22.Math

  @typedoc """
  A shuffling move.
  """
  @type t :: {:deal, number} | {:cut, number} | :reverse

  @typedoc """
  The result of one or more shuffling moves as a pair of multipler and offset.

  This is a collapsed representation of moves that remains efficient even for
  working with large numbers of moves and repetitions. The new position of a
  card can be found by multiplying the card's position by the multiplier, then
  adding the offset.
  """
  @type result :: {number, number}

  @doc """
  Parse a string into a shuffling move.

  ## Examples

      iex> Day22.Move.from_string("deal with increment 8")
      {:deal, 8}
      iex> Day22.Move.from_string("cut -103")
      {:cut, -103}
      iex> Day22.Move.from_string("deal into new stack")
      :reverse
  """
  @spec from_string(String.t()) :: t
  def from_string(str)

  def from_string("deal with increment " <> n), do: {:deal, String.to_integer(n)}
  def from_string("cut " <> n), do: {:cut, String.to_integer(n)}
  def from_string("deal into new stack"), do: :reverse

  @doc """
  Performs a shuffling move, altering the multiplier and offset accordingly.

  ## Examples

  ### From starting position

      iex> Day22.Move.perform({1, 0}, 7, {:deal, 3})
      {3, 0}
      iex> Day22.Move.perform({1, 0}, 7, {:cut, 4})
      {1, 3}
      iex> Day22.Move.perform({1, 0}, 7, :reverse)
      {6, 6}

  ### From more interesting places

      iex> Day22.Move.perform({1, 4}, 10, {:deal, 7})
      {7, 8}
      iex> Day22.Move.perform({7, 8}, 10, :reverse)
      {3, 1}
      iex> Day22.Move.perform({3, 1}, 10, {:cut, -3})
      {3, 4}

  """
  @spec perform(result, number, t) :: result
  def perform(result, size, move)

  def perform({m, b}, size, {:deal, n}), do: mod({m * n, b * n}, size)
  def perform({m, b}, size, {:cut, n}), do: mod({m, b - n}, size)
  def perform({m, b}, size, :reverse), do: mod({-m, -b - 1}, size)

  defp mod({m, b}, size) do
    {Integer.mod(m, size), Integer.mod(b, size)}
  end

  @doc """
  Perform a sequence of a list of moves, returning the multiplier and offset
  needed to determine the final position of a card after applying the moves.

  ## Examples

      iex> Day22.Move.perform_list([{:cut, 6}, {:deal, 7}, :reverse], 11)
      {4, 8}
      iex> Day22.Move.perform_list([{:deal, 7}, :reverse, :reverse], 11)
      {7, 0}

  """
  @spec perform_list(list(t), number) :: result
  def perform_list(moves, size) do
    Enum.reduce(moves, {1, 0}, &perform(&2, size, &1))
  end

  @doc """
  Perform a shuffling move in reverse, giving a result that can find the
  original position of a card assuming it was shuffled with the given move.

  ## Examples

  ### Back to starting position

      iex> Day22.Move.undo({3, 0}, 7, {:deal, 3})
      {1, 0}
      iex> Day22.Move.undo({1, 3}, 7, {:cut, 4})
      {1, 0}
      iex> Day22.Move.undo({6, 6}, 7, :reverse)
      {1, 0}

  ### Back from more interesting places

      iex> Day22.Move.undo({7, 8}, 10, {:deal, 7})
      {1, 4}
      iex> Day22.Move.undo({3, 1}, 10, :reverse)
      {7, 8}
      iex> Day22.Move.undo({3, 4}, 10, {:cut, -3})
      {3, 1}

  """
  @spec undo(result, number, t) :: result
  def undo(result, size, move)

  def undo({m, b}, size, {:deal, n}), do: mod({m * mod_inv(n, size), b * mod_inv(n, size)}, size)
  def undo({m, b}, size, {:cut, n}), do: mod({m, b + n}, size)
  def undo({m, b}, size, :reverse), do: perform({m, b}, size, :reverse)

  @doc """
  Undoes a list of moves, returning the multiplier and offset needed to
  determine the original position of a card after having the moves performed.

  ## Examples

      iex> Day22.Move.undo_list([{:cut, 6}, {:deal, 7}, :reverse], 11)
      {3, 9}
      iex> Day22.Move.undo_list([{:deal, 7}, :reverse, :reverse], 11)
      {8, 0}

  """
  @spec undo_list(list(t), number) :: result
  def undo_list(moves, size) do
    Enum.reduce(Enum.reverse(moves), {1, 0}, &undo(&2, size, &1))
  end

  @doc """
  Repeats a move or sequence of moves a given number of times.

  ## Examples

      iex> Day22.Move.repeat({4, 3}, 11, 0)
      {1, 0}
      iex> Day22.Move.repeat({4, 3}, 11, 1)
      {4, 3}
      iex> Day22.Move.repeat({4, 3}, 11, 2)
      {5, 4}
      iex> Day22.Move.repeat({4, 3}, 11, 3)
      {9, 8}

  """
  @spec repeat(result, number, number) :: result
  def repeat(result, size, times)

  def repeat({m, b}, size, times) do
    mtimes = pow(m, times, size)

    mod(
      {
        mtimes,
        b * (mtimes - 1) * mod_inv(m - 1, size)
      },
      size
    )
  end
end
