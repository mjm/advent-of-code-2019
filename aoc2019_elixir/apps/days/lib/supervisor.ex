defmodule Days.Supervisor do
  @moduledoc false

  use Supervisor

  def start_link(opts) do
    Supervisor.start_link(__MODULE__, :ok, opts)
  end

  @impl true
  def init(:ok) do
    children = [
      {Day01, name: Day01}
    ]

    Supervisor.init(children, strategy: :one_for_one)
  end
end
