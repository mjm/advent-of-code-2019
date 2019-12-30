defmodule Day22.Math do
  @moduledoc """
  Helpful math functions for doing modular arithmetic.
  """

  import Bitwise

  @doc """
  Calculates the modular inverse of `a` mod `b`.
  """
  @spec mod_inv(number, number) :: number
  def mod_inv(a, b) do
    {1, x, _} = extended_gcd(a, b)
    x
  end

  defp extended_gcd(a, b) do
    extended_gcd_loop({b, a}, {0, 1}, {1, 0})
  end

  defp extended_gcd_loop({0, r0}, {_, s0}, {_, t0}), do: {r0, s0, t0}

  defp extended_gcd_loop({r, r0}, {s, s0}, {t, t0}) do
    q = Integer.floor_div(r0, r)
    extended_gcd_loop({r0 - q * r, r}, {s0 - q * s, s}, {t0 - q * t, t})
  end

  @doc """
  Performs modular exponentiation on integers.

  ## Examples

      iex> Day22.Math.pow(5, 2, 3)
      1
      iex> Day22.Math.pow(3, 4, 17)
      13

  """
  @spec pow(number, number, number) :: number
  def pow(b, e, m) do
    pow(Integer.mod(b, m), e, m, 1)
  end

  defp pow(_, _, 1, _), do: 0
  defp pow(_, 0, _, r), do: r

  defp pow(b, e, m, r) do
    r =
      case Integer.mod(e, 2) do
        1 -> Integer.mod(r * b, m)
        _ -> r
      end

    pow(Integer.mod(b * b, m), e >>> 1, m, r)
  end
end
