defmodule Aoc2019Elixir.MixProject do
  use Mix.Project

  def project do
    [
      apps_path: "apps",
      version: "0.1.0",
      start_permanent: Mix.env() == :prod,
      deps: deps(),
      docs: [
        groups_for_modules: [
          Common: [~r/^Aoc/, ~r/^Days/],
          Intcode: ~r/^Intcode/,
          Problems: ~r/^Day[[:digit:]]{2}/
        ]
      ]
    ]
  end

  # Dependencies listed here are available only for this
  # project and cannot be accessed from applications inside
  # the apps folder.
  #
  # Run "mix help deps" for examples and options.
  defp deps do
    [
      {:ex_doc, "~> 0.21", only: :dev, runtime: false}
    ]
  end
end
