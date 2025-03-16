defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def parse_input([x]), do: String.to_float(x)

  def solve(x) when x >= 38.0, do: 1
  def solve(x) when x >= 37.5, do: 2
  def solve(_), do: 3

  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end
end
