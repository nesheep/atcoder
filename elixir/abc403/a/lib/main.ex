defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> IO.puts()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input([_ | a]), do: Enum.map(a, &String.to_integer/1)

  defp solve(a) do
    a
    |> Enum.with_index(1)
    |> Enum.filter(fn {_, i} -> rem(i, 2) == 1 end)
    |> Enum.map(fn {x, _} -> x end)
    |> Enum.sum()
  end
end
