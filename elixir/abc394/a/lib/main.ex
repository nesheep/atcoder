defmodule Main do
  def stdin_values, do: IO.read(:eof) |> String.split()

  def parse_input([s]), do: s

  def solve(s) when is_binary(s) do
    String.graphemes(s)
    |> Enum.filter(&(&1 == "2"))
    |> Enum.join("")
  end

  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end
end
