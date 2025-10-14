defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()
  defp parse_input(l), do: Enum.map(l, &String.to_integer/1)
  defp solve([x, c]), do: div(x, 1000 + c) * 1000
end
