defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input([s]), do: s

  defp solve(s) do
    String.graphemes(s)
    |> Enum.frequencies()
    |> Enum.find(fn {_, n} -> n == 1 end)
    |> then(fn {c, _} -> c end)
  end
end
