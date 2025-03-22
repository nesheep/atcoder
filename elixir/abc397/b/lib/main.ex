defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def parse_input([s]), do: s

  def solve(s) do
    s
    |> String.graphemes()
    |> Enum.reduce({0, 0}, fn c, {n, a} ->
      case {c, rem(n, 2)} do
        {"i", 0} -> {n + 1, a}
        {"o", 1} -> {n + 1, a}
        _ -> {n + 2, a + 1}
      end
    end)
    |> case do
      {n, a} when rem(n, 2) == 0 -> a
      {_, a} -> a + 1
    end
  end

  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end
end
