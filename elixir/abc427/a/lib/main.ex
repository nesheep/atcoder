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
    s = String.graphemes(s)
    m = s |> length() |> div(2)

    Enum.with_index(s)
    |> Enum.filter(fn {_, i} -> i != m end)
    |> Enum.map(fn {c, _} -> c end)
    |> to_string()
  end
end
