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
    s
    |> String.to_charlist()
    |> Enum.filter(&(&1 < ?a))
    |> List.to_string()
  end
end
