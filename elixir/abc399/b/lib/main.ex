defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> Enum.each(&IO.puts/1)
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input([_ | p]), do: p |> Enum.map(&String.to_integer/1)

  defp solve(p) do
    p
    |> Enum.with_index(1)
    |> Enum.sort(fn {a, _}, {b, _} -> b < a end)
    |> Enum.with_index(1)
    |> Enum.map(fn {{s, i}, r} -> {r, i, s} end)
    |> Enum.scan(fn {_, i1, s1} = a, {r2, _, s2} ->
      if s1 == s2, do: {r2, i1, s1}, else: a
    end)
    |> Enum.sort(fn {_, i1, _}, {_, i2, _} -> i1 < i2 end)
    |> Enum.map(fn {r, _, _} -> r end)
  end
end
