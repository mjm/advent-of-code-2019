defmodule Intcode do
  @moduledoc """
  Helper functions for working with Intcode computers.
  """

  defdelegate send_input(computer, value), to: Intcode.Computer
  defdelegate memory_from_string(str), to: Intcode.Memory, as: :from_string
end
