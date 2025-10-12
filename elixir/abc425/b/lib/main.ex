defmodule Main do
  def main do
    stdin_values()
    |> parse_input()
    |> solve()
    |> output()
  end

  defp stdin_values, do: IO.read(:eof) |> String.split()

  defp parse_input(l) when is_list(l) do
    [n | a] = Enum.map(l, &String.to_integer/1)
    {n, a}
  end

  defp solve({n, a}) when is_list(a) do
    a_i = Enum.with_index(a, 1)
    s = Enum.filter(a, &(&1 != -1)) |> MapSet.new()

    q =
      Enum.zip(
        Enum.filter(a_i, fn {x, _} -> x == -1 end) |> Enum.map(fn {_, i} -> i end),
        Enum.filter(1..n, fn x -> !MapSet.member?(s, x) end)
      )
      |> Map.new()

    p =
      Enum.map(a_i, fn
        {x, _} when x != -1 -> x
        {_, i} -> Map.get(q, i)
      end)

    if Enum.uniq(p) |> length() == n, do: p, else: nil
  end

  defp output(p) when is_list(p) do
    IO.puts("Yes")
    IO.puts(Enum.join(p, " "))
  end

  defp output(_), do: IO.puts("No")
end
