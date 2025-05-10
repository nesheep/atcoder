defmodule Main do
  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input([_, s, t]), do: {s, t}

  defp solve({s, t}) do
    Enum.zip(String.graphemes(s), String.graphemes(t))
    |> Enum.count(fn {a, b} -> a != b end)
  end

  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end
end
