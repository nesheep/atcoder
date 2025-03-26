defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def parse_input([n]), do: String.to_integer(n)

  def a(n), do: String.duplicate("-", div(n - 1, 2))
  def b(n) when rem(n, 2) == 0, do: "=="
  def b(_), do: "="

  def solve(n), do: "#{a(n)}#{b(n)}#{a(n)}"

  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end
end
